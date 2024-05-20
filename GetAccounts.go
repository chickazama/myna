package myna

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/chickazama/myna/models"
)

func (c *Client) GetAccounts() ([]models.AccountV2, error) {
	var ret []models.AccountV2
	req, err := http.NewRequest(http.MethodGet, ACCOUNTS_URL_PROD, nil)
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
	var wrapper models.WrapperAccountV2
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.Accounts
	return ret, nil
}
