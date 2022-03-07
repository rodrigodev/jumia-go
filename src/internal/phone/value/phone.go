package value

type Phone struct {
	Country     string `json:"country"`
	State       string `json:"state"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
}

type Customer struct {
	Phone string `json:"phone"`
}
