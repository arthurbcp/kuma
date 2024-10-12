package domain

type Module struct {
	Id          string   `yaml:"id"`
	Description string   `yaml:"description"`
	Version     string   `yaml:"version"`
	Runs        []string `yaml:"runs"`
}

func NewModule(module map[string]interface{}, runs map[string]Run) Module {
	moduleRuns := []string{}
	for key, run := range runs {
		moduleRuns = append(moduleRuns, run.File+"/"+key)
	}
	return Module{
		Id:          module["id"].(string),
		Description: module["description"].(string),
		Version:     module["version"].(string),
		Runs:        moduleRuns,
	}
}
