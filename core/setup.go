package core

import (
	"os"
	"path/filepath"
	"pylang/config"
	"pylang/internal/function"
	"pylang/internal/module"
)

var Application config.Application

func Start(path string, args *config.Args) {
	Application.Path.Work = path
	Application.Path.PylangBuild = filepath.Join(Application.Path.Work, "pylang-cache", "build")
	Application.Path.PylangCompiler = filepath.Join(Application.Path.Work, "pylang-cache", "compile")

	// Creacion de carpetas
	function.Print(config.Log{
		Message:   "Creating working directory",
		TimeStart: 300,
	})

	for _, path := range []string{Application.Path.PylangBuild, Application.Path.PylangCompiler} {
		if err := os.MkdirAll(path, 0755); err != nil {
			return
		}
	}

	function.Print(config.Log{
		Message:   "Implementing arguments",
		TimeStart: 250,
	})

	Application.Parameters = *args

	// Codigo en bytes del proyecto
	function.Print(config.Log{
		Message:   "Cloning the working directory and Python interpreter to bytes",
		TimeStart: 400,
	})

	if err := module.BytesEmbed(&Application); err != nil {
		return
	}
}
