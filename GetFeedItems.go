package myna

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AssociatedFeedRoundUp struct {
	GoalCategoryUid string            `json:"goalCategoryUid"`
	Amount          CurrencyAndAmount `json:"amount"`
}

type BatchPaymentDetails struct {
	BatchPaymentUID  string `json:"batchPaymentUid"`
	BatchPaymentType string `json:"batchPaymentType"`
}

type CurrencyAndAmount struct {
	Currency   string `json:"currency"`
	MinorUnits int64  `json:"minorUnits"`
}

type feedItemsWrapper struct {
	FeedItems []FeedItem `json:"feedItems"`
}

type FeedItem struct {
	FeedItemUID                        string                `json:"feedItemUid"`
	CategoryUID                        string                `json:"categoryUid"`
	Amount                             CurrencyAndAmount     `json:"amount"`
	SourceAmount                       CurrencyAndAmount     `json:"sourceAmount"`
	Direction                          string                `json:"direction"`
	UpdatedAt                          string                `json:"updatedAt"`
	TransactionTime                    string                `json:"transactionTime"`
	SettlementTime                     string                `json:"settlementTime"`
	RetryAllocationUntilTime           string                `json:"retryAllocationUntilTime"`
	Source                             string                `json:"source"`
	SourceSubType                      string                `json:"sourceSubType"`
	Status                             string                `json:"status"`
	TransactingApplicationUserUID      string                `json:"transactingApplicationUserUid"`
	CounterPartyType                   string                `json:"counterPartyType"`
	CounterPartyUID                    string                `json:"counterPartyUid"`
	CounterPartyName                   string                `json:"counterPartyName"`
	CounterPartySubEntityUID           string                `json:"counterPartySubEntityUid"`
	CounterPartySubEntityName          string                `json:"counterPartySubEntityName"`
	CounterPartySubEntityIdentifier    string                `json:"counterPartySubEntityIdentifier"`
	CounterPartySubEntitySubIdentifier string                `json:"counterPartySubEntitySubIdentifier"`
	ExchangeRate                       float64               `json:"exchangeRate"`
	TotalFees                          float64               `json:"totalFees"`
	TotalFeeAmount                     CurrencyAndAmount     `json:"totalFeeAmount"`
	Reference                          string                `json:"reference"`
	Country                            string                `json:"country"`
	SpendingCategory                   string                `json:"spendingCategory"`
	UserNote                           string                `json:"userNote"`
	RoundUp                            AssociatedFeedRoundUp `json:"roundUp"`
	HasAttachment                      bool                  `json:"hasAttachment"`
	HasReceipt                         bool                  `json:"hasReceipt"`
	BatchPaymentDetails                BatchPaymentDetails   `json:"batchPaymentDetails"`
}

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
	var wrapper feedItemsWrapper
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.FeedItems
	return ret, nil
}
