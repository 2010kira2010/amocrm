package amocrm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// PostMessageMode and PopupMode are the only options for
// amoCRM OAuth2.0 "mode" request parameter.
const (
	PostMessageMode = "post_message"
	PopupMode       = "popup"
	userAgent       = "AmoCRM-API-Golang-Client"
	apiVersion      = uint8(4)
	requestTimeout  = 20 * time.Second
)

type GrantType struct {
	code   string
	fields []string
}

var (
	authorizationCodeGrant = GrantType{
		code:   "authorization_code",
		fields: []string{"code"},
	}
	refreshTokenGrant = GrantType{
		code:   "refresh_token",
		fields: []string{"refresh_token"},
	}
	// clientCredentialsGrant = GrantType{
	// 	code:   "client_credentials",
	// 	fields: []string{},
	// }
	// passwordGrant = GrantType{
	// 	code:   "password",
	// 	fields: []string{"username", "password"},
	// }
)

// api implements Client interface.
type api struct {
	clientID     string
	clientSecret string
	redirectURL  string

	domain string
	token  Token

	http *http.Client

	storage TokenStorage
}

func newAPI(clientID, clientSecret, redirectURL string, storage TokenStorage) *api {
	return &api{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
		http: &http.Client{
			Timeout: requestTimeout,
		},

		storage: storage,
	}
}

func (a *api) do(ep endpoint, method string, q url.Values, h http.Header, data interface{}) (*http.Response, error) {
	if a.token == nil {
		return nil, errors.New("invalid token")
	}

	if a.token.Expired() {
		if err := a.refreshToken(); err != nil {
			return nil, err
		}
	}

	header := a.header()
	for k, v := range h {
		if _, reserved := header[k]; !reserved {
			header[k] = v
		}
	}

	var body io.Reader
	if data != nil {
		header.Set("Content-Type", "application/json")
		jdata, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jdata)
	}

	apiURL, err := a.url(ep.path(), q)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(method, apiURL.String(), body)
	if err != nil {
		return nil, err
	}
	r.Header = header

	return a.http.Do(r)
}

func (a *api) read(response *http.Response, target interface{}) (err error) {
	defer func() {
		if clErr := response.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body: %w", clErr)
			}
		}
	}()

	if response.StatusCode >= 400 {
		var data []byte
		data, err = ioutil.ReadAll(response.Body)
		err = fmt.Errorf("invalid status: %v", string(data))
		return
	}

	err = json.NewDecoder(response.Body).Decode(target)
	return
}

func (a *api) checkToken() error {
	if a.token.Expired() {
		if err := a.refreshToken(); err != nil {
			return err
		}
	}
	return nil
}

func (a *api) setToken(token Token) error {
	if token == nil {
		return errors.New("invalid token")
	}
	a.token = token

	if a.storage != nil {
		return a.storage.SetToken(token)
	}

	return nil
}

func (a *api) loadToken() (Token, error) {
	if a.storage == nil {
		return nil, nil
	}

	return a.storage.GetToken()
}

func (a *api) setDomain(domain string) error {
	if !isValidDomain(domain) {
		return errors.New("invalid domain")
	}

	a.domain = domain
	return nil
}

func (a *api) authorizationURL(state, mode string) (*url.URL, error) {
	if state == "" {
		return nil, oauth2Err("empty state")
	}
	if mode != PostMessageMode && mode != PopupMode {
		return nil, oauth2Err("unexpected mode")
	}

	query := url.Values{
		"mode":      []string{mode},
		"state":     []string{state},
		"client_id": []string{a.clientID},
	}.Encode()

	authURL := "https://www.amocrm.ru/oauth?" + query

	return url.Parse(authURL)
}

func (a *api) getToken(grant GrantType, options url.Values, header http.Header) (Token, error) {
	if !isValidDomain(a.domain) {
		return nil, oauth2Err("invalid accounts domain")
	}

	// Validate required grantType-specific fields
	for _, key := range grant.fields {
		if values, ok := options[key]; len(values) == 0 || !ok {
			return nil, oauth2Err("missing required %s grant parameter %s", grant.code, key)
		}
	}

	// Default request parameters
	data := url.Values{
		"client_id":     []string{a.clientID},
		"client_secret": []string{a.clientSecret},
		"redirect_uri":  []string{a.redirectURL},
		"grant_type":    []string{grant.code},
	}

	// Merge options with default parameters
	for k, v := range options {
		if _, reserved := data[k]; !reserved {
			data[k] = v
		}
	}

	// Set request URL
	tokenURL, err := a.url("/oauth2/access_token", nil)
	if err != nil {
		return nil, oauth2Err("build request url")
	}

	// Set request headers
	reqHeader := a.baseHeader()
	reqHeader["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	for k, v := range header {
		if _, reserved := reqHeader[k]; !reserved {
			reqHeader[k] = v
		}
	}

	// Create request body
	reqBody := ioutil.NopCloser(strings.NewReader(data.Encode()))

	// Build request
	req := &http.Request{
		Method: http.MethodPost,
		Header: reqHeader,
		URL:    tokenURL,
		Body:   reqBody,
	}

	resp, err := a.http.Do(req)
	if err != nil {
		return nil, oauth2Err("send request")
	}

	respBody, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if closeBodyErr := resp.Body.Close(); closeBodyErr != nil {
		return nil, oauth2Err("close response body")
	}
	if err != nil {
		return nil, oauth2Err("fetch response body")
	}

	if statusCode := resp.StatusCode; statusCode < 200 || statusCode > 299 {
		return nil, oauth2Err("fetch token: response: %v - %s, request: %+v", resp.Status, respBody, req)
	}

	var jsonToken tokenJSON
	if err = json.Unmarshal(respBody, &jsonToken); err != nil {
		return nil, oauth2Err("parse token from json")
	}

	token := &tokenSource{
		accessToken:  jsonToken.AccessToken,
		tokenType:    jsonToken.TokenType,
		refreshToken: jsonToken.RefreshToken,
		expiresAt:    time.Now().Add(time.Duration(jsonToken.ExpiresIn) * time.Second),
	}

	if token.accessToken == "" {
		return nil, oauth2Err("server response missing access_token")
	}

	return token, nil
}

func (a *api) refreshToken() error {
	if a.token.RefreshToken() == "" {
		return oauth2Err("empty refresh token")
	}

	token, err := a.getToken(refreshTokenGrant, url.Values{
		"grant_type":    []string{"refresh_token"},
		"refresh_token": []string{a.token.RefreshToken()},
	}, nil)
	if err != nil {
		return err
	}

	a.token = token

	data, err := ioutil.ReadFile("amocrm_token.json")
	if err != nil {
		log.Fatalf("Error reading settings file: %v", err)
	}

	var settings crmConnJSON
	if err := json.Unmarshal(data, &settings); err != nil {
		log.Fatalf("Error unmarshalling settings JSON: %v", err)
	}
	settings.RefreshToken = token.RefreshToken()
	settings.AccessToken = token.AccessToken()
	settings.AccessTokenExpiresAt = strconv.FormatInt(token.ExpiresAt().Unix(), 10)

	// Кодирование измененного JSON
	updatedJsonData, err := json.MarshalIndent(settings, "", "    ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("amocrm_token.json", updatedJsonData, os.ModePerm); err != nil {
		fmt.Println("set WriteFile TokenStored:", err)
	}
	if a.storage != nil {
		return a.storage.SetToken(token)
	}

	return nil
}

func (a *api) url(path string, q url.Values) (*url.URL, error) {
	if !isValidDomain(a.domain) {
		return nil, oauth2Err("invalid accounts domain")
	}

	endpointURL := "https://" + a.domain + path + "?" + q.Encode()

	return url.Parse(endpointURL)
}

func (a *api) header() http.Header {
	authHeader := a.token.TokenType() + " " + a.token.AccessToken()

	header := a.baseHeader()
	header["Authorization"] = []string{authHeader}

	return header
}

func (a *api) baseHeader() http.Header {
	return http.Header{
		"User-Agent": []string{userAgent},
	}
}

func isValidDomain(domain string) bool {
	if domain == "" {
		return false
	}

	parts := strings.Split(domain, ".")
	if len(parts) != 3 ||
		parts[0] == "" ||
		parts[0] == "www" ||
		len(parts[0]) > 63 ||
		parts[1] != "amocrm" ||
		parts[2] != "ru" && parts[2] != "com" {
		return false
	}

	return true
}

func oauth2Err(format string, args ...interface{}) error {
	return fmt.Errorf("oauth2: "+format, args...)
}
