package main

import (
	"fmt"
	"os"
	"path/filepath"
	"pylang/cmd"
	"pylang/config"
	"pylang/core"
	"pylang/internal/function"
	"pylang/internal/logger"
)

func main() {
	var (
		Config config.Configs
		err    error
	)

	config.Icon()

	if err := cmd.RootCmD.Execute(); err != nil {
		os.Exit(1)
	}

	// ---

	cmd.Path, err = function.ResolvePath(cmd.Path)
	if err != nil {
		logger.Print(config.Log{
			Message: "The directory could not be resolved",
			Origin:  "function(resolvePath)", IsAlert: false, Err: err,
		})
		return
	}

	// ---

	logger.Print(config.Log{
		Message: fmt.Sprintf("Working directory: %s%s%s", config.Bold, cmd.Path, config.Reset), TimeStart: 100,
	})

	logger.Print(config.Log{
		Message: "Loading work environment configuration", TimeStart: 190,
	})

	// --
	Config, err = function.ReadConfig(filepath.Join(cmd.Path, "pylang-config.conf"))
	if err != nil {
		logger.Print(config.Log{
			Message: "The working environment configuration could not be loaded",
			Origin:  "function(ReadConfig)", IsAlert: false, Err: err,
		})
		return
	}

	Config.SO = cmd.SO
	Config.ARCH = cmd.ARCH
	// --

	logger.Print(config.Log{
		Message:   "[pylang] Starting compilation...",
		TimeStart: 100,
	})

	core.Start(cmd.Path, cmd.Name, &Config)
}
