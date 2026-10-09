package module

import (
	"fmt"
	"os"
	"pylang/config"
	"pylang/internal/compressed"
	"pylang/internal/logger"
)

func ToBytes(app *config.Application) error {
	s := logger.Animation("Starting file compression", 500)
	s.Start()

	// --
	pathWork, err := compressed.CompressedModel(
		app.Config.ModelCompressed, app.Name,
		app.Path.Work, app.Config.ExcluideCompiler,
	)

	if err != nil {
		s.Stop()
		logger.Print(config.Log{
			Message: "Compression could not be applied to the file group in the directory",
			Origin:  "ToBytes(compressedModel)", Err: err, IsAlert: false,
		})
		return err
	}
	// --

	// --
	pathInterpreter, err := compressed.CompressedModel(
		app.Config.ModelCompressed, fmt.Sprintf("interpreter-%s", app.Name),
		app.Path.Interpreter, app.Config.ExcluideInterpreter,
	)

	if err != nil {
		s.Stop()
		logger.Print(config.Log{
			Message: "Compression could not be applied to the file group in the directory",
			Origin:  "ToBytes(compressedModel)", Err: err, IsAlert: false,
		})
		return err
	}
	// --

	s.Stop()
	s.Suffix = "Successful compression"
	logger.Print(config.Log{
		Message: "Compiling to bytecode", TimeStart: 390,
	})

	// --
	app.ContentByte.Code, err = os.ReadFile(pathWork)
	if err != nil {
		logger.Print(config.Log{
			Message: "The working directory could not be compiled into bytecode",
			Origin:  "ToBytes(*os.ReadFile)", Err: err, IsAlert: false,
		})
		return err
	}

	app.ContentByte.Runtime, err = os.ReadFile(pathInterpreter)
	if err != nil {
		logger.Print(config.Log{
			Message: "The interpreter could not be compiled into bytecode",
			Origin:  "ToBytes(*os.ReadFile)", Err: err, IsAlert: false,
		})
		return err
	}
	// --

	return nil
}
