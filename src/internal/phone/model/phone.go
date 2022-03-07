package model

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (Customer) TableName() string {
	return "customer"
}
