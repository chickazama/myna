package models

type WrapperAccountV2 struct {
	Accounts []AccountV2 `json:"accounts"`
}
