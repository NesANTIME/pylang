package function

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func ResolvePath(path string) (string, error) {
	var (
		realpath string
		err      error
	)

	if path == "." {
		realpath, err = os.Getwd()
	} else {
		realpath, err = filepath.Abs(path)
	}

	info, err := os.Stat(realpath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return realpath, fmt.Errorf("%s no existe", realpath)
		}

		return realpath, err
	}

	if !info.IsDir() {
		return realpath, fmt.Errorf("%s no es un directorio", realpath)
	}

	return realpath, err
}
