package domain

type Run struct {
	Description string        `json:"description"`
	Steps       []interface{} `json:"steps"`
	File        string        `json:"file"`
}

func NewRun(description string, steps []interface{}, file string) Run {
	return Run{
		Description: description,
		Steps:       steps,
		File:        file,
	}
}
