package client

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/amaanq/coc.go"

	"bot/env"
)

var ErrCocMaintenance = errors.New("inMaintenance Message: API is currently in maintenance, please come back later")

type CocClient struct {
	*coc.Client // create new type to extend it
}

func NewCocClient() (*CocClient, error) {
	emails := strings.Split(env.COC_API_EMAILS.Value(), ",")
	passwords := strings.Split(env.COC_API_PASSWORDS.Value(), ",")
	if len(emails) != len(passwords) {
		return nil, errors.New("invalid COC-API credentials")
	}

	credentials := make(map[string]string, len(emails))
	for i, email := range emails {
		credentials[email] = passwords[i]
	}

	client, err := coc.New(credentials)
	if err != nil {
		return nil, fmt.Errorf("error creating COC-Client: %v", err)
	}

	cocClient := &CocClient{Client: client}
	return cocClient, nil
}

func (client *CocClient) GetClans(tags []string) ([]*coc.Clan, error) {
	if tags == nil || len(tags) == 0 {
		return nil, errors.New("invalid tags to fetch clans")
	}

	var wg sync.WaitGroup
	clans, errs := make([]*coc.Clan, len(tags)), make([]error, len(tags))

	wg.Add(len(tags))
	for i := range tags {
		go func(n int) {
			defer wg.Done()
			clans[n], errs[n] = client.GetClan(tags[n])
		}(i)
	}
	wg.Wait()

	return clans, errors.Join(errs...)
}

func (client *CocClient) IsMaintenanceErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), ErrCocMaintenance.Error())
}
