package main

import (
	"github.com/centurylinkcloud/clc-go-cli/auth"
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/formatter_provider"
	"github.com/centurylinkcloud/clc-go-cli/model_adjuster"
	"github.com/centurylinkcloud/clc-go-cli/model_loader"
	"github.com/centurylinkcloud/clc-go-cli/model_validator"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"github.com/centurylinkcloud/clc-go-cli/parser"
	"github.com/centurylinkcloud/clc-go-cli/state"

	"io"
)

func run(args []string) string {
	if len(args) == 0 {
		return ussage()
	}
	cmdArg := ""
	optionArgs := args[1:]
	if len(args) >= 2 {
		cmdArg = args[1]
		optionArgs = args[2:]
	}
	cmd, err := command_loader.LoadCommand(args[0], cmdArg)
	if err != nil {
		return err.Error()
	}
	if cmd.Command() == "" {
		optionArgs = args[1:]
	}
	parsedArgs, err := parser.ParseArguments(optionArgs)
	if err != nil {
		return err.Error()
	}
	options, err := options.ExtractFrom(parsedArgs)
	if err != nil {
		return err.Error()
	}
	if options.Help {
		return cmd.ShowHelp()
	}
	err = model_loader.LoadModel(parsedArgs, cmd.InputModel())
	if err != nil {
		return err.Error()
	}
	err = model_validator.ValidateModel(cmd.InputModel())
	if err != nil {
		return err.Error()
	}
	err = model_adjuster.ApplyDefaultBehaviour(cmd.InputModel())
	if err != nil {
		return err.Error()
	}
	conf, err := config.LoadConfig()
	if err != nil {
		return err.Error()
	}
	if cmd.Resource() == "login" {
		if options.User == "" || options.Password == "" {
			return "Both --user and --password options must be specified."
		}
		conf.User = options.User
		conf.Password = options.Password
		config.Save(conf)
		return ""
	}
	cn, err := auth.AuthenticateCommand(options, conf)
	if err != nil {
		return err.Error()
	}
	err = cmd.Execute(cn)
	if err != nil {
		if err == io.EOF {
			return ""
		}
		return err.Error()
	}
	err = state.SaveLastResult(cmd.OutputModel())
	if err != nil {
		return err.Error()
	}
	f, err := formatter_provider.GetOutputFormatter(options)
	if err != nil {
		return err.Error()
	}
	outputModel := cmd.OutputModel()
	if options.Query != "" {
		queried, err := parser.ParseQuery(outputModel, options.Query)
		if err != nil {
			return err.Error()
		} else if queried == nil {
			return "No results found for the given query."
		} else {
			outputModel = queried
		}
	}
	output, err := f.FormatOutput(outputModel)
	if err != nil {
		return err.Error()
	}
	return output
}

func ussage() string {
	res := "Ussage: clc <resource> <command> [options and parameters], for example 'clc server create --name my-server ...'\n"
	res += "To get help and list all avaliable resources or commands, you can use 'clc --help' or 'clc <resource> --help' or 'clc <resource> <command> --help'\n"
	return res
}
