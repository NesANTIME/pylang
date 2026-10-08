package main

import (
	"fmt"
	"os"
	"pylang/cmd"
	"pylang/config"
	"pylang/core"
	"pylang/internal/function"
)

func main() {
	config.Icon()

	if err := cmd.RootCmD.Execute(); err != nil {
		os.Exit(1)
	}

	path, err := function.ResolvePath(cmd.Path)
	if err != nil {
		function.Print(config.Log{
			Origin:  "function(resolvePath)",
			IsAlert: false,
			Err:     err,
			Message: "The directory could not be resolved",
		})
		return
	}

	cmd.Path = path

	function.Print(config.Log{
		Message:   fmt.Sprintf("Working directory: %s%s%s", config.Bold, cmd.Path, config.Reset),
		TimeStart: 100,
	})

	function.Print(config.Log{
		Message:   "[pylang] Starting compilation...",
		TimeStart: 100,
	})

	core.Start(cmd.Path, &cmd.Arg)
}
