package clashsync

import (
	"errors"
	"strings"

	"github.com/aaantiii/goclash"

	"github.com/aaantiii/lostapp/services/clashsync/env"
)

func NewCocClient() (*goclash.Client, error) {
	emails := strings.Split(env.COC_EMAIL.Value(), ",")
	passwords := strings.Split(env.COC_PASSWORD.Value(), ",")
	if len(emails) != len(passwords) {
		return nil, errors.New("invalid COC-API credentials")
	}

	credentials := make(goclash.Credentials, len(emails))
	for i, email := range emails {
		credentials[email] = passwords[i]
	}

	client, err := goclash.New(credentials)
	if err != nil {
		return nil, err
	}

	return client, nil
}
