package main

import (
	"os"

	"github.com/entraktest/cmd"
	"github.com/entraktest/configs"
	"github.com/spf13/pflag"
)

func main() {

	root := cmd.RootCmd()
	var filename string
	fs := pflag.NewFlagSet("Root", pflag.ContinueOnError)
	fs.StringVarP(&filename,
		"file",
		"f",
		"",
		"Custom configuration filename",
	)
	root.Flags().AddFlagSet(fs)
	configuration := configs.New(filename, cmd.ConfigPath...)
	root.AddCommand(
		cmd.NewHttpCmd(
			configuration,
		).BaseCmd,
	)
	if err := root.Execute(); err != nil {
		panic(err.Error())
		os.Exit(1)
	}
}
