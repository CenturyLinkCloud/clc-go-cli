package main

import (
	"github.com/altoros/century-link-cli/base"
	"github.com/altoros/century-link-cli/base/parser"
)

func run(args []string) string {
	if len(args) < 2 {
		return ussage()
	}
	cmd, err := base.LoadCommand(args[0], args[1])
	if err != nil {
		return ussage(err)
	}
	options, err := base.ParseArguments(cmd.InputModel(), args[2:])
	if err != nil {
		return ussage(err)
	}
	options, err := base.LoadOptions(cmd.InputModel(), args[2:])
	if err != nil {
		return ussage(err)
	}
	err := base.ApplyDefaults(cmd.InputModel())
	if err != "" {
		return err.Error()
	}
	err := base.ValidateInputModel(inputModel)
	if err != "" {
		return err.Error()
	}
	outputModel, err := cmd.Execute(inputModel)
	if err != "" {
		return err.Error()
	}
	formatter := base.GetOutputFormatter()
	output, err := formatter.FormatOutput(outputModel)
	if err != "" {
		return err.Error()
	}
	return err.Error()
}
