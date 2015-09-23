# CenturyLink CLI

This is the Command Line Interface (CLI) for manipulating the CenturyLink Infrastructure as a Service (IaaS).

## Getting Started

### Download the tool

Click a link below to download the latest OS system release.

[MacOS](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2015-09-21/clc-2015-09-21-darwin.tar.gz) | [Linux](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2015-09-21/clc-2015-09-21-linux-amd64.tar.gz) | [Windows](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2015-09-21/clc-2015-09-21-windows-x64.zip)

**Note:** You can see previous releases and release notes on the [releases page](https://github.com/CenturyLinkCloud/clc-go-cli/releases).

### Log in to your IaaS account

You need to be authenticated with a username and password in order to execute CLI commands. There are plenty of ways to set the credentials:

* A config (see [the Config section](#set-up-the-configuration-file) for more details).
* A `login` command:

  `clc login --user bob --password passw0rd`.

  This puts the passed credentials into the config.
* `CLC_USER` and `CLC_PASSWORD` environment variables:

  `CLC_USER=bob CLC_PASSWORD=passw0rd clc server list`

  or on Windows in PowerShell:

  `$env:CLC_USER="bob"; $env:CLC_PASSWORD="passw0rd"; clc.exe server list`.

  **Note:** If specified, these values take precedence over the values from the configuration file and environment variables (if any).
* `--user` and `--password` command options:

  `clc server list --user bob --password passw0rd`.

  **Note:** If specified, these values take precedence over the values from the configuration file and environment variables (if any).

### Set up the configuration file

The program uses a configuration file located at `$HOME/.clc/config.yml` on Linux/Unix/Mac and `C:\Users\%username%\clc\config.yml` on Windows. A config file is created automatically on the first execution of any command. The file is in YAML format. The following fields count:

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

Choose a profile either via a `--profile` option or a `CLC_PROFILE` environment variable: `clc server list --profile alice`.

### Specify a default data center

A number of commands require a default data center variable (via the `--data-center` option) in order to ensure that the command only operates on the entities (groups, servers, policies, etc) belonging to the specific data center.

There is an option to set a default data center so that you do not need to specify it explicitly with every command.

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

### Identify entities by ID or by name

There are many commands that depend on the identification of specific entities such as `server`, `group` or `network`.
A common pattern for specyfing the entity ID in the command line is `--<entity>-id`, like `--server-id` or `--load-balancer-id`.

Alternatively, the name of an entity can be specified instead of the ID (the common pattern is `--<entity>-name`). This approach has some important subtleties to mention:

* You can't specify both an ID and a name for the same entity.
* If there is more than one entity with the specified name, an error occurs.
* Autocomplete, if turned on, works for names and does **not** work for IDs. See [the Autocomplete section](#autocomplete-for-entity-names) for details.

### Enjoy the tool

Below is a list of CLI commands to help bring you up-to-speed on using the tool.

**Explore the list of data centers:**

`clc data-center list`

**Find server template IDs that contain the word "UBUNTU" in some data center `<data-center>`:**

```
clc data-center get-deployment-capabilities --data-center <data-center> --query templates.name --output text | grep UBUNTU
```

**Search for the root group ID of the data center under consideration:**

```
clc group list --all --filter location-id=<data-center> --query id --output text
```

Or, the same thing can be accomplished by issuing:

```
clc group list --data-center <data-center> --query id --output text
```

**Get the list of subgroups. Use a "SubGroup" alias for subgroups IDs in the output:**

```
clc group get --group-id <root-group-id> --query 'groups.{SubGroup:id}'
```

**Create your own group inside the one queried:**

```
clc group create --name "my group" --description "A group of mine" --parent-group-id <group-id> --custom-fields "id=<some-field>,value=<some-value>" "id=<another-field>,value=<another-value>"
```

**Note:** Pay attention to how we set custom fields. According to the command help, the `--custom-fields` argument accepts an array of objects with 2 keys each: `id` and `value`. The tool interprets multiple space-separated values as an array and each object can be specified using the `key1=value1,key2=value2,..`-notation, which is described in more detail further in the document.

**Create a server:**

```
clc server create --name myserv --source-server-id <template-id> --group-id <group-id> --cpu 1 --memory-gb 1
```

The same can be accomplished with a piece of JSON:

```
clc server create '{"name":"myserv","source-server-id":"<template-id>","group-id":"<group-id>","cpu":1,"memory-gb":1}'
```

Be careful with JSON. Keys and string values have to be enclosed in **double** quotes. Also, an expression may fail to be parsed unless it is enclosed in quotes, mainly because commas and spaces usually have special meanings in shells.

There is another notation for describing objects:

```
clc server create "name=myserv,source-server-id='<template-id>',group-id='<group-id>',cpu=1,memory-gb=1"
```

In this case you can use both `'` and `"` for both values and the whole expression, but be sure to escape special characters, as it has been partly described for JSON.

**Note:** this notation does not support nested objects and arrays.

You can also mix both options:

```
clc server create '{"name":"myserv"}' source-server-id='<template-id>' --group-id <group-id> --cpu 1 --memory-gb 1
```

**Note:** Be sure to put all the data **not bound to any command key** first, otherwise it will be interpreted as a value or an item of an array for the preceding command key.

**Wait until the server has been created:**

```
clc wait
```

**Query only servers with status "active" and see the output as a table:**

```
clc server list --all --filter status=active --output table
```

**Increase the server's CPUs count and log the HTTP request/response data:**

```
clc server update --server-id <server_id> --cpu 2 --trace
```

**Show billing details of the servers of the group as a table:**

```
clc group get-billing-details --group-id <group-id> --query groups.<group-id>.servers --output table
```

**Make a skeleton of a command for getting groups with servers inside:**

```
clc group list --filter 'servers-count>0' --generate-cli-skeleton > groups_with_servers.json
```

**Apply the skeleton:**

```
clc group list --from-file groups_with_servers.json
```

## Autocomplete

### Bash

Release tarballs for Linux/Unix/Darwin (starting from the release `2015-08-18`) contain 2 files for enabling autocomplete: `bash_autocomplete` and `install_autocompletion`. Execute `source bash_autocomplete` to turn on autocomplete for the current terminal session. `install_autocompletion` is provided for you to install autocomplete user-wide. The script copies the `bash_autocomplete` contents into `~/.bash_completion/clc` and updates `~/.bashrc` accordingly.

### PowerShell

Only v3 support is provided because previous versions do not support custom autocomplete handlers. PowerShell v3 is distributed as a part of Windows Management Framework 3.0, which can be downloaded [from here](http://www.microsoft.com/en-us/download/details.aspx?id=34595). You can check the version by typing `$PSVersionTable.PSVersion`.

To turn on autocomplete execute `.\powershell3_autocomplete.ps1`. You can find the file in the release tarball for Windows.

### Autocomplete for entity names

'Entity names' is a special item in the autocomplete list because options are fetched from the server. This kind of autocomplete only works under certain circumstances and the process may take a relatively long time.
Things you should note:

* The functionality does not work until the user has been authenticated. Authentication is needed to perform API requests.
* Since options are generated on the fly as the user enters a command, an entity name lookup for the data-center-dependent commands is only made within the default data center. If there is no default set, no options will appear.
* In bash, a waiting indicator in the form of dot rotation is shown for the time options that are fetched after the Tab has been pressed. Windows PowerShell does not support this kind of interaction - the input is simply blocked until the options have arrived.
* A cache is implemented to avoid making long subsequent requests to the server. The cache entry lifetime is 30 seconds.

## Getting Help

Explore the available resources, commands, options and other useful guidance using the `--help` option:
`clc --help`, `clc <resource> --help` and `clc <resouce> <command> --help` are all at your service.

The documentation of the underlying HTTP API can be found [here](https://www.ctl.io/api-docs/v2/).

## The Development Process

Development is set up for Unix/Linux/Mac systems. Some of the
instructions below may not work properly on Windows.

* [Install Go](https://golang.org/).
* Install Godep: `go get github.com/tools/godep`.
* Clone this repo (do **not** use `go get`).
* [Ensure your $GOPATH is set correctly](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable).
* Install dependencies with Godep: enter the repo's root and `godep restore`.
* Use the dev script to run commands: `./dev <resource> <command>`.
* Install go vet: `go get code.google.com/p/go.tools/cmd/vet`.
* Before commit check that `gofmt -d=true ./..` and `go vet ./...` do not produce any output (except for that coming from `Godeps/_workspace` - ignore it) and check that all tests pass via `./run_tests`.

If you want to make an executable, simply run `./scripts/build`. The binary will appear in the `./out` folder.
