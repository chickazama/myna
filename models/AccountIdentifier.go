package models

type AccountIdentifier struct {
	IdentifierType    string `json:"identifierType"`
	BankIdentifier    string `json:"bankIdentifier"`
	AccountIdentifier string `json:"accountIdentifier"`
}
