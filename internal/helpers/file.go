package helpers

import (
	"io"
	"os" // Import the os package

	"github.com/gookit/color"
	"github.com/spf13/afero"
)

// Initialize the default filesystem. You can replace this with a mock filesystem in tests.
var AppFs = afero.NewOsFs()

// CreateDirectoryIfNotExists creates a directory if it does not already exist.
func CreateDirectoryIfNotExists(path string) error {
	exists, err := afero.DirExists(AppFs, path)
	if err != nil {
		return err
	}
	if !exists {
		// Use os.FileMode directly
		err := AppFs.MkdirAll(path, os.ModePerm)
		if err != nil {
			color.Gray.Printf("  ❌ " + path + "\n")
			return err
		}
		color.Gray.Printf("  ✅ " + path + "\n")
	}
	return nil
}

// ReadFile reads the content of a file and returns it as a string.
func ReadFile(filePath string) (string, error) {
	file, err := AppFs.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// CreateFileIfNotExists creates a file if it does not already exist.
func CreateFileIfNotExists(filename string) (afero.File, error) {
	exists, err := afero.Exists(AppFs, filename)
	if err != nil {
		return nil, err
	}
	if !exists {
		return CreateFile(filename)
	}
	return nil, nil
}

// CreateFile creates or truncates the named file.
func CreateFile(filename string) (afero.File, error) {
	file, err := AppFs.Create(filename)
	if err != nil {
		return file, err
	}
	return file, nil
}

// WriteFile writes a string to a file, overwriting it if it already exists.
func WriteFile(filename string, content string) error {
	// Open the file with write permissions, create if not exists, truncate if exists
	// Use os.FileMode directly
	file, err := AppFs.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// ReadDir reads the directory named by path and returns a slice of file names.
func ReadDir(path string) ([]string, error) {
	entries, err := afero.ReadDir(AppFs, path)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, entry := range entries {
		fileNames = append(fileNames, entry.Name())
	}

	return fileNames, nil
}
