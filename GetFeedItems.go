package myna

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetFeedItems(a *AccountV2, t0 string) ([]FeedItem, error) {
	var ret []FeedItem
	url := fmt.Sprintf("%s/feed/account/%s/category/%s?accountUid=%s&categoryUid=%s&changesSince=%s", BASE_URL_PROD, a.AccountUID, a.DefaultCategory, a.AccountUID, a.DefaultCategory, t0)
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
	var wrapper WrapperFeedItem
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.FeedItems
	return ret, nil
}
