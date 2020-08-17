package Entity

type Account struct {
	Owner              string `json:"owner"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Balance            int    `json:"balance"`
	CurrencyType       string `json:"currencyType"`
	ZeroBalanceAccount bool   `json:"zeroBalanceAccount"`
}
