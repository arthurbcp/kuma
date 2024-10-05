package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	testCases := []struct {
		name          string
		projectPath   string
		templatesPath string
	}{
		{
			name:          "Valid paths",
			projectPath:   "/path/to/project",
			templatesPath: "/path/to/templates",
		},
		{
			name:          "Empty paths",
			projectPath:   "",
			templatesPath: "",
		},
		{
			name:          "Relative paths",
			projectPath:   "./project",
			templatesPath: "../templates",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := NewConfig(tc.projectPath, tc.templatesPath)

			assert.NotNil(t, config, "Config should not be nil")
			assert.Equal(t, tc.projectPath, config.ProjectPath, "ProjectPath should match input")
			assert.Equal(t, tc.templatesPath, config.TemplatesPath, "TemplatesPath should match input")
		})
	}
}
