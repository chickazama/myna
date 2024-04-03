package myna

import (
	"encoding/json"
	"io"
	"net/http"
)

type accountV2Wrapper struct {
	Accounts []AccountV2 `json:"accounts"`
}

func (c *Client) GetAccounts() ([]AccountV2, error) {
	var ret []AccountV2
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
	var wrapper accountV2Wrapper
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.Accounts
	return ret, nil
}
