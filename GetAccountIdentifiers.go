package myna

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AccountIdentifier struct {
	IdentifierType    string `json:"identifierType"`
	BankIdentifier    string `json:"bankIdentifier"`
	AccountIdentifier string `json:"accountIdentifier"`
}

type AccountIdentifiers struct {
	AccountIdentifier  string              `json:"accountIdentifier"`
	BankIdentifier     string              `json:"bankIdentifier"`
	IBAN               string              `json:"iban"`
	BIC                string              `json:"bic"`
	AccountIdentifiers []AccountIdentifier `json:"accountIdentifiers"`
}

func (c *Client) GetAccountIdentifiers(a *AccountV2) (AccountIdentifiers, error) {
	var ret AccountIdentifiers
	url := fmt.Sprintf("%s/%s%s", ACCOUNTS_URL_PROD, a.AccountUID, IDENTIFIERS_ENDPOINT)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ret, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.bearerToken)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return ret, err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return ret, err
	}
	// fmt.Printf("%s\n", buf)
	err = json.Unmarshal(buf, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
