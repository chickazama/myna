package myna

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetCards() ([]Card, error) {
	var ret []Card
	url := fmt.Sprintf("%s%s", BASE_URL_PROD, "/cards")
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
	var wrapper WrapperCard
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.Cards
	return ret, nil
}
