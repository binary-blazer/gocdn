package lib

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func GetFilePath(fileId string) (string, error) {
	var dirPath string
	err := filepath.Walk("uploads", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.Contains(info.Name(), fileId) {
			dirPath = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if dirPath == "" {
		return "", os.ErrNotExist
	}
	imagePaths, err := filepath.Glob(filepath.Join(dirPath, "image.*"))
	if err != nil {
		return "", err
	}

	if len(imagePaths) == 0 {
		return "", os.ErrNotExist
	}
	imagePath := imagePaths[0]

	if !FileExists(imagePath) {
		return "", os.ErrNotExist
	}
	return imagePath, nil
}

func JoinStrings(list []string, sep string) string {
	return strings.Join(list, sep)
}

func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func CreateFolder(path string) error {
	if !FileExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFile(path string) error {
	if !FileExists(path) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func WriteYAML(path string, data interface{}) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
