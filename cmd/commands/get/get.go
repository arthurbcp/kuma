// get.go
//
// Package get defines the 'get' subcommand for the Kuma CLI.
// It handles generating project scaffolds based on Go templates.
package get

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	execHandlers "github.com/arthurbcp/kuma-cli/cmd/commands/exec/handlers"
	"github.com/arthurbcp/kuma-cli/cmd/ui/selectInput"
	"github.com/arthurbcp/kuma-cli/cmd/ui/utils/program"
	"github.com/arthurbcp/kuma-cli/cmd/ui/utils/steps"
	"github.com/arthurbcp/kuma-cli/internal/domain"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/github"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	Repo     string
	Template string
)

var Templates = map[string]domain.Template{
	"arthurbcp/typescript-rest-openapi-services": domain.NewTemplate(
		"typescript-rest-openapi-services",
		"Create a library TypeScript with services typed for all endpoints described in a file Open API 2.0",
		[]string{"typescript", "openapi", "rest", "library"},
	),
}

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
	for repository, template := range Templates {
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

func download(cmd *cobra.Command) {
	client := github.NewClient(nil)

	if Template == "" && Repo == "" {
		cmd.Help()
		style.LogPrint("\nplease specify a template or a repository")
		os.Exit(1)
	}

	repo := Template
	if _, ok := Templates[Template]; !ok {
		repo = Repo
	}
	style.LogPrint("getting templates from github repository...")
	splitRepo := strings.Split(repo, "/")
	if len(splitRepo) != 2 {
		style.ErrorPrint("invalid repository name: " + Repo)
		os.Exit(1)
	}
	org := splitRepo[0]
	repoName := splitRepo[1]
	downloadRepo(client, org, repoName, "", ".kuma-files")
	style.CheckMarkPrint("templates downloaded successfully!\n")
	vars := map[string]interface{}{
		"data": map[string]interface{}{},
	}
	execHandlers.HandleRun("initial", vars)
	os.Exit(0)
}

// downloadFile writes the content of the file to the local file system
func downloadFile(content *github.RepositoryContent, baseDir string) error {
	// Ensure the content is a file
	if *content.Type == "file" {
		// Create file path
		filePath := filepath.Join(baseDir, *content.Path)
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return err
		}

		// Fetch the file's contents via the download URL
		if content.DownloadURL != nil {
			resp, err := http.Get(*content.DownloadURL)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			// Read the content
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			// Write the content to the file
			return os.WriteFile(filePath, data, 0644)
		}
	}
	return nil
}

// downloadRepo recursively downloads repository contents (directories and files)
func downloadRepo(client *github.Client, owner, repo, path, baseDir string) error {
	ctx := context.Background()

	// Fetch the contents of the path (directory or file)
	_, dirContents, _, err := client.Repositories.GetContents(ctx, owner, repo, path, nil)
	if err != nil {
		return err
	}

	// Iterate over each content item
	for _, content := range dirContents {
		if *content.Type == "dir" {
			// If it's a directory, recursively download its contents
			err = downloadRepo(client, owner, repo, *content.Path, baseDir)
			if err != nil {
				return err
			}
		} else if *content.Type == "file" {
			// If it's a file, download it
			err = downloadFile(content, baseDir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// init sets up flags for the 'get' subcommand and binds them to variables.
func init() {
	// Repository name
	GetCmd.Flags().StringVarP(&Repo, "repo", "r", "", "Github repository")
	templates := make([]string, 0, len(Templates))
	for key := range Templates {
		templates = append(templates, key)
	}
	GetCmd.Flags().StringVarP(&Template, "template", "t", "", fmt.Sprintf("KUMA official template repositories:\n - %s",
		color.Gray.Sprintf(strings.Join(templates, "\n - "))))
}
