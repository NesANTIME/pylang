package compressed

import (
	"fmt"
	"path/filepath"
	"pylang/config"
	"pylang/internal/logger"
)

func CompressedModel(model string, nameFile string, pathWork string, excluides []string) (string, error) {
	path := filepath.Join(pathWork, nameFile)

	var (
		pathFinal string
		err       error = nil
	)

	switch model {
	case "zip":
		pathFinal, err = compressedZip(pathWork, path, excluides)

	default:
		pathFinal, err = compressedZip(pathWork, path, excluides)

		logger.Print(config.Log{
			Message: fmt.Sprintf(
				"The compression type is invalid or incompatible. The compression model used was %s", model),
			Origin: "CompressedModel(*ModelCompressed)", IsAlert: true,
			Err: fmt.Errorf("The compression type is invalid or incompatible"),
		})
	}

	return pathFinal, err
}
