package myna

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chickazama/myna/models"
)

func (c *Client) GetAccountHolderName() (models.AccountHolderName, error) {
	var ret models.AccountHolderName
	url := fmt.Sprintf("%s%s/name", BASE_URL_PROD, ACCOUNT_HOLDER_ENDPOINT)
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
