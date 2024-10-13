package services

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/arthurbcp/kuma/cmd/shared"
	"github.com/arthurbcp/kuma/internal/domain"
	"github.com/arthurbcp/kuma/internal/helpers"
	"github.com/arthurbcp/kuma/pkg/filesystem"
	"gopkg.in/yaml.v3"
)

// ModuleService handles operations related to Kuma modules, such as adding,
// removing, retrieving, and listing modules. It interacts with the filesystem
// to persist module configurations.
type ModuleService struct {
	path string                         // Base path where modules are stored
	fs   filesystem.FileSystemInterface // Filesystem interface for file operations
}

// NewModuleService creates and returns a new instance of ModuleService.
//
// Parameters:
// - path: The base directory path where modules are located.
// - fs: An implementation of the FileSystemInterface for file operations.
//
// Returns:
// - A pointer to the initialized ModuleService.
func NewModuleService(path string, fs filesystem.FileSystemInterface) *ModuleService {
	return &ModuleService{
		path: path,
		fs:   fs,
	}
}

// Add adds a new module to the Kuma modules configuration file.
// It ensures the modules file exists, retrieves the module details,
// converts the module to a map, updates the modules configuration, and
// writes the updated configuration back to the file.
//
// Parameters:
// - newModule: The name of the module to be added.
//
// Returns:
// - An error if any step fails; otherwise, nil.
func (s *ModuleService) Add(newModule string) error {
	modulesFile := shared.KumaFilesPath + "/kuma-modules.yaml"

	_, err := s.fs.CreateFileIfNotExists(modulesFile)
	if err != nil {
		return err
	}

	modules, err := helpers.UnmarshalFile(modulesFile, s.fs)
	if err != nil {
		return err
	}

	module, err := s.Get(newModule)
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
	s.fs.WriteFile(modulesFile, string(yamlContent))
	return nil
}

// Remove deletes a module from the Kuma modules configuration file.
// It unmarshals the current configuration, removes the specified module,
// and writes the updated configuration back. If no modules remain, it
// clears the modules file.
//
// Parameters:
// - module: The name of the module to be removed.
//
// Returns:
// - An error if any step fails; otherwise, nil.
func (s *ModuleService) Remove(module string) error {
	modulesFile := shared.KumaFilesPath + "/kuma-modules.yaml"
	modules, err := helpers.UnmarshalFile(modulesFile, s.fs)
	if err != nil {
		return err
	}

	delete(modules, module)

	if len(modules) == 0 {
		s.fs.WriteFile(modulesFile, "")
		return nil
	}

	yamlContent, err := yaml.Marshal(modules)
	if err != nil {
		return err
	}
	s.fs.WriteFile(modulesFile, string(yamlContent))
	return nil
}

// Get retrieves a specific module by its name.
// It reads the module's configuration file and associated runs,
// constructs a Module domain object, and returns it.
//
// Parameters:
// - module: The name of the module to retrieve.
//
// Returns:
// - A Module object containing the module's configuration and runs.
// - An error if any step fails.
func (s *ModuleService) Get(module string) (domain.Module, error) {
	configPath := fmt.Sprintf("%s/%s/kuma-config.yaml", s.path, module)
	configData, err := helpers.UnmarshalFile(configPath, s.fs)
	if err != nil {
		return domain.Module{}, err
	}

	runsService := NewRunService(fmt.Sprintf("%s/%s/%s", s.path, module, shared.KumaRunsPath), s.fs)
	runs, err := runsService.GetAll()
	if err != nil {
		return domain.Module{}, err
	}
	return domain.NewModule(configData, runs), nil
}

// GetAll retrieves all modules from the Kuma modules configuration file.
// It unmarshals the modules file, converts each module to the Module domain type,
// and returns a map of module names to Module objects.
//
// Returns:
// - A map where keys are module names and values are Module objects.
// - An error if any step fails.
func (s *ModuleService) GetAll() (map[string]domain.Module, error) {
	modulesFile := shared.KumaFilesPath + "/kuma-modules.yaml"
	modules, err := helpers.UnmarshalFile(modulesFile, s.fs)
	if err != nil {
		return nil, err
	}
	modulesMap := map[string]domain.Module{}
	for key, module := range modules {
		moduleString, err := json.Marshal(module)
		if err != nil {
			return nil, err
		}
		m := &domain.Module{}
		err = json.Unmarshal(moduleString, m)
		if err != nil {
			return nil, err
		}
		modulesMap[key] = *m
	}
	return modulesMap, nil
}

// GetModuleName extracts and returns the module name from a repository string.
// The repository string is expected to be in the format "owner/module".
//
// Parameters:
// - repo: The repository string containing the module name.
//
// Returns:
// - The extracted module name.
func (s *ModuleService) GetModuleName(repo string) string {
	splitRepo := strings.Split(repo, "/")
	if len(splitRepo) < 2 {
		return ""
	}
	return splitRepo[1]
}

// GetRun retrieves a specific run associated with a module.
// It reads the run's configuration file, verifies the run exists,
// constructs a Run domain object, and returns it.
//
// Parameters:
// - module: A pointer to the Module object containing run information.
// - runKey: The key identifying the specific run to retrieve.
// - modulePath: The filesystem path to the module's directory.
//
// Returns:
// - A pointer to the Run object containing run details.
// - An error if any step fails or the run is not found.
func (s *ModuleService) GetRun(module *domain.Module, runKey string, modulePath string) (*domain.Run, error) {
	moduleRun, exists := module.Runs[runKey]
	if !exists {
		return nil, fmt.Errorf("run key '%s' does not exist in the module", runKey)
	}

	runFilePath := fmt.Sprintf("%s/%s", modulePath, moduleRun.File)
	runs, err := helpers.UnmarshalFile(runFilePath, s.fs)
	if err != nil {
		return nil, err
	}

	runContent, ok := runs[runKey]
	if !ok {
		return nil, fmt.Errorf("run not found: %s", runKey)
	}

	run := domain.NewRun(
		runKey,
		runContent.(map[string]interface{})["description"].(string),
		runContent.(map[string]interface{})["steps"].([]interface{}),
		moduleRun.File,
	)

	return &run, nil
}
