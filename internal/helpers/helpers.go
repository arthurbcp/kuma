package helpers

import (
	"text/template"

	"github.com/arthurbcp/kuma-cli/pkg/filesystem"
)

//go:generate mockgen -source=helpers.go -destination=../mocks/helpers.go -package=mocks
type HelpersInterface interface {
	StringContains(s []string, e string) bool
	InterfaceContains(s []interface{}, e string) bool
	HeaderPrint(text string)
	CheckMarkPrint(text string)
	CrossMarkPrint(text string)
	ErrorPrint(text string)
	DebugPrint(header, text string)
	ReplaceVars(text string, vars interface{}, funcs template.FuncMap) (string, error)
	PrettyJson(in string) string
	PrettyMarshal(data interface{}) (string, error)
	UnmarshalFile(fileName string, fs filesystem.FileSystemInterface) (map[string]interface{}, error)
	UnmarshalByExt(file string, configData []byte) (map[string]interface{}, error)
	UnmarshalJson(configData []byte) (map[string]interface{}, error)
	UnmarshalYaml(configData []byte) (map[string]interface{}, error)
	GetFuncMap() template.FuncMap
}

type Helpers struct {
}

func NewHelpers() *Helpers {
	return &Helpers{}
}
