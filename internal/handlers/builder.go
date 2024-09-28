package handlers

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
)

type BuilderHandler struct {
	builder *domain.Builder
	vars    map[string]interface{}
}

func NewBuilderHandler(builder *domain.Builder) *BuilderHandler {
	return &BuilderHandler{
		builder: builder,
	}
}

func (h *BuilderHandler) Build(vars map[string]interface{}) error {
	h.vars = vars
	helpers.HeaderPrint("APPLYING TEMPLATES")
	err := h.createDirAndFilesRecursive("", h.builder.Data.Structure, h.builder.Config.ProjectPath)
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
}

// createDirRecursive creates directories recursively based on the YAML structure
// key: current directory name
// node: nested directories or empty if no subdirectories
// basePath: accumulated path from previous recursion levels
func (h *BuilderHandler) createDirAndFilesRecursive(key string, node interface{}, basePath string) error {
	// Construct the current path by joining the base path with the current key
	currentPath := filepath.Join(basePath, key)
	// Create the directory with appropriate permissions
	err := helpers.CreateDirectoryIfNotExists(currentPath)
	if err != nil {
		return err
	}

	switch children := node.(type) {
	case map[string]interface{}:
		// If the node is a map with string keys, iterate and recurse
		for childKey, childValue := range children {
			// Check if the key contains a dot, indicating a file name
			if len(strings.Split(childKey, ".")) > 1 {
				err := h.createFileAndApplyTemplate(currentPath, childKey, childValue.(map[string]interface{}))
				if err != nil {
					helpers.CrossMarkPrint(filepath.Join(currentPath, childKey))
					return err
				}
				helpers.CheckMarkPrint(filepath.Join(currentPath, childKey))
				continue
			}
			childKey, err := helpers.ReplaceVars(childKey, childValue, helpers.FuncMap)
			if err != nil {
				return err
			}
			err = h.createDirAndFilesRecursive(childKey, childValue, currentPath)
			if err != nil {
				return err
			}
		}
	default:
		// If the node does not have subdirectories, do nothing
	}

	return nil
}

func (h *BuilderHandler) createFileAndApplyTemplate(currentPath string, fileName string, data map[string]interface{}) error {
	filePath := filepath.Join(currentPath, fileName)
	file, err := helpers.CreateFile(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	t, err := h.getTemplate(data)
	if err != nil {
		return err
	}
	data = map[string]interface{}{
		"Data": data,
		// TODO: Add Global Logic
	}
	return t.Execute(file, data)
}

func (h *BuilderHandler) getTemplate(data map[string]interface{}) (*template.Template, error) {
	templateName := data["Template"].(string)
	if templateName == "" {
		return nil, fmt.Errorf("template is required")
	}
	templateFile := h.builder.Config.TemplatesPath + "/" + templateName
	allTemplates := []string{templateFile}
	if includes, ok := data["Includes"].([]interface{}); ok {
		for _, include := range includes {
			allTemplates = append(allTemplates, filepath.Join(h.builder.Config.TemplatesPath, include.(string)))
		}
	}
	return template.New(templateName).Funcs(helpers.FuncMap).ParseFiles(allTemplates...)
}
