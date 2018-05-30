# go-modica #
[![Build Status](https://travis-ci.org/matthewhartstonge/go-modica.svg?branch=master)](https://travis-ci.org/matthewhartstonge/go-modica)
[![Go Report Card](https://goreportcard.com/badge/github.com/matthewhartstonge/go-modica)](https://goreportcard.com/report/github.com/matthewhartstonge/go-modica)

go-modica is a Go Client library for accessing [Modicagroup's RESTful APIs.][modica api uri]

go-modica requires Go version 1.8 or greater.

[modica api uri]: https://confluence.modicagroup.com/display/DC/Modica+API+Documentation

## Usage ##

```go
import "github.com/matthewhartstonge/go-modica"
```

Construct a new Modica client, then use the various services on the client to 
access different parts of the Modica API. For example:

```go
client := modica.NewClient("foo", "bar", nil)

// Get a message that has been sent
msg, _ := client.MobileGateway.GetMessage(654321)
```

Some API methods have optional parameters that can be passed. For example:

```go
client := modica.NewClient("foo", "bar", nil)

// Send an SMS message
myCoolNewMessageToSend := &modica.Message{
    Destination: "+642123456789",
    Content:     "Hello, test message!",
}
msg, _ := client.MobileGateway.CreateMessage(myCoolNewMessageToSend)
```

For more sample code snippets, head over to the [example][exampledir] directory.

[exampledir]: https://github.com/matthewhartstonge/go-modica/tree/master/example

### Authentication ###

The go-modica library handles client basic authentication for you and asks for 
these when creating a new api client. You can find your client credentials under:
 
[Omni Dashboard][omnidashboard] > Applications >  Mobile Gateway (API)

* `ClientID`: is the `Application Name` 
* `ClientSecret`: is the generated `password`

```go
client := modica.NewClient("ClientID", "ClientSecret", nil)
```

[omnidashboard]: https://omni.modicagroup.com

## Roadmap ##

This library is being initially developed for an internal application at
LINC-ED, so API methods will likely be implemented in the order that they are
needed by that application. Eventually, I would like to cover the entire
Modica API, so contributions are of course always welcome. The calling pattern 
is pretty well established, so adding new methods is relatively straightforward.

## Versioning ##
In general, go-modica follows [semver](https://semver.org/) as closely as we can 
for tagging releases of the package.

* The major version is incremented with any incompatible change to the exported 
	Go API.
* The minor version is incremented with any backwards-compatible changes to 
	functionality, including new features.
* The patch version is incremented with any backwards-compatible bug fixes.

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
