package helpers

import "html/template"

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
}

type Helpers struct {
}

func NewHelpers() *Helpers {
	return &Helpers{}
}
