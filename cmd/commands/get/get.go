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

	"github.com/arthurbcp/kuma-cli/cmd/commands/run"
	"github.com/arthurbcp/kuma-cli/internal/helpers"
	"github.com/google/go-github/github"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	Repo     string
	Template string
)

var Templates = map[string]string{
	"typescript-rest-openapi-services": "arthurbcp/typescript-rest-openapi-services",
}

// GetCmd represents the 'get' subcommand.
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a templates repository remotely",
	Run: func(cmd *cobra.Command, args []string) {
		download(cmd)
	},
}

func download(cmd *cobra.Command) {
	helpers := helpers.NewHelpers()
	client := github.NewClient(nil)

	if Template == "" && Repo == "" {
		cmd.Help()
		fmt.Println("\nplease specify a template or a repository")
		os.Exit(1)
	}

	repo, ok := Templates[Template]
	if !ok {
		repo = Repo
	}
	helpers.TitlePrint("getting templates from github repository...")
	splitRepo := strings.Split(repo, "/")
	if len(splitRepo) != 2 {
		helpers.TitlePrint("invalid repository name: " + Repo)
		os.Exit(1)
	}
	org := splitRepo[0]
	repoName := splitRepo[1]
	downloadRepo(client, org, repoName, "", ".kuma-files")
	helpers.CheckMarkPrint("templates downloaded successfully!")
	vars := map[string]interface{}{
		"data": map[string]interface{}{},
	}
	run.ExecRun("initial", vars)
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
