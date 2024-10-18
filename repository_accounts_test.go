package amocrm_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/2010kira2010/amocrm"
)

func TestAccounts_Current(t *testing.T) {
	noTokenClient := amocrm.New(clientID, clientSecret, redirectURL)

	almostValidClient := amocrm.New(clientID, clientSecret, redirectURL)
	_ = almostValidClient.SetToken(amocrm.NewToken(accessToken, refreshToken, tokenType, time.Time{}))
	_ = almostValidClient.SetDomain("example.amocrm.ru")

	relations := []string{
		amocrm.WithUUID,
		amocrm.WithVersion,
		amocrm.WithAmojoID,
		amocrm.WithTaskTypes,
		amocrm.WithUserGroups,
		amocrm.WithAmojoRights,
		amocrm.WithDatetimeSettings,
	}

	cases := []struct {
		client amocrm.Client
		config amocrm.AccountsConfig
		wanted *amocrm.Account
		error  error
	}{
		{
			client: noTokenClient,
			error:  errors.New("unexpected account relation: example"),
			wanted: (*amocrm.Account)(nil),
			config: amocrm.AccountsConfig{
				Relations: []string{"example"},
			},
		},
		{
			client: noTokenClient,
			error:  errors.New("get accounts: invalid token"),
			wanted: (*amocrm.Account)(nil),
			config: amocrm.AccountsConfig{Relations: relations},
		},
		{
			client: almostValidClient,
			error:  nil,
			wanted: &amocrm.Account{},
			config: amocrm.AccountsConfig{Relations: relations},
		},
	}

	for _, tc := range cases {
		got, err := tc.client.Accounts().Current(tc.config)
		require.Exactly(t, tc.wanted, got)

		if tc.error == nil {
			require.NoError(t, err)
		} else {
			require.EqualError(t, err, tc.error.Error())
		}
	}
}
