package domain

type Config struct {
	ProjectPath   string
	TemplatesPath string
}

func NewConfig(projectPath string, templatesPath string) *Config {
	return &Config{
		ProjectPath:   projectPath,
		TemplatesPath: templatesPath,
	}
}
