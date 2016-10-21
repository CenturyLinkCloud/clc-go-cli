# CenturyLink Cloud CLI

This is the Command Line Interface (CLI) for the CenturyLink Cloud.

## Getting Started

### Download the tool

Click a link below to download the latest release for each OS.

[MacOS tar.gz](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2016-07-21/clc-2016-07-21-darwin-amd64.tar.gz) | [MacOS pkg](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2016-07-21/clc-2016-07-21.pkg) | [Linux tar.gz](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2016-07-21/clc-2016-07-21-linux-amd64.tar.gz) | [Windows zip](https://github.com/CenturyLinkCloud/clc-go-cli/releases/download/2016-07-21/clc-2016-07-21-windows-amd64.zip)

**Note:** You can see previous releases and release notes on the [releases page](https://github.com/CenturyLinkCloud/clc-go-cli/releases).

### Install it

#### On Linux:

Extract the archive. Its contents look like:

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

Extract the archive. Its contents look like:

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

### Log in to your CenturyLink Cloud account

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
* `profiles`: a hash of alternative credentials. See [Profiles](#profiles).
* `defaultdatacenter`: a short code for a default data center. See [the corresponding section](#specify-a-default-data-center).

An example of a configuration file:

```
user: bob
password: passw0rd
defaultformat: "table"
defaultdatacenter: "CA1"
profiles:
  alice:
    user: alice
    password: pa33w0rd
```

### Profiles

Each profile is a pair of alternative credentials to use. Profiles are specified in the configuration file.

To choose a profile for a single command invokation use either a `--profile` option or `CLC_PROFILE` environment variable: `clc server list --profile alice`.

Also, you can set up your default credentials from a profile via the `login` command: `clc login --profile alice`. Be careful, though, because your previous defaults will be overriden this way. Therefore it is a grood idea to have a profile for every of your users.

### Specify a default data center

A number of commands require a default data center variable (via the `--data-center` option) in order to ensure that the command only operates on the entities (groups, servers, policies, etc) belonging to the specific data center.

There is an option to set a default data center so that you do not need to specify it with every command.

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

### Switch accounts

Your account is determined automatically for you upon authentication. But you might also have sub accounts that have to be dealt with. Thus, the tool allows you to specify a custom account via an `--account-alias` option: `clc server list --account-alias MYSUBACC`.

### Identify entities by ID or by name

There are many commands that depend on the identification of specific entities such as `server`, `group` or `network`.
A common pattern for specifying the entity ID in the command line is `--<entity>-id`, like `--server-id` or `--load-balancer-id`.

Alternatively, the name of an entity can be specified instead of the ID (the common pattern is `--<entity>-name`). This approach has some important subtleties to mention:

* You can't specify both an ID and a name for the same entity.
* If there is more than one entity with the specified name, an error occurs.
* Autocomplete, if turned on, works for names and does **not** work for IDs. See [the Autocomplete section](#autocomplete-for-entity-names) for details.

### List of CLI commands

Below is a list of CLI commands to help bring you up-to-speed on using the tool.

####Explore the list of data centers:

`clc data-center list`

####Find server template IDs that contain the word "UBUNTU" in some data center `<data-center>`:

```
clc data-center get-deployment-capabilities --data-center <data-center> --query templates.name --output text | grep UBUNTU
```

####Search for the root group ID of the data center under consideration:

```
clc group list --all --filter location-id=<data-center> --query id --output text
```

Or, the same thing can be accomplished by issuing:

```
clc group list --data-center <data-center> --query id --output text
```

####Get the list of subgroups. Use a "SubGroup" alias for subgroups IDs in the output:

```
clc group get --group-id <root-group-id> --query 'groups.{SubGroup:id}'
```

####Create your own group inside the one queried:

```
clc group create --name "my group" --description "A group of mine" --parent-group-id <group-id> --custom-fields "id=<some-field>,value=<some-value>" "id=<another-field>,value=<another-value>"
```

**Note:** Pay attention to how we set custom fields. According to the command help, the `--custom-fields` argument accepts an array of objects with 2 keys each: `id` and `value`. The tool interprets multiple space-separated values as an array and each object can be specified using the `key1=value1,key2=value2,..`-notation, which is described in more detail further in the document.

####Create a server:

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

####Wait until the server has been created:

```
clc wait
```

####Query only servers with status "active" and see the output as a table:

```
clc server list --all --filter status=active --output table
```

####Increase the server's CPUs count and log the HTTP request/response data:

```
clc server update --server-id <server_id> --cpu 2 --trace
```

####Show billing details of the servers of the group as a table:

```
clc group get-billing-details --group-id <group-id> --query groups.<group-id>.servers --output table
```

####Make a skeleton of a command for getting groups with servers inside:

```
clc group list --filter 'servers-count>0' --generate-cli-skeleton > groups_with_servers.json
```

####Apply the skeleton:

```
clc group list --from-file groups_with_servers.json
```

## Autocomplete

Autocomplete currently works for:

* Resources (server, data-center, group, etc)
* Commands
* Command options and arguments
* The `--output` values
* The `--profile` values
* Values of the arguments that have a limited set of possible values (like `server create --type standard|hyperscale|bareMetal`)
* Values of the arguments that are actually entity names

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

## Contributing

Development is set up for Unix/Linux/Mac systems. Some of the
instructions below may not work properly on Windows.

### Preparing an environment

* [Install Go](https://golang.org/).
* Install Godep: `go get github.com/tools/godep`.
* Clone this repo (do **not** use `go get`).
* [Ensure your $GOPATH is set correctly](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable). Working in a clean environment without any other packages on $GOPATH is highly encouraged to avoid conflicts with the dependencies of the tool. Using a [gvm tool](https://github.com/moovweb/gvm) is a good choice for setting up a clean environment. Also, gvm is a convenient tool for installing cross-compilation prerequisites.
* Install dependencies with Godep: enter the repo's root and `godep restore`.
* Install go vet: `go get code.google.com/p/go.tools/cmd/vet`.

### Developing

* The TDD approach is recommended - write a failing test first, then fix it.

* Use a `dev` script to run commands as you change the code:

    ```
    ./dev <resource> <command>
    ```

  This way you do not need to rebuild the tool every time you alter something.

* Before making a pull request check that `gofmt -d=true ./..` and `go vet ./...` do not produce any output (except for that coming from `Godeps/_workspace` - ignore it).

* Do not commit until the unit tests have passed (`./run_tests`).

* If you want to make an executable, simply run `./scripts/build`. The binary will appear in the `./out` folder.

* The integration tests can be running `./run_integration_tests`.

* The API file can be regenerated by running `./scripts/generate_api`.

### Building the releases

Generally, any Linux/Darwin machine should work for building the releases. A Darwin machine is required though if you want to build a `MacOS .pkg`.

* Install [gvm](https://github.com/moovweb/gvm)

* Install the cross-compilation prerequisites:

```
./scripts/install_platform_commands
```

* Build the releases:

```
./scripts/build_releases <version>
```

At first, the script updates `base/constants.go` with the given version. This is needed for the tool to use the relevant user agent information. After that, the script builds a binary for each of the following OS/arch flavors:

- Linux/amd4
- Windows/amd64
- MacOS/amd64

The binaries are packaged along with utility scripts as described in the [Install](#install-it) section. The folders are then archived - a `.tar.gz` file is made for Linux and Mac; a `.zip` file is made for Windows.

Here is a full list of the created artifacts:

* `clc-$version-linux-amd64/`
* `clc-$version-linux-amd64.tar.gz`
* `clc-$version-darwin-amd64`
* `clc-$version-darwin-amd64.tar.gz`
* `clc-$version-windows-amd64/`
* `clc-$version-windows-amd64.zip`

#### Building a .pkg for MacOS

* Build a regular MacOS release using the command from the previous section

* Execute the following script to build a `.pkg` file:

```
./scripts/build_darwin_pkg <version>
```

**Note:** the version has to match the version you specified in the previous section.

You should see 2 artifacts after executing this script:

* `clc-$version-pkg`
* `clc-$version.pkg`

## Security

The CenturyLink Cloud Go CLI leverages our public API that serves all requests over HTTPS. Therefore credentials are encrypted when being transfered. Credentials stored on the local machine in the config.yml are not encrypted during installation and you are encouraged to use industry standard encryption tools in order to provide additional protection for them.

## License

The project is licensed under the [Apache License v2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
