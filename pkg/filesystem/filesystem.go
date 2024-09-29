package filesystem

import (
	"io"
	"os" // Import the os package

	"github.com/gookit/color"
	"github.com/spf13/afero"
)

type FileSystem struct {
	Fs afero.Fs
}

func NewFileSystem(fs afero.Fs) *FileSystem {
	return &FileSystem{
		Fs: fs,
	}
}

// CreateDirectoryIfNotExists creates a directory if it does not already exist.
func (s *FileSystem) CreateDirectoryIfNotExists(path string) error {
	exists, err := afero.DirExists(s.Fs, path)
	if err != nil {
		return err
	}
	if !exists {
		// Use os.FileMode directly
		err := s.Fs.MkdirAll(path, os.ModePerm)
		if err != nil {
			color.Gray.Printf("  ❌ " + path + "\n")
			return err
		}
		color.Gray.Printf("  ✅ " + path + "\n")
	}
	return nil
}

// ReadFile reads the content of a file and returns it as a string.
func (s *FileSystem) ReadFile(filePath string) (string, error) {
	file, err := s.Fs.Open(filePath)
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
func (s *FileSystem) CreateFileIfNotExists(filename string) (afero.File, error) {
	exists, err := afero.Exists(s.Fs, filename)
	if err != nil {
		return nil, err
	}
	if !exists {
		return s.CreateFile(filename)
	}
	return nil, nil
}

// CreateFile creates or truncates the named file.
func (s *FileSystem) CreateFile(filename string) (afero.File, error) {
	file, err := s.Fs.Create(filename)
	if err != nil {
		return file, err
	}
	return file, nil
}

// WriteFile writes a string to a file, overwriting it if it already exists.
func (s *FileSystem) WriteFile(filename string, content string) error {
	err := afero.WriteFile(s.Fs, filename, []byte(content), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// ReadDir reads the directory named by path and returns a slice of file names.
func (s *FileSystem) ReadDir(path string) ([]string, error) {
	entries, err := afero.ReadDir(s.Fs, path)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, entry := range entries {
		fileNames = append(fileNames, entry.Name())
	}

	return fileNames, nil
}
