package main

import (
	cli "github.com/centurylinkcloud/clc-go-cli"
)

func run(args []string) string {
	if len(args) == 0 {
		return ussage()
	}
	cmdArg := ""
	if len(args) >= 2 {
		cmdArg = args[1]
	}
	cmd, err := cli.LoadCommand(args[0], cmdArg)
	if err != nil {
		return err.Error()
	}
	parsedArgs, err := cli.ParseArguments(args[2:])
	if err != nil {
		return err.Error()
	}
	options, err := cli.LoadOptions(parsedArgs)
	if err != nil {
		return err.Error()
	}
	err = cli.LoadModel(parsedArgs, cmd.InputModel())
	if err != nil {
		return err.Error()
	}
	err = cli.ValidateModel(cmd.InputModel())
	if err != nil {
		return err.Error()
	}
	err = cli.ApplyDefaultBehaviour(cmd.InputModel())
	if err != nil {
		return err.Error()
	}
	cn, err := cli.AuthenticateCommand(options)
	err = cmd.Execute(cn)
	if err != nil {
		return err.Error()
	}
	formatter := cli.GetOutputFormatter(options)
	output, err := formatter.FormatOutput(cmd.OutputModel())
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
