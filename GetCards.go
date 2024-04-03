package myna

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CurrencyFlag struct {
	Enabled  bool   `json:"enabled"`
	Currency string `json:"currency"`
}

type cardsWrapper struct {
	Cards []Card `json:"cards"`
}

type Card struct {
	CardUid                   string         `json:"cardUid"`
	PublicToken               string         `json:"publicToken"`
	Enabled                   bool           `json:"enabled"`
	WalletNotificationEnabled bool           `json:"walletNotificationEnabled"`
	POSEnabled                bool           `json:"posEnabled"`
	ATMEnabled                bool           `json:"atmEnabled"`
	OnlineEnabled             bool           `json:"onlineEnabled"`
	MobileWalletEnabled       bool           `json:"mobileWalletEnabled"`
	GamblingEnabled           bool           `json:"gamblingEnabled"`
	MagStripeEnabled          bool           `json:"magStripeEnabled"`
	Cancelled                 bool           `json:"cancelled"`
	ActivationRequested       bool           `json:"activationRequested"`
	Activated                 bool           `json:"activated"`
	EndOfCardNumber           string         `json:"endOfCardNumber"`
	CurrencyFlags             []CurrencyFlag `json:"currencyFlag"`
	CardAssociationUid        string         `json:"cardAssociationUid"`
	GamblingToBeEnabledAt     string         `json:"gamblingToBeEnabledAt"`
}

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
	var wrapper cardsWrapper
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.Cards
	return ret, nil
}
