package params

type TodoCreate struct {
	Task        string `json:"task" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsCompleted bool   `json:"is_completed"`
}
