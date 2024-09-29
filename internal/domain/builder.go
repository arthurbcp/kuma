package domain

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// BuilderData encapsulates the structure and templates data parsed from configuration files.
type BuilderData struct {
	// Structure defines the directory and file hierarchy to be created.
	Structure map[string]interface{}

	// Templates defines the templates to be applied to the generated files.
	Templates map[string]interface{}
}

// Builder is responsible for managing the configuration and data required to build the project structure.
type Builder struct {
	// Config holds the configuration paths for the project and templates.
	Config *Config

	// Data holds the parsed structure and templates data.
	Data *BuilderData
}

// NewBuilder initializes a new Builder instance.
//
// Parameters:
//   - file: The path to the configuration file (JSON or YAML).
//   - vars: A map of variables to replace placeholders in the configuration file.
//   - config: A pointer to the Config struct containing project and templates paths.
//
// Returns:
//
//	A pointer to a Builder instance if successful, or an error if initialization fails.
func NewBuilder(file string, vars map[string]interface{}, config *Config) (*Builder, error) {
	builder := Builder{}
	err := builder.SetBuilderData(file, vars)
	if err != nil {
		return nil, err
	}
	err = builder.SetConfig(config)
	if err != nil {
		return nil, err
	}
	return &builder, nil
}

// SetBuilderData parses the configuration file and populates the BuilderData.
//
// Parameters:
//   - file: The path to the configuration file.
//   - vars: A map of variables for placeholder replacement in the configuration.
//
// Returns:
//
//	An error if parsing fails, otherwise nil.
func (b *Builder) SetBuilderData(file string, vars map[string]interface{}) error {
	helpers.HeaderPrint("PARSING CONFIG")

	// Read the content of the configuration file.
	configData, err := helpers.ReadFile(file)
	if err != nil {
		return err
	}

	// Replace variables in the configuration data.
	configData, err = helpers.ReplaceVars(configData, vars, helpers.FuncMap)
	if err != nil {
		return err
	}

	// Determine the file type based on its extension and unmarshal accordingly.
	switch filepath.Ext(file) {
	case ".yaml", ".yml":
		data, err := UnmarshalYamlConfig([]byte(configData))
		if err != nil {
			return err
		}
		b.Data = data
	case ".json":
		data, err := UnmarshalJsonConfig([]byte(configData))
		if err != nil {
			return err
		}
		b.Data = data
	default:
		return fmt.Errorf("invalid file extension: %s", file)
	}
	return nil
}

// SetConfig assigns the provided Config to the Builder.
//
// Parameters:
//   - config: A pointer to the Config struct.
//
// Returns:
//
//	An error if setting the configuration fails, otherwise nil.
func (b *Builder) SetConfig(config *Config) error {
	b.Config = config
	return nil
}

// UnmarshalJsonConfig parses JSON configuration data into BuilderData.
//
// Parameters:
//   - configData: A byte slice containing JSON-formatted configuration data.
//
// Returns:
//
//	A pointer to BuilderData and an error if unmarshaling fails.
func UnmarshalJsonConfig(configData []byte) (*BuilderData, error) {
	config := BuilderData{}
	err := json.Unmarshal(configData, &config)
	if err != nil {
		return &config, err
	}
	// Note: The original code does not populate BuilderData from 'c'.
	return &config, nil
}

// UnmarshalYamlConfig parses YAML configuration data into BuilderData.
//
// Parameters:
//   - configData: A byte slice containing YAML-formatted configuration data.
//
// Returns:
//
//	A pointer to BuilderData and an error if unmarshaling fails.
func UnmarshalYamlConfig(configData []byte) (*BuilderData, error) {
	config := BuilderData{}
	c := map[interface{}]interface{}{}
	err := yaml.Unmarshal(configData, &c)
	if err != nil {
		return &config, err
	}
	// Decode the map into BuilderData using mapstructure.
	err = mapstructure.Decode(c, &config)
	if err != nil {
		return &config, err
	}
	return &config, nil
}
