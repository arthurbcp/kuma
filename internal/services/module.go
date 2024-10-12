package services

import (
	"strings"

	"github.com/arthurbcp/kuma/cmd/shared"
	"github.com/arthurbcp/kuma/internal/domain"
	"github.com/arthurbcp/kuma/internal/helpers"
	"github.com/arthurbcp/kuma/pkg/filesystem"
	"gopkg.in/yaml.v3"
)

type ModuleService struct {
	path string
	fs   filesystem.FileSystemInterface
}

func NewModuleService(path string, fs filesystem.FileSystemInterface) *ModuleService {
	return &ModuleService{
		path: path,
		fs:   fs,
	}
}

func (m *ModuleService) Add(newModule string) error {
	modulesFile := shared.KumaFilesPath + "/kuma-modules.yaml"
	_, err := m.fs.CreateFileIfNotExists(modulesFile)
	if err != nil {
		return err
	}
	modules, err := helpers.UnmarshalFile(modulesFile, m.fs)
	if err != nil {
		return err
	}

	module, err := m.GetModule(newModule)
	if err != nil {
		return err
	}

	mapModule, err := helpers.StructToMap(module)
	if err != nil {
		return err
	}
	modules[newModule] = mapModule

	yamlContent, err := yaml.Marshal(modules)
	if err != nil {
		return err
	}
	m.fs.WriteFile(modulesFile, string(yamlContent))
	return nil
}

func (m *ModuleService) GetModule(module string) (domain.Module, error) {
	configData, err := helpers.UnmarshalFile(m.path+"/"+module+"/kuma-config.yaml", m.fs)
	if err != nil {
		return domain.Module{}, err
	}
	runsService := NewRunService(m.path+"/"+module+"/"+shared.KumaRunsPath, m.fs)
	runs, err := runsService.GetAll()
	if err != nil {
		return domain.Module{}, err
	}
	return domain.NewModule(configData, runs), nil
}

func (m *ModuleService) GetModuleName(repo string) string {
	splitRepo := strings.Split(repo, "/")
	return splitRepo[1]
}
