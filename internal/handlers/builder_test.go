package handlers

import (
	"testing"

	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name            string
		templateContent string
		includeContent  string
		structure       map[string]interface{}
		data            map[string]interface{}
		expectedFile    string
		expectedContent string
		expectedError   string
	}{
		{
			name:            "Template with includes",
			includeContent:  `{{define "include"}}{{ .include}}{{end}}`,
			templateContent: `{{.Data.key}} {{ block "include" .Data }}{{end}}`,
			structure: map[string]interface{}{
				"dir": map[string]interface{}{
					"file.txt": map[string]interface{}{
						"Template": "template.txt",
						"Data": map[string]interface{}{
							"key":     "value",
							"include": "include value",
						},
						"Includes": []interface{}{"include.txt"},
					},
				},
			},
			data:            map[string]interface{}{"key": "value"},
			expectedFile:    "project/dir/file.txt",
			expectedContent: "value include value",
		},
		{
			name:            "Missing template",
			templateContent: "",
			includeContent:  "",
			structure: map[string]interface{}{
				"file.txt": map[string]interface{}{
					"Template": "non_existent.txt",
					"Data":     map[string]interface{}{},
				},
			},
			expectedError: "open templates/non_existent.txt: file does not exist",
		},
		{
			name:            "Missing include file",
			templateContent: `{{.Data.key}} {{ block "include" .Data }}{{end}}`,
			includeContent:  "",
			structure: map[string]interface{}{
				"file.txt": map[string]interface{}{
					"Template": "template.txt",
					"Data":     map[string]interface{}{"key": "value"},
					"Includes": []interface{}{"non_existent.txt"},
				},
			},
			expectedError: "open templates/non_existent.txt: file does not exist",
		},
		{
			name:            "Template required",
			templateContent: "",
			includeContent:  "",
			structure: map[string]interface{}{
				"file.txt": map[string]interface{}{
					"Data": map[string]interface{}{},
				},
			},
			expectedError: "template is required",
		},
		{
			name:            "Invalid include type",
			templateContent: `{{.Data.key}}`,
			includeContent:  "",
			structure: map[string]interface{}{
				"file.txt": map[string]interface{}{
					"Template": "template.txt",
					"Data":     map[string]interface{}{"key": "value"},
					"Includes": []interface{}{123}, // Invalid include type
				},
			},
			expectedError: "invalid include type: 123",
		},
		{
			name:            "Invalid template syntax",
			templateContent: "{{.Data.key} {{ block include .Data }}{{end}}",
			includeContent:  "",
			structure: map[string]interface{}{
				"file.txt": map[string]interface{}{
					"Template": "template.txt",
					"Data":     map[string]interface{}{"key": "value"},
				},
			},
			expectedError: "error parsing template file templates/template.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aferoFs := afero.NewMemMapFs()

			// Create the templates directory
			err := aferoFs.MkdirAll("templates", 0755)
			assert.NoError(t, err)

			// Create and write the template file
			if tt.templateContent != "" {
				err = afero.WriteFile(aferoFs, "templates/template.txt", []byte(tt.templateContent), 0644)
				assert.NoError(t, err)
			}

			// Create and write the include file
			if tt.includeContent != "" {
				err = afero.WriteFile(aferoFs, "templates/include.txt", []byte(tt.includeContent), 0644)
				assert.NoError(t, err)
			}

			fs := filesystem.NewFileSystem(aferoFs)
			h := helpers.NewHelpers()

			builder := &domain.Builder{
				Fs:      fs,
				Helpers: h,
				Data: &domain.BuilderData{
					Structure: tt.structure,
				},
				Config: &domain.Config{
					ProjectPath:   "project",
					TemplatesPath: "templates",
				},
			}

			handler := NewBuilderHandler(builder)

			// Test build
			err = handler.Build()

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)

				// Verify that the file was created
				exists, err := afero.Exists(aferoFs, tt.expectedFile)
				assert.NoError(t, err)
				assert.True(t, exists)

				// Check the content of the created file
				content, err := afero.ReadFile(aferoFs, tt.expectedFile)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedContent, string(content))
			}
		})
	}
}
