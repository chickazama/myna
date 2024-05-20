package myna

type AddressesV2 struct {
	Current  AddressV2   `json:"current"`
	Previous []AddressV2 `json:"previous"`
}
