package handlers

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/arthurbcp/kuma/internal/domain"
	"github.com/arthurbcp/kuma/internal/helpers"
	"github.com/arthurbcp/kuma/pkg/style"
	"github.com/spf13/afero"
)

// BuilderHandler manages the building process of the project structure.
// It interacts with the domain.Builder to create directories and files based on templates.
type BuilderHandler struct {
	// builder is the domain Builder responsible for providing structure and template data.
	builder *domain.Builder
}

// NewBuilderHandler creates and returns a new BuilderHandler instance.
//
// Parameters:
//   - builder: A pointer to the domain.Builder instance.
//
// Returns:
//
//	A pointer to a BuilderHandler.
func NewBuilderHandler(builder *domain.Builder) *BuilderHandler {
	return &BuilderHandler{
		builder: builder,
	}
}

// Build initiates the building process by applying templates and creating
// directories and files as defined in the Builder's data.
//
// Returns:
//
//	An error if the build process fails, otherwise nil.
func (h *BuilderHandler) Build() error {
	style.LogPrint("applying templates...")
	// Start recursive creation of directories and files from the root.
	err := h.createDirAndFilesRecursive("", h.builder.Data.Structure, h.builder.Config.ProjectPath)
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
}

// createDirAndFilesRecursive recursively creates directories and files based on the provided structure.
//
// Parameters:
//   - key: The current directory or file name.
//   - node: The nested structure (directories or file definitions).
//   - basePath: The accumulated file system path from previous recursion levels.
//
// Returns:
//
//	An error if directory or file creation fails, otherwise nil.
func (h *BuilderHandler) createDirAndFilesRecursive(key string, node interface{}, basePath string) error {
	// Construct the current path by joining the base path with the current key.
	currentPath := filepath.Join(basePath, key)

	// Create the directory if it does not exist.
	err := h.builder.Fs.CreateDirectoryIfNotExists(currentPath)
	if err != nil {
		return err
	}

	switch children := node.(type) {
	case map[string]interface{}:
		// Iterate through the map to handle subdirectories and files.
		for childKey, childValue := range children {
			// Check if the key contains a dot, indicating a file name.
			if len(strings.Split(childKey, ".")) > 1 {
				// Create the file and apply the corresponding template.
				err := h.createFileAndApplyTemplate(currentPath, childKey, childValue.(map[string]interface{}))
				if err != nil {
					style.CrossMarkPrint(filepath.Join(currentPath, childKey))
					return err
				}
				style.CheckMarkPrint(filepath.Join(currentPath, childKey))
				continue
			}

			// Replace variables in the directory name.
			childKey, err := helpers.ReplaceVars(childKey, childValue, helpers.GetFuncMap())
			if err != nil {
				return err
			}

			// Recursively create subdirectories and files.
			err = h.createDirAndFilesRecursive(childKey, childValue, currentPath)
			if err != nil {
				return err
			}
		}
	default:
		// If the node does not have subdirectories or files, do nothing.
	}

	return nil
}

// createFileAndApplyTemplate creates a file and applies the specified template to it.
//
// Parameters:
//   - currentPath: The directory path where the file will be created.
//   - fileName: The name of the file to be created.
//   - data: A map containing template data and metadata.
//
// Returns:
//
//	An error if file creation or template application fails, otherwise nil.
func (h *BuilderHandler) createFileAndApplyTemplate(currentPath string, fileName string, data map[string]interface{}) error {
	// Construct the full file path.
	filePath := filepath.Join(currentPath, fileName)

	// Create the file.
	file, err := h.builder.Fs.CreateFile(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Retrieve and parse the template.
	t, err := h.getTemplate(data)
	if err != nil {
		return err
	}
	// Prepare the data for template execution.
	data = map[string]interface{}{
		"data":   data["data"],
		"global": h.builder.Data.Global,
	}
	// Execute the template and write to the file.
	return t.Execute(file, data)
}

// getTemplate retrieves and parses the template files based on the provided data.
//
// Parameters:
//   - data: A map containing template metadata, including the template name and any includes.
//
// Returns:
//
//	A pointer to the parsed template.Template and an error if parsing fails.
func (h *BuilderHandler) getTemplate(data map[string]interface{}) (*template.Template, error) {
	// Extract the main template name from the data.
	templateName, ok := data["template"].(string)
	if !ok || templateName == "" {
		return nil, fmt.Errorf("template is required")
	}

	// Construct the path to the main template file.
	templateFile := filepath.Join(h.builder.Config.TemplatesPath, templateName)

	// Initialize a slice to hold all template file paths, including any includes.
	allTemplates := []string{templateFile}

	// If there are additional templates to include, add their paths.
	if includes, ok := data["includes"].([]interface{}); ok {
		for _, include := range includes {
			includeStr, ok := include.(string)
			if !ok {
				return nil, fmt.Errorf("invalid include type: %v", include)
			}
			allTemplates = append(allTemplates, filepath.Join(h.builder.Config.TemplatesPath, includeStr))
		}
	}

	// Create a new template object
	tmpl := template.New(templateName).Funcs(helpers.GetFuncMap())

	// Iterate through all the template paths, read them from the afero filesystem, and parse them
	for _, tmplFile := range allTemplates {
		// Read the template file using afero
		content, err := afero.ReadFile(h.builder.Fs.GetAferoFs(), tmplFile)
		if err != nil {
			return nil, fmt.Errorf("error reading template file %s: %w", tmplFile, err)
		}

		// Parse the template content
		_, err = tmpl.Parse(string(content))
		if err != nil {
			return nil, fmt.Errorf("error parsing template file %s: %w", tmplFile, err)
		}
	}

	// Return the parsed template
	return tmpl, nil
}
