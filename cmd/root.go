package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Name string
	Path string

	SO   string
	ARCH string

	RootCmD = &cobra.Command{
		Use:   "pylang",
		Short: "Compiler for embedding Python environments within Go executables (Golang)",

		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("A subcommand is required: {compile}")
		},
	}
)

var compileCmd = &cobra.Command{
	Use:   "compile <path>",
	Short: "Compile the project to the indicated path",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("compile requires 1 argument (the path), received %d", len(args))
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		Path = args[0]
		return nil
	},
}

func init() {
	compileCmd.Flags().StringVarP(&SO, "cc-os", "s", string(runtime.GOOS), "Specify processor architecture for cross-compilation")
	compileCmd.Flags().StringVarP(&ARCH, "cc-arch", "a", string(runtime.GOARCH), "specify operating system for cross-compilation")

	compileCmd.Flags().StringVarP(&Name, "output", "o", "setup", "output file/directory")

	RootCmD.AddCommand(compileCmd)
}
