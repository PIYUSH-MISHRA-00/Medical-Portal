package models

type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Condition string `json:"condition"`
}