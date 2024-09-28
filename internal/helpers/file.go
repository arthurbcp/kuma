package helpers

import (
	"io"
	"os"

	"github.com/gookit/color"
)

func CreateDirectoryIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			color.Gray.Printf("  ❌ " + path + "\n")
			return err
		}
		color.Gray.Printf("  ✅ " + path + "\n")
	}
	return nil
}

func ReadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
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

func CreateFileIfNotExists(filename string) (*os.File, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return CreateFile(filename)
	}

	return nil, nil
}

func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return file, err
	}

	return file, nil
}

func WriteFile(filename string, content string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, []byte(content), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func ReadDir(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}
