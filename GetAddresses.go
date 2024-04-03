package myna

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AddressV2 struct {
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Line3       string `json:"line3"`
	PostTown    string `json:"postTown"`
	PostCode    string `json:"postCode"`
	CountryCode string `json:"countryCode"`
}

type AddressesV2 struct {
	Current  AddressV2   `json:"current"`
	Previous []AddressV2 `json:"previous"`
}

func (c *Client) GetAddresses() (AddressesV2, error) {
	var ret AddressesV2
	url := fmt.Sprintf("%s%s", BASE_URL_PROD, ADDRESSES_ENDPOINT)
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
