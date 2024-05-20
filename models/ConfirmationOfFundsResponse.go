package models

type ConfirmationOfFundsResponse struct {
	RequestedAmountAvailableToSpend                 bool `json:"requestedAmountAvailableToSpend"`
	AccountWouldBeInOverdraftIfRequestedAmountSpent bool `json:"accountWouldBeInOverdraftIfRequestedAmountSpent"`
}
