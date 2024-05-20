package models

type BalanceV2 struct {
	ClearedBalance        SignedCurrencyAndAmount `json:"clearedBalance"`
	EffecticeBalance      SignedCurrencyAndAmount `json:"effectiveBalance"`
	PendingTransactions   SignedCurrencyAndAmount `json:"pendingTransactions"`
	AcceptedOverdraft     SignedCurrencyAndAmount `json:"acceptedOverdraft"`
	Amount                SignedCurrencyAndAmount `json:"amount"`
	TotalClearedBalance   SignedCurrencyAndAmount `json:"totalClearedBalance"`
	TotalEffectiveBalance SignedCurrencyAndAmount `json:"totalEffectiveBalance"`
}
