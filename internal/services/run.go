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
	data, err := helpers.UnmarshalFile(s.path, s.fs)
	if err != nil {
		return nil, err
	}
	var runs = make(map[string]domain.Run)
	for key, run := range data {
		runs[key] = domain.NewRun(
			run.(map[string]interface{})["description"].(string),
			run.(map[string]interface{})["steps"].([]interface{}),
		)
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
