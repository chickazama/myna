package myna

import "net/http"

type Client struct {
	bearerToken string
	httpClient  *http.Client
}

func New(bearerToken string, httpClient *http.Client) *Client {
	ret := new(Client)
	ret.bearerToken = bearerToken
	ret.httpClient = httpClient
	return ret
}
