package amocrm

import (
	"crypto/rand"
	"encoding/hex"
	"net/url"
)

// Provider is a wrapper for authorization and making requests.
type Client interface {
	AuthorizeURL(state, mode string) (*url.URL, error)
	TokenByCode(code string) (Token, error)
	LoadTokenOrAuthorize(code string) error
	CheckToken() error
	SetToken(token Token) error
	SetDomain(domain string) error
	Accounts() Accounts
	Leads() Leads
	Contacts() Contacts
	Companies() Companies
	Pipelines() Pipelines
	Calls() Calls
	Catalogs() Catalogs
	Users() Users
	EventsV2() EventsV2
}

// Verify interface compliance.
var _ Client = (*amoCRM)(nil)

type amoCRM struct {
	api *api
}

// RandomState generates a new random state.
func RandomState() string {
	// Converting bytes to hex will always double length. Hence, we can reduce
	// the amount of bytes by half to produce the correct length of 32 characters.
	key := make([]byte, 16)

	// https://golang.org/pkg/math/rand/#Rand.Read
	// Ignore errors as it always returns a nil error.
	_, _ = rand.Read(key)

	return hex.EncodeToString(key)
}

// New allocates and returns a new amoCRM API Client.
func New(clientID, clientSecret, redirectURL string) Client {
	return &amoCRM{
		api: newAPI(clientID, clientSecret, redirectURL, nil),
	}
}

func NewWithStorage(tokenStorage TokenStorage, clientID, clientSecret, redirectURL string) Client {
	return &amoCRM{
		api: newAPI(clientID, clientSecret, redirectURL, tokenStorage),
	}
}

// AuthorizeURL returns a URL of page to ask for permissions.
func (a *amoCRM) AuthorizeURL(state, mode string) (*url.URL, error) {
	if state == "" {
		return nil, oauth2Err("empty state")
	}
	if mode != PostMessageMode && mode != PopupMode {
		return nil, oauth2Err("unexpected mode")
	}

	query := url.Values{
		"mode":      []string{mode},
		"state":     []string{state},
		"client_id": []string{a.api.clientID},
	}.Encode()

	authURL := "https://www.amocrm.ru/oauth?" + query

	return url.Parse(authURL)
}

// SetToken stores given token to sign API requests.
func (a *amoCRM) CheckToken() error {
	return a.api.checkToken()
}

// SetToken stores given token to sign API requests.
func (a *amoCRM) SetToken(token Token) error {
	return a.api.setToken(token)
}

// SetToken stores given domain to build accounts-specific API endpoints.
func (a *amoCRM) SetDomain(domain string) error {
	return a.api.setDomain(domain)
}

func (a *amoCRM) LoadTokenOrAuthorize(authCode string) error {
	token, err := a.api.loadToken()
	if err != nil {
		return err
	}

	if token != nil {
		return a.api.setToken(token)
	}

	token, err = a.TokenByCode(authCode)
	if err != nil {
		return err
	}

	return a.api.setToken(token)
}

// TokenByCode makes a handshake with amoCRM, exchanging given
// authorization code for a set of tokens.
func (a *amoCRM) TokenByCode(code string) (Token, error) {
	return a.api.getToken(authorizationCodeGrant, url.Values{
		"code":       []string{code},
		"grant_type": []string{"authorization_code"},
	}, nil)
	return nil, nil
}

// Accounts returns accounts repository.
func (a *amoCRM) Accounts() Accounts {
	return newAccounts(a.api)
}

func (a *amoCRM) Leads() Leads {
	return newLeads(a.api)
}

func (a *amoCRM) Contacts() Contacts {
	return newContacts(a.api)
}

func (a *amoCRM) Companies() Companies {
	return newCompanies(a.api)
}

func (a *amoCRM) Pipelines() Pipelines {
	return newPipelines(a.api)
}

func (a *amoCRM) Calls() Calls {
	return newCalls(a.api)
}

func (a *amoCRM) Catalogs() Catalogs {
	return newCatalogs(a.api)
}

func (a *amoCRM) Tasks() Tasks {
	return newTasks(a.api)
}

func (a *amoCRM) Users() Users {
	return newUsers(a.api)
}

func (a *amoCRM) EventsV2() EventsV2 {
	return newEventsV2(a.api)
}
