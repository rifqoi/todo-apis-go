package models

type Todo struct {
	ID          int    `json:"id"`
	ToList      string `json:"to_list"`
	Description string `json:"description"`
}
