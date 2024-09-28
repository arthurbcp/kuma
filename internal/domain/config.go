package domain

// Config holds the configuration paths required for the application.
// It includes paths for the project directory and the templates directory.
type Config struct {
	// ProjectPath specifies the root directory where the project will be generated.
	ProjectPath string

	// TemplatesPath specifies the directory where template files are located.
	TemplatesPath string
}

// NewConfig creates and returns a new Config instance with the provided
// project and templates paths.
//
// Parameters:
//   - projectPath: The file system path to the project directory.
//   - templatesPath: The file system path to the templates directory.
//
// Returns:
//
//	A pointer to a Config struct initialized with the given paths.
func NewConfig(projectPath string, templatesPath string) *Config {
	return &Config{
		ProjectPath:   projectPath,
		TemplatesPath: templatesPath,
	}
}
