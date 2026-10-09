package core

import (
	"path/filepath"
	"pylang/config"
	"pylang/core/utils"
	"pylang/internal/logger"
	"pylang/internal/module"
)

var (
	app config.Application
)

func Start(path string, name string, cfg *config.Configs) {
	logger.Print(config.Log{
		Message: "Implementing arguments", TimeStart: 250,
	})

	app.Name = name
	app.Path.Work = path

	app.Path.PylangCache = filepath.Join(app.Path.Work, "pylang-cache", "cache")
	app.Path.PylangBuild = filepath.Join(app.Path.Work, "pylang-cache", "build")
	app.Path.PylangCompiler = filepath.Join(app.Path.Work, "pylang-cache", "compile")

	app.Config = *cfg

	// ---
	if err := utils.CreateFolders([]string{app.Path.PylangBuild,
		app.Path.PylangCompiler, app.Path.PylangCache}); err != nil {

		logger.Print(config.Log{
			Message: "Error creating cache folder",
			IsAlert: false,
			Err:     err,
			Origin:  "utils(CreateFolders)",
		})

		return
	}
	// ---

	logger.Print(config.Log{
		Message: "Connecting to embedded Python version repository", TimeStart: 200,
	})

	// --

	// --

	logger.Print(config.Log{
		Message: "Cloning the working directory and Python interpreter", TimeStart: 400,
	})

	// -
	if err := module.ToBytes(&app); err != nil {
		return
	}
	// -

	logger.Print(config.Log{
		Message: "Compilation to bytecode, successful", TimeStart: 300,
	})

}
