## About ##

This is a short paper about a couple of important traits of the tool design worth reading if you want to contribute to the project.

## Defining Commands ##

Commands are defined in `init.go`. Each definition has the following structure (please, read through the comments to get an idea of what each of the parameters is used for):

```
registerCommandBase(InputModel, OutputModel, commands.CommandExcInfo{ // input model and output model are explained further in the doc
    Verb: "POST", // "GET", "PUT", "PATCH", "DELETE"
    Url: "https://api.ctl.io/v2/myresources",
    Resource: myresource,
    Command: mycommand, // the invocation would start with "clc myresource mycommand"
    Help: help.Command{
        Brief: []string{ // this will be shown as a part of the "clc myresource --help" output
            "Applies my command on my resource",
            "This text will be displayed starting from a new line",
        },
        Arguments: []help.Argument{ // this will be shown as a part of the "clc myresource mycommand --help" output
            {
                "--my-argument",
                []string{
                    "Required. Applies an argument.",
                    "This text will be displayed on a new line",
                },
            },
        },
    },
})
```
### Input Model ###

Input model is constructed out of a user input, if any, serialized into JSON, and attached to the API request as a payload and/or as a GET query string.

Generally, a pointer to a struct is passed to `registerCommandBase` as an input model.

If an API call does not expect any payload use `nil`.

Each console argument corresponds to a struct field. For example, `--my-field` corresponds to the struct field with a name `MyField`.

### URL Interpolation ###

You can interpolate the URL templates with the input values.

Say, you have a model:

```

type MyInputModel struct {
    MyParam string
}
```

And you want the `MyParam` value to be a part of the URL. Then, there are 3 rules to follow:

* The field (`MyParam` in this case) has to be string.
* The field has to be marked with a `URIParam` tag like this: ``MyParam string `URIParam:"yes"` ``.
* The URL may refer to the field's value as `{MyParam}`.

`{MyParam}` may appear anywhere in the URL including the query part:

```
https://api.ctl.io/v2/myresources/{MyResourceId}?myparam={MyParam}&onemoreparam={OneMoreParam}
```

Note that if all query parameters are empty strings the query is not constructed at all.

Moreover, you can refer to a nested field of a composed structure:

```
type MyComposedModel struct {
    MyInputModel `argument:"composed" URIParam:"MyParam"`
}
```

The `argument` tag is described in details further in the document.

An `accountAlias` parameter gets substituted for you. You do not usually need to create the corresponding model field.

### Supported Tags ###

Please, read through the comments to get to know various supported struct tags. 

```
type MyInputModel struct {
  ThisIsRquired string `valid:"required"` // a non-empty value is required; the tag is from https://github.com/asaskevich/govalidator
  Enumerable string `oneOf:"X,Y,Z"` // will fail unless Enumerable is either X or Y or Z
  QueryParam string `URIParam:"yes"` // see the URL Interpolation section; only strings are supported
  InnerModel `argument:"composed" URIParam:"InnerId,InnerCode"` // treats the inner model as if its fields were the outer model's fields; the URIParam tag here grabs the fields `InnerId` and `InnerCode` out of the inner model
  NoAPI string `json:"-"`
  NoInput string `json:"ImplicitParam" argument:"ignore"` // there is no `--no-input` console argument but the field gets serialized to JSON
}
```

### Custom Input Model Behaviour ###

#### Custom Validation ####

The following interface lets you perform custom validation:

```
func (m *MyModel) Validate() error {
    // Check whatever you have to here.
    // Use `errors.EmptyField` for consistency if you want to report a missing value for a field of some custom data type.
}
```

#### Custom Behaviour ####

Model post-processing can be performed **after** the validation via the following interface:

```
func (m *MyModel) ApplyDefaultBehaviour() error {
    // For example, convert a string to upper case or apply an unusual format to a time value
}
```

#### Inferring Entity Names and Autocomplete Options ####

The interface below is used to provide users with a possibility to specify a name (or another concise parameter) instead of ID (or another non-human-readable parameter). There is a number of common models implementing this interface which you can embed into other models including:

* `server.Server`
* `network.Network`
* `group.Group`
* `balancer.LoadBalancer`

The behaviour is also used in autocomplete where the list of values of a given property is often needed.

```
func (m *MyModel) InferID(cn base.Connection) error {
    // Load entities from the server, find the target entities by some set of specified parameters, and assign the missing ones (usually, ID).
}

func (m *MyModel) GetNames(cn base.Connection, property string) ([]string, error) {
    // Load entities and collect a list of the values of the given property (usually, name)
}
```

### Output Model ###

Output model is an object the response payload is loaded into. Note that you should pass a pointer into `registerCommandBase` in `init.go`.

JSON is normally loaded into a struct or a slice of structs. When no response payload is expected use `new(string)` - it will hold an error message in case of a failure.

