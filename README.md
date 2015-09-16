# CenturyLink CLI

Command Line Interface for manipulating the CenturyLink IaaS.

## Getting Started

### Download the tool

Download a release tarball compiled for your platform from the [releases page](https://github.com/CenturyLinkCloud/clc-go-cli/releases). Extract an executable (`clc.exe` on Windows and `clc` on other platforms) and optionally put it on your PATH.

### Log in to your IaaS account

You need to be authenticated with a username and password in order to execute CLI commands. There are plenty of ways to set the credentials:

1. A config (see [the Config section](#set-up-the-configuration-file) for more details).
2. A `login` command:
  
  `clc login --user bob --password passw0rd`.
  
  This puts the passed credentials into the config.
3. `CLC_USER` and `CLC_PASSWORD` environment variables:

  `CLC_USER=bob CLC_PASSWORD=passw0rd clc server list`
  
  or on Windows in PowerShell:
  
  `$env:CLC_USER="bob"; $env:CLC_PASSWORD="passw0rd"; clc.exe server list`.
  
  If specified, they take precedence over the values from the configuration file, if any.
4. `--user` and `--password` command options:

  `clc server list --user bob --password passw0rd`.
  
  If specified, they take precedence over the values from the configuration file and environment variables, if any.

### Set up the configuration file

The program uses a configuration file located at `$HOME/.clc/config.yml` on Linux/Unix/Mac and `C:\Users\%username%\clc\config.yml` on Windows. One is created automatically on the first execution of any command. The file is in YAML format. The following fields count:

* `user` and `password`: the credentials used for authentication.
* `defaultformat`: a default output format, either `json`, `table` or `text`.
* `profiles`: a hash of alternative credentials.
* `defaultdatacenter`: a short code for a default data center. See [the corresponding section](#specify-a-default-data-center).

An example of a configuration file:

```
user: bob
password: passw0rd
defaultformat: "table"
defaultdatacenter: "CA1"
profiles:
  alice:
    username: alice
    passwod: pa33w0rd
```

Choose a profile either via a `--profile` option or `CLC_PROFILE` environment variable: `clc server list --profile alice`.

### Specify a default data center

A number of commands require a data center to be specified (via the `--data-center` option) what limits entities (groups, servers, policies, etc) operated upon to only those belonging to this data center.

There is a possibility to set a default one so that you do not need to specify it explicitly with every command.

You can either set a data center in the config using the `defaultdatacenter` field or execute a command:

```
clc data-center set-default --data-center <a-short-code-for-a-data-center>
```

You can query the current default with:

```
clc data-center show-default
```

Or unset it using:

```
clc data-center unset-default
```

### Enjoy the tool

Below are some examples of CLI commands so that you can faster get to use it efficiently.

Explore the list of data centers:

`clc data-center list`

Find server template ids that contain the word "UBUNTU" in some data center `<data-center>`:

```
clc data-center get-deployment-capabilities --data-center <data-center> --query templates.name --output text | grep UBUNTU
```

Search for the root group id of the data center under consideration:

```
clc group list --all --filter location-id=<data-center> --query id --output text
```

Or, the same thing can be accomplished by issuing:

```
clc group list --data-center <data-center> --query id --output text
```

Get the list of subgroups. Use a "SubGroup" alias for subgroups ids in the output:

```
clc group get --group-id <root-group-id> --query 'groups.{SubGroup:id}'
```

Create a server:

```
clc server create --name myserv --source-server-id <template-id> --group-id <group-id> --cpu 1 --memoryGB 1
```

Wait until the server has been created:

```
clc wait
```

Query only servers with status "active" and see the output as a table:

```
clc server list --all --filter status=active --output table
```

Increase the server's CPUs count and log the HTTP request/response data:

```
clc server update --server-id <server_id> --cpu 2 --trace
```

Show billing details of the servers of the group as a table:

```
clc group get-billing-details --group-id <group-id> --query groups.<group-id>.servers --output table
```

Make a skeleton of a command for getting groups with servers inside:

```
clc group list --filter 'servers-count>0' --generate-cli-skeleton > groups_with_servers.json
```

Apply the skeleton:

```
clc group list --from-file groups_with_servers.json
```

## Autocomplete

### Bash

Release tarballs for Linux/Unix/Darwin (starting from the release `2015-08-18`) contain 2 files for enabling autocomplete: `bash_autocomplete` and `install_autocompletion`. Execute `source bash_autocomplete` to turn on autocomplete for the current terminal session. `install_autocompletion` is provided for you to install autocomplete user-wide. The script, upon invoking,
copies the `bash_autocomplete` contents into `~/.bash_completion/clc` and updates `~/.bashrc` accordingly.

### PowerShell

Only v3 support is provided because previous versions do not support custom autocomplete handlers. PowerShell v3 is distributed as a part of Windows Management Framework 3.0, which can be downloaded [from here](http://www.microsoft.com/en-us/download/details.aspx?id=34595). You can check the version by typing `$PSVersionTable.PSVersion`.

To turn on autocomplete execute `.\powershell3_autocomplete.ps1`. You can find the file in the release tarball for Windows.

## Getting Help

Explore the available resources, commands, options and other useful guidance using the `--help` option:
`clc --help`, `clc <resource> --help` and `clc <resouce> <command> --help` are all at your service.

The documentation of the underlying HTTP API can be found [here](https://www.ctl.io/api-docs/v2/).

## The Development Process

The development is supposed to happen on Unix/Linux/Mac systems. Some of the
instructions below may not work properly on Windows.

* [Install Go](https://golang.org/).
* Install Godep: `go get github.com/tools/godep`.
* Clone this repo (do **not** use `go get`).
* [Ensure your $GOPATH is set correctly](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable).
* Install dependencies with Godep: enter the repo's root and `godep restore`.
* Use the dev script to run commands: `./dev <resource> <command>`.
* Install go vet: `go get code.google.com/p/go.tools/cmd/vet`.
* Before commit check that `gofmt -d=true ./..` and `go vet ./...` do not produce any output (except for that coming from `Godeps/_workspace` - ignore it) and check that all tests pass via `./run_tests`.

If you want to make an executable, simply run `./build`. The binary will appear in the `./out` folder.
