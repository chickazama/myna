package myna

type AddressV2 struct {
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Line3       string `json:"line3"`
	PostTown    string `json:"postTown"`
	PostCode    string `json:"postCode"`
	CountryCode string `json:"countryCode"`
}
