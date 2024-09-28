package domain

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type BuilderData struct {
	Structure map[string]interface{}
	Templates map[string]interface{}
}

type Builder struct {
	Config *Config
	Data   *BuilderData
}

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

func (b *Builder) SetBuilderData(file string, vars map[string]interface{}) error {
	helpers.HeaderPrint("PARSING CONFIG")

	configData, err := helpers.ReadFile(file)
	if err != nil {
		return err
	}
	configData, err = helpers.ReplaceVars(configData, vars, helpers.FuncMap)
	if err != nil {
		return err
	}
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

func (b *Builder) SetConfig(config *Config) error {
	b.Config = config
	return nil
}

func UnmarshalJsonConfig(configData []byte) (*BuilderData, error) {
	config := BuilderData{}
	c := map[interface{}]interface{}{}
	err := json.Unmarshal(configData, &c)
	if err != nil {
		return &config, err
	}
	return &config, nil
}

func UnmarshalYamlConfig(configData []byte) (*BuilderData, error) {
	config := BuilderData{}
	c := map[interface{}]interface{}{}
	err := yaml.Unmarshal(configData, &c)
	if err != nil {
		return &config, err
	}
	err = mapstructure.Decode(c, &config)
	if err != nil {
		return &config, err
	}
	return &config, nil
}
