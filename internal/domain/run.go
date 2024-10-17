package domain

type Run struct {
	Key         string        `json:"key"`
	Description string        `json:"description"`
	Steps       []interface{} `json:"steps"`
	File        string        `json:"file"`
}

func NewRun(key string, description string, steps []interface{}, file string) Run {
	return Run{
		Key:         key,
		Description: description,
		Steps:       steps,
		File:        file,
	}
}
