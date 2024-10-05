package filesystem

import "github.com/spf13/afero"

//go:generate mockgen -source=filesystem_interface.go  -destination=./mocks/filesystem.go -package=filesystem_mocks
type FileSystemInterface interface {
	GetAferoFs() afero.Fs
	CreateDirectoryIfNotExists(path string) error
	ReadFile(filePath string) (string, error)
	CreateFileIfNotExists(filename string) (afero.File, error)
	CreateFile(filename string) (afero.File, error)
	WriteFile(filename string, content string) error
	ReadDir(path string) ([]string, error)
}
