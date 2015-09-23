# CenturyLink CLI

Command Line Interface for CenturyLink Cloud.

## Getting Started

### Download the tool

The latest release:

[MacOS tar.gz](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2015-09-21/clc-2015-09-21-darwin.tar.gz) | [MacOS pkg](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2015-09-21/clc-2015-09-21.pkg) | [Linux tar.gz](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2015-09-21/clc-2015-09-21-linux-amd64.tar.gz) | [Windows zip](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2015-09-21/clc-2015-09-21-windows-x64.zip)

See previous releases and release notes on the [releases page](https://github.com/CenturyLinkCloud/clc-go-cli/releases).

### Install it

#### On Linux:

Extract the archive. Its contents look like

```
clc-$VERSION-linux-amd64
|
 -- clc
 -- install_autocompletion
 -- autocomplete
   |
    -- bash_autocomplete
```

You can immediately start using the `clc` binary, or put it somewhere on your `PATH` for convenience. In order to turn on bash autocomplete you have to source the `bash_autocomplete` script. `install_autocompletion` copies this script to `~/.bash_completion/clc` and puts a line sourcing it to the `~/.bashrc` file so that autocomplete is turned on automatically in every terminal session.

#### On MacOS:

There are 2 options of installing the tool: a tar archive and a pkg file.

The tar archive is pretty much the same as the one for Linux. The only difference is that `install_autocompletion` alters `~/.bash_profile`, not `~/.bashrc`.

The pkg file is an easy way to set up things. It installs everything for you. The binary is placed at `/usr/local/bin`. The `~/.bash_completion/clc` script is created and a line sourcing this script is added to `~/.bash_profile` to enable bash autocomplete.

#### On Windows:

Extract the archive. Its contents look like

```
clc-$VERSION-windows-x64
|
 -- clc.exe
 -- autocomplete
   |
    -- powershell3_autocomplete.ps1
```

You can start using the binary right away, or put it somewhere on your PATH for convenience. To turn on PowerShell autocomplete execute the `powershell3_autocomplete.ps1` script.

Note that autocomplete only works with PowerShell version >= 3. PowerShell v3 is distributed as a part of Windows Management Framework 3.0, which can be downloaded [from here](http://www.microsoft.com/en-us/download/details.aspx?id=34595). You can check the version by typing `$PSVersionTable.PSVersion`.

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

### Identify entities by ID or by name

There is a great deal of commands that depend on the instances of certain entities, such as `server`, `group`, `network`.
A common pattern for specyfing such entities in the command line is `--<entity>-id`, like `--server-id` or `--load-balancer-id`.

Alternatively, a name of an entity can be specified (the common pattern is `--<entity>-name`). This approach has some important subtleties to mention:

* You can't specify both an ID and a name for the same entity.
* If there is more than one entity with the specified name, an error occurs.
* Autocomplete, if turned on, works for names and does **not** work for IDs. See [the Autocomplete section](#autocomplete-for-entity-names) for details.

### Enjoy the tool

Below are some examples of CLI commands to help you faster get to use the tool.

Explore the list of data centers:

`clc data-center list`

Find server template IDs that contain the word "UBUNTU" in some data center `<data-center>`:

```
clc data-center get-deployment-capabilities --data-center <data-center> --query templates.name --output text | grep UBUNTU
```

Search for the root group ID of the data center under consideration:

```
clc group list --all --filter location-id=<data-center> --query id --output text
```

Or, the same thing can be accomplished by issuing:

```
clc group list --data-center <data-center> --query id --output text
```

Get the list of subgroups. Use a "SubGroup" alias for subgroups IDs in the output:

```
clc group get --group-id <root-group-id> --query 'groups.{SubGroup:id}'
```

Create your own group inside the one queried:

```
clc group create --name "my group" --description "A group of mine" --parent-group-id <group-id> --custom-fields "id=<some-field>,value=<some-value>" "id=<another-field>,value=<another-value>"
```

Note how we set custom fields here. According to the command help, the `--custom-fields` argument accepts an array of objects with 2 keys each: `id` and `value`. The tool interprets multiple space-separated values as an array and each object can be specified using the `key1=value1,key2=value2,..`-notation, which is described in more detail further in the document.

Create a server:

```
clc server create --name myserv --source-server-id <template-id> --group-id <group-id> --cpu 1 --memory-gb 1
```

The same can be accomplished with a piece of JSON:

```
clc server create '{"name":"myserv","source-server-id":"<template-id>","group-id":"<group-id>","cpu":1,"memory-gb":1}'
```

Be careful with JSON though. Keys and string values have to be enclosed in **double** quotes. Also, an expression may fail to be parsed unless it is enclosed in quotes, mainly because commas and spaces usually have special meanings in shells.

Moreover, there is yet another notation for describing objects:

```
clc server create "name=myserv,source-server-id='<template-id>',group-id='<group-id>',cpu=1,memory-gb=1"
```

In this case you can use both `'` and `"` for both values and the whole expression but be sure to escape special characters as it has been partly described for JSON. Note that this notation does not support nested objects and arrays.

Finally, you can mix the ways described:

```
clc server create '{"name":"myserv"}' source-server-id='<template-id>' --group-id <group-id> --cpu 1 --memory-gb 1
```

Be sure to put all the data **not bound to any command key** first, otherwise it will be interpreted as a value or an item of an array for the preceding command key.

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

### Autocomplete for entity names

Entity names is a special item in the autocomplete list because options are fetched from the server. Therefore such kind of autocomplete only works under certain circumstances and the process may take a relatively long time.
Below are the things you should be aware of:

* The functionality does not work until the user has been authenticated. Authentication is needed to perform API requests.
* Since options are generated on the fly as the user enters a command, an entity name lookup for the data-center-dependent commands is only made within the default data center. If there is no default set, no options are to appear.
* In bash, a waiting indicator in the form of dot rotation is shown for the time options are being fetched after the Tab has been pressed. Windows PowerShell does not support such kind of interaction: in it, the input is simply blocked until the options have arrived.
* A cache is implemented to avoid making long subsequent requests to the server. The cache entry lifetime is 30 seconds.

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

If you want to make an executable, simply run `./scripts/build`. The binary will appear in the `./out` folder.
