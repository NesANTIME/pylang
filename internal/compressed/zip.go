package compressed

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func shouldExclude(path string, info os.FileInfo, excludes []string) bool {
	name := info.Name()

	for _, exclude := range excludes {
		if name == exclude {
			return true
		}

		if strings.HasPrefix(exclude, "*") &&
			strings.HasSuffix(name, strings.TrimPrefix(exclude, "*")) {
			return true
		}
	}
	return false
}

// ---

func compressedZip(pathdir string, pathfile string, excludes []string) (string, error) {
	zipfile, err := os.Create(pathfile)
	if err != nil {
		return pathfile, err
	}
	defer zipfile.Close()

	zipArchive := zip.NewWriter(zipfile)
	defer zipArchive.Close()

	err = filepath.Walk(pathfile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if shouldExclude(path, info, excludes) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(pathfile, path)
		if err != nil {
			return err
		}

		writer, err := zipArchive.Create(relativePath)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, file)

		file.Close()

		return err
	})

	if err != nil {
		return pathfile, err
	}

	return pathfile, nil
}
