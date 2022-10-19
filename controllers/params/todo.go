package params

type TodoCreate struct {
	ToList      string `json:"to_list"`
	Description string `json:"description"`
}

type TodoUpdate struct {
	ToList      string `json:"to_list"`
	Description string `json:"description"`
}
