package models

type AccountIdentifiers struct {
	AccountIdentifier  string              `json:"accountIdentifier"`
	BankIdentifier     string              `json:"bankIdentifier"`
	IBAN               string              `json:"iban"`
	BIC                string              `json:"bic"`
	AccountIdentifiers []AccountIdentifier `json:"accountIdentifiers"`
}
