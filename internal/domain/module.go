package domain

type Module struct {
	Description string   `json:"description"`
	Version     string   `json:"version"`
	Runs        []string `json:"runs"`
}

func NewModule(module map[string]interface{}, runs map[string]Run) Module {
	moduleRuns := []string{}
	for key, run := range runs {
		moduleRuns = append(moduleRuns, run.File+"/"+key)
	}
	return Module{
		Description: module["description"].(string),
		Version:     module["version"].(string),
		Runs:        moduleRuns,
	}
}
