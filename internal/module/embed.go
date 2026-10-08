package module

import (
	"os"
	"path/filepath"
	"pylang/config"
	"pylang/internal/compressed"
)

var (
	pathCode    string
	pathRuntime string

	err error
)

func BytesEmbed(Application *config.Application) error {
	excluides := []string{".git", "safepy-cache", "__pycache__", ".env"}
	excluides = append(excluides, Application.Parameters.ExcluideFoldersAndFiles...)

	pathfile := filepath.Join(Application.Path.PylangCompiler, Application.Parameters.NameFile)

	switch Application.Parameters.ModelCompressed {

	case "zip":
		pathCode, err = compressed.Zip(Application.Path.Work, pathfile, excluides)
		if err != nil {
			return err
		}

	default:
		pathCode, err = compressed.Zip(Application.Path.Work, pathfile, excluides)
		if err != nil {
			return err
		}
	}

	Application.ContentByte.Code, err = os.ReadFile(pathCode)
	if err != nil {
		return err
	}

	// añadir contentbyte del interprete

	return nil
}
