package shared

var (
	// TemplateVariables holds the variables for template replacement during the generate process.
	TemplateVariables map[interface{}]interface{}

	// KumaConfigFilePath specifies the path to the Kuma configuration file.
	KumaConfigFilePath string = ".kuma-files/kuma-config.yaml"

	// KumaRunsPath defines the path to the directory containing Kuma runs.
	KumaRunsPath string = ".kuma-files/kuma-runs.yaml"

	// KumaTemplatesPath defines the path to the directory containing Kuma templates.
	KumaTemplatesPath string = ".kuma-files/kuma-templates"
)
