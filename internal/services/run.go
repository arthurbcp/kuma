package services

import (
	"fmt"

	"github.com/arthurbcp/kuma/internal/domain"
	"github.com/arthurbcp/kuma/internal/helpers"
	"github.com/arthurbcp/kuma/pkg/filesystem"
)

type RunService struct {
	path string
	fs   filesystem.FileSystemInterface
}

func NewRunService(path string, fs filesystem.FileSystemInterface) *RunService {
	return &RunService{
		path: path,
		fs:   fs,
	}
}

func (s *RunService) GetAll() (map[string]domain.Run, error) {
	deprecateRunsFileMsg := "\nif your a using runs.yaml file, please move it to the runs folder"
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
			runs[key] = domain.NewRun(
				run.(map[string]interface{})["description"].(string),
				run.(map[string]interface{})["steps"].([]interface{}),
			)
		}
	}

	return runs, nil
}

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
