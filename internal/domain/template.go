package domain

type Template struct {
	Name        string
	Tags        []string
	Description string
}

func NewTemplate(name, description string, tags []string) Template {
	return Template{
		Name:        name,
		Tags:        tags,
		Description: description,
	}
}
