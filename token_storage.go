// The MIT License (MIT)
//
// Copyright (c) 2021 Alexey Khan
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package amocrm

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type TokenStorage interface {
	SetToken(Token) error
	GetToken() (Token, error)
}

type JSONFileTokenStorage struct {
	File string
}

type JSONToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (self JSONFileTokenStorage) SetToken(token Token) error {
	jt := JSONToken{
		AccessToken:  token.AccessToken(),
		RefreshToken: token.RefreshToken(),
		TokenType:    token.TokenType(),
		ExpiresAt:    token.ExpiresAt(),
	}

	data, err := json.Marshal(jt)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(self.File, data, os.ModePerm)
}

func (self JSONFileTokenStorage) GetToken() (Token, error) {
	data, err := ioutil.ReadFile(self.File)
	if err != nil {
		return nil, nil
	}

	var jt JSONToken
	if err := json.Unmarshal(data, &jt); err != nil {
		return nil, err
	}

	return NewToken(jt.AccessToken, jt.RefreshToken, jt.TokenType, jt.ExpiresAt), nil
}
