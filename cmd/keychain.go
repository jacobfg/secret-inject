package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/keybase/go-keychain"
)

var (
	ErrCredentialsNotFound = errors.New("credentials not found in native keychain")
	serverURL              = "https://api.github.com"
)

func GetSecret(serverUrl string, label string) (string, string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassInternetPassword)
	query.SetLabel(label)

	err := splitServer3(serverURL, query)
	if err != nil {
		fmt.Print(err)
		// 	return "", "", err
	}

	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnAttributes(true)
	query.SetReturnData(true)

	results, err := keychain.QueryItem(query)
	if err != nil {
		fmt.Print(err)
		return "", "", err
	}

	if len(results) == 0 {
		return "", "", ErrCredentialsNotFound
	}

	return results[0].Account, string(results[0].Data), nil
}

// https://github.com/Versent/saml2aws/blob/master/helper/osxkeychain/osxkeychain.go
func splitServer3(serverURL string, item keychain.Item) (err error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return
	}

	hostAndPort := strings.Split(u.Host, ":")
	SetServer(item, hostAndPort[0])
	if len(hostAndPort) == 2 {
		SetPort(item, hostAndPort[1])
	}

	SetProtocol(item, u.Scheme)
	SetPath(item, u.Path)

	return
}
