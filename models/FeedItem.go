package models

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
