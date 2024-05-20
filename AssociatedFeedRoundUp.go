package myna

type AssociatedFeedRoundUp struct {
	GoalCategoryUid string            `json:"goalCategoryUid"`
	Amount          CurrencyAndAmount `json:"amount"`
}
