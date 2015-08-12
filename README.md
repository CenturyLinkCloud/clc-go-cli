# CenturyLink CLI

Command Line Interface for manilulating the CenturyLink IaaS.

## Getting Started

Download a binary compiled for your platform from the [releases page](https://github.com/CenturyLinkCloud/clc-go-cli/releases).

## Getting Help

Explore the available resources, commands, options and other useful guidance using the `--help` option:
`clc --help`, `clc <resource> --help` and `clc <resouce> <command> --help` are all at your service.

The documentation of the underlying HTTP API can be found [here](https://www.ctl.io/api-docs/v2/).

## The Development Process

* [Install Go](https://golang.org/).
* Install Godep: `go get github.com/tools/godep`.
* Clone this repo (you do not have to use `go get`).
* [Ensure your $GOPATH is set correctly](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable).
* Install dependencies with Godep: enter the repo's root and `godep restore`.
* Use the dev script to run commands: `./dev <resource> <command>`.
* Install go vet: `go get code.google.com/p/go.tools/cmd/vet`.
* Before commit check that `gofmt -d=true ./..` and `go vet ./...` do not produce any output and check that all tests pass via `./run_tests`.

If you want to make an executabe, simply run `./build`. The binary will appear in the `./out` folder.
