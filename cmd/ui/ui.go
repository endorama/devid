/*
ui package proxies cli.Ui struct to provide a shared UI usable by all commands
(core and from plugins).

The package public interface is similar to cli.Ui (albeit not equal).
*/
package ui

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

var ui = &cli.ColoredUi{ //nolint:gochecknoglobals // UI is shared
	OutputColor: cli.UiColorNone,
	InfoColor:   cli.UiColorNone,
	ErrorColor:  cli.UiColorRed,
	WarnColor:   cli.UiColorYellow,

	Ui: &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	},
}

func Ask(string) {
	panic("Not yet implemented")
}

func AskSecret(string) {
	panic("Not yet implemented")
}

func Error(e error) {
	ui.Error(e.Error())
}

func Fatal(e error, ec int) {
	ui.Error(e.Error())
	os.Exit(ec)
}

func Info(format string, a ...interface{}) {
	ui.Info(fmt.Sprintf(format, a...))
}

func Output(format string, a ...interface{}) {
	ui.Output(fmt.Sprintf(format, a...))
}

func Warn(format string, a ...interface{}) {
	ui.Warn(fmt.Sprintf(format, a...))
}
