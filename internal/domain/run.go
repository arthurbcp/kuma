package domain

type Run struct {
	Description string        `json:"description"`
	Steps       []interface{} `json:"steps"`
}

func NewRun(description string, steps []interface{}) Run {
	return Run{
		Description: description,
		Steps:       steps,
	}
}
