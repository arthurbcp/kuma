package services

import (
	"fmt"

	"github.com/arthurbcp/kuma/internal/domain"
	"github.com/arthurbcp/kuma/internal/helpers"
	"github.com/arthurbcp/kuma/pkg/filesystem"
)

// RunService handles operations related to Kuma runs, such as retrieving
// all runs or fetching a specific run by name. It interacts with the
// filesystem to read run configurations stored in YAML files.
type RunService struct {
	path string                         // Base path where runs are stored
	fs   filesystem.FileSystemInterface // Filesystem interface for file operations
}

// NewRunService creates and returns a new instance of RunService.
//
// Parameters:
// - path: The base directory path where runs are located.
// - fs: An implementation of the FileSystemInterface for file operations.
//
// Returns:
// - A pointer to the initialized RunService.
func NewRunService(path string, fs filesystem.FileSystemInterface) *RunService {
	return &RunService{
		path: path,
		fs:   fs,
	}
}

// GetAll retrieves all runs from the specified runs directory.
// It reads each run's YAML configuration file, unmarshals the data,
// and constructs a map of run names to Run domain objects.
//
// Returns:
// - A map where keys are run names and values are Run objects.
// - An error if any step fails, such as reading the directory or unmarshaling files.
func (s *RunService) GetAll() (map[string]domain.Run, error) {
	deprecateRunsFileMsg := "\nif you're using a runs.yaml file, please move it to the runs folder"
	files, err := s.fs.ReadDir(s.path)
	if err != nil {
		return nil, fmt.Errorf("error reading runs directory: %w%s", err, deprecateRunsFileMsg)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no runs found in %s%s", s.path, deprecateRunsFileMsg)
	}
	runs := make(map[string]domain.Run)
	for _, fileName := range files {
		data, err := helpers.UnmarshalFile(s.path+"/"+fileName, s.fs)
		if err != nil {
			return nil, err
		}
		for key, run := range data {
			if _, ok := runs[key]; ok {
				return nil, fmt.Errorf("conflict between runs found for the run %s\n rename one of them and try again", key)
			}
			runs[key] = domain.NewRun(
				key,
				run.(map[string]interface{})["description"].(string),
				run.(map[string]interface{})["steps"].([]interface{}),
				fileName,
			)
		}
	}

	return runs, nil
}

// Get retrieves a specific run by its name.
// It first fetches all runs using GetAll and then searches for the run
// with the specified name.
//
// Parameters:
// - name: The name of the run to retrieve.
//
// Returns:
// - The Run object corresponding to the provided name.
// - An error if the run is not found or if fetching all runs fails.
func (s *RunService) Get(name string) (domain.Run, error) {
	runs, err := s.GetAll()
	if err != nil {
		return domain.Run{}, err
	}
	run, ok := runs[name]
	if !ok {
		return domain.Run{}, fmt.Errorf("run not found: %s", name)
	}
	return run, nil
}
