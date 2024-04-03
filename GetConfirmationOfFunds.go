package myna

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ConfirmationOfFundsResponse struct {
	RequestedAmountAvailableToSpend                 bool `json:"requestedAmountAvailableToSpend"`
	AccountWouldBeInOverdraftIfRequestedAmountSpent bool `json:"accountWouldBeInOverdraftIfRequestedAmountSpent"`
}

func (c *Client) GetConfirmationOfFunds(a *AccountV2, amountMinorUnits int64) (ConfirmationOfFundsResponse, error) {
	var ret ConfirmationOfFundsResponse
	url := fmt.Sprintf("%s/%s%s?accountUid=%s&targetAmountInMinorUnits=%d", ACCOUNTS_URL_PROD, a.AccountUID, CONFIRMATION_OF_FUNDS_ENDPOINT, a.AccountUID, amountMinorUnits)
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
	fmt.Printf("%s\n", buf)
	err = json.Unmarshal(buf, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
