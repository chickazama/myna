package myna

type Card struct {
	CardUid                   string         `json:"cardUid"`
	PublicToken               string         `json:"publicToken"`
	Enabled                   bool           `json:"enabled"`
	WalletNotificationEnabled bool           `json:"walletNotificationEnabled"`
	POSEnabled                bool           `json:"posEnabled"`
	ATMEnabled                bool           `json:"atmEnabled"`
	OnlineEnabled             bool           `json:"onlineEnabled"`
	MobileWalletEnabled       bool           `json:"mobileWalletEnabled"`
	GamblingEnabled           bool           `json:"gamblingEnabled"`
	MagStripeEnabled          bool           `json:"magStripeEnabled"`
	Cancelled                 bool           `json:"cancelled"`
	ActivationRequested       bool           `json:"activationRequested"`
	Activated                 bool           `json:"activated"`
	EndOfCardNumber           string         `json:"endOfCardNumber"`
	CurrencyFlags             []CurrencyFlag `json:"currencyFlag"`
	CardAssociationUid        string         `json:"cardAssociationUid"`
	GamblingToBeEnabledAt     string         `json:"gamblingToBeEnabledAt"`
}
