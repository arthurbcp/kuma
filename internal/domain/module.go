package domain

type ModuleRun struct {
	Description string `json:"description"`
	File        string `json:"file"`
}

type Module struct {
	Key         string               `json:"key"`
	Description string               `json:"description"`
	Version     string               `json:"version"`
	Runs        map[string]ModuleRun `json:"runs"`
}

func NewModule(key string, module map[string]interface{}, runs map[string]Run) Module {
	runsMap := map[string]ModuleRun{}
	for key, run := range runs {
		runsMap[key] = ModuleRun{
			Description: run.Description,
			File:        run.File,
		}
	}
	return Module{
		Key:         key,
		Description: module["description"].(string),
		Version:     module["version"].(string),
		Runs:        runsMap,
	}
}
