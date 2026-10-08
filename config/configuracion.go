package config

import "fmt"

const (
	Bold  = "\033[1m"
	Reset = "\033[0m"
)

func Icon() {
	var Icon []string = []string{
		"┌─┐┬ ┬┬  ┌─┐┌┐┌┌─┐", "├─┘└┬┘│  ├─┤││││ ┬", "┴   ┴ ┴─┘┴ ┴┘└┘└─┘ by NesANTIME",
	}

	for _, i := range Icon {
		fmt.Println(i)
	}
}

// ---
