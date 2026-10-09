package utils

import (
	"os"
	"pylang/config"
	"pylang/internal/logger"
)

func CreateFolders(paths []string) error {
	logger.Print(config.Log{
		Message: "Creating cache working directory", TimeStart: 300,
	})

	for _, path := range paths {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	return nil
}
