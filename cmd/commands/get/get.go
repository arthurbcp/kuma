// get.go
//
// Package get defines the 'get' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package get

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	execCmd "github.com/arthurbcp/kuma/cmd/commands/exec"
	"github.com/arthurbcp/kuma/cmd/shared"
	"github.com/arthurbcp/kuma/cmd/ui/selectInput"
	"github.com/arthurbcp/kuma/cmd/ui/utils/program"
	"github.com/arthurbcp/kuma/cmd/ui/utils/steps"
	"github.com/arthurbcp/kuma/internal/domain"
	"github.com/arthurbcp/kuma/internal/helpers"
	"github.com/arthurbcp/kuma/internal/services"
	"github.com/arthurbcp/kuma/pkg/filesystem"
	"github.com/arthurbcp/kuma/pkg/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gookit/color"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	Repo     string
	Template string
)

// GetCmd represents the 'get' subcommand.
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a templates repository remotely",
	Run: func(cmd *cobra.Command, args []string) {
		if Template == "" && Repo == "" {
			Template = handleTea()
		}
		download(cmd)
	},
}

func handleTea() string {
	program := program.NewProgram()
	var options = make([]steps.Item, 0)

	for repository, template := range shared.Templates {
		options = append(options, steps.NewItem(
			template.Name,
			repository,
			template.Description,
			template.Tags,
		))
	}

	output := &selectInput.Selection{}
	p := tea.NewProgram(selectInput.InitialSelectInputModel(options, output, "Select a template or type \"o\" to use a different repository", true, program))
	_, err := p.Run()

	program.ExitCLI(p)

	if err != nil {
		style.ErrorPrint("error running program: " + err.Error())
		os.Exit(1)
	}
	return output.Choice
}

func addKumaModule(newModule string) error {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	modulesFile := shared.KumaFilesPath + "/kuma-modules.yaml"
	_, err := fs.CreateFileIfNotExists(modulesFile)
	if err != nil {
		return err
	}
	modules, err := helpers.UnmarshalFile(modulesFile, fs)
	if err != nil {
		return err
	}

	module, err := getModule(newModule)
	if err != nil {
		return err
	}

	mapModule, err := helpers.StructToMap(module)
	if err != nil {
		return err
	}
	modules[newModule] = mapModule

	yamlContent, err := yaml.Marshal(modules)
	if err != nil {
		return err
	}
	fs.WriteFile(modulesFile, string(yamlContent))
	return nil
}

func getModule(module string) (domain.Module, error) {
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	configData, err := helpers.UnmarshalFile(shared.KumaFilesPath+"/"+module+"/kuma-config.yaml", fs)
	if err != nil {
		return domain.Module{}, err
	}
	runsService := services.NewRunService(shared.KumaRunsPath+"/"+module, fs)
	runs, err := runsService.GetAll()
	if err != nil {
		return domain.Module{}, err
	}
	return domain.NewModule(configData, runs), nil
}

func download(cmd *cobra.Command) {
	if Template == "" && Repo == "" {
		cmd.Help()
		style.LogPrint("\nplease specify a template or a repository")
		os.Exit(1)
	}

	repo := Template
	if _, ok := shared.Templates[Template]; !ok {
		repo = Repo
	}
	style.LogPrint("getting templates from github repository as a submodule...")
	fs := filesystem.NewFileSystem(afero.NewOsFs())
	err := fs.CreateDirectoryIfNotExists(shared.KumaFilesPath)
	if err != nil {
		style.ErrorPrint("error creating kuma files directory: " + err.Error())
		os.Exit(1)
	}
	err = addGitSubmodule(repo)
	if err != nil {
		style.ErrorPrint("error adding submodule: " + err.Error())
		os.Exit(1)
	}
	style.CheckMarkPrint("templates downloaded successfully!\n")
	err = addKumaModule(GetModuleName(repo))
	if err != nil {
		style.ErrorPrint("error adding kuma module: " + err.Error())
		os.Exit(1)
	}
	execCmd.Execute()
	os.Exit(0)
}

func GetModuleName(repo string) string {
	splitRepo := strings.Split(repo, "/")
	return splitRepo[1]
}

func addGitSubmodule(repo string) error {
	cmd := exec.Command("git", "submodule", "add", shared.GitHubURL+"/"+repo, shared.KumaFilesPath+"/"+GetModuleName(repo))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// init sets up flags for the 'get' subcommand and binds them to variables.
func init() {
	// Repository name
	GetCmd.Flags().StringVarP(&Repo, "repo", "r", "", "Github repository")
	templates := make([]string, 0, len(shared.Templates))
	for key := range shared.Templates {
		templates = append(templates, key)
	}
	GetCmd.Flags().StringVarP(&Template, "template", "t", "", fmt.Sprintf("KUMA official template repositories:\n - %s",
		color.Gray.Sprintf(strings.Join(templates, "\n - "))))
}
