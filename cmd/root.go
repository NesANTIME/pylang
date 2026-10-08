package cmd

import (
	"fmt"
	"pylang/config"

	"github.com/spf13/cobra"
)

var (
	Path string
	Arg  config.Args

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
	compileCmd.Flags().StringVarP(&Arg.NameFile, "output", "o", "setup", "output file/directory")
	compileCmd.Flags().StringVarP(&Arg.ModelCompressed, "model", "m", "zip", "Compression model used: (zip)")

	RootCmD.AddCommand(compileCmd)
}
