package function

import (
	"fmt"
	"os"
	"pylang/config"
	"time"

	"github.com/briandowns/spinner"
)

const (
	red    = "\033[31m"
	yellow = "\033[33m"

	bold  = "\033[1m"
	reset = "\033[0m"
)

func Print(Log config.Log) {
	var message string

	if Log.Err != nil && !Log.IsAlert {
		message = fmt.Sprintf(
			"%s[pylang]%s %s\n  %sError {%s%s%s}%s\n    %sDetail:%s %s",
			red, reset, Log.Message, bold, reset, Log.Origin,
			bold, reset, bold, reset, Log.Err,
		)

	} else if Log.Err != nil && Log.IsAlert {
		message = fmt.Sprintf(
			"%s[pylang]%s %s\n  %sWarning {%s%s%s}%s\n    %sDetail:%s %s",
			yellow, reset, Log.Message, bold, reset, Log.Origin,
			bold, reset, bold, reset, Log.Err,
		)

	} else {
		message = fmt.Sprintf("%s[pylang]%s %s ", bold, reset, Log.Message)
		time.Sleep(time.Duration(Log.TimeStart) * time.Millisecond)
	}

	fmt.Println(message)
}

func Animation(msg string, frameMs int) *spinner.Spinner {
	s := spinner.New(
		spinner.CharSets[14],
		time.Duration(frameMs)*time.Millisecond,
		spinner.WithWriter(os.Stderr),
	)
	s.Suffix = " " + msg
	return s
}
