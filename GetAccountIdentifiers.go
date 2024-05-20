package myna

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chickazama/myna/models"
)

func (c *Client) GetAccountIdentifiers(a *models.AccountV2) (models.AccountIdentifiers, error) {
	var ret models.AccountIdentifiers
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
