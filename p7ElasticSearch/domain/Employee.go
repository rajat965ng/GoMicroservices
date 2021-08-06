package domain

type Employee struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Designation string `json:"designation"`
}
