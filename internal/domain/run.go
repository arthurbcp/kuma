package domain

type Run struct {
	Key         string        `json:"key"`
	Description string        `json:"description"`
	Steps       []interface{} `json:"steps"`
	File        string        `json:"file"`
	Visible     bool          `json:"visible"`
}

func NewRun(key string, description string, steps []interface{}, file string, visible bool) Run {
	return Run{
		Key:         key,
		Description: description,
		Steps:       steps,
		Visible:     visible,
		File:        file,
	}
}
