# ipify

An unofficial client library for [ipify][]: *A Simple IP Address API*. Based on the official library.


## Purpose

[ipify][] is the best IP address lookup service on the internet.  It's fast, simple, scalable, and open source.

In short: if you need a way to pragmatically get your public IP address, ipify is the best possible choice!

This library will retrieve your public IP address from ipify's API service, and return it as a string.  It can't get any simpler than that. Usage documentation is below.

This library also has some other nice features you might care about:

- If a request fails for any reason, it is re-attempted 3 times using an exponential backoff algorithm for maximum effectiveness.
- This library handles errors properly, and usage examples below show you how to deal with errors in a foolproof way.
- This library only makes API requests over HTTPS.


## Installation

To install `ipify`, simply run:

```console
$ go get github.com/tech10/ipify
```

This will install the latest version of the library automatically.


## Usage

Using this library is very simple.  Here's a simple example:

```go
package main

import (
	"fmt"
	"github.com/tech10/ipify"
)

func main() {

	// Initialize variables.
	// We will be resolving to IPV4, IPV6, or which ever comes first.
	// This will fully demonstrate the use of the library.

	var ip string
	var ip4 string
	var ip6 string
	var err error

	// This will resolve to IPV4 or IPV6, which ever comes first.
	ip, err = ipify.GetIp()
	if err != nil {
		fmt.Println("Couldn't get my IP address:", err)
	} else {
		fmt.Println("My IP address is:", ip)
	}

	// For IPV4 only.
	ip4, err = ipify.GetIp4()
	if err != nil {
		fmt.Println("Couldn't get my IPV4 address:", err)
	} else {
		fmt.Println("My IPV4 address is:", ip4)
	}

	// For IPV6 only.
	ip6, err = ipify.GetIp6()
	if err != nil {
		fmt.Println("Couldn't get my IP address:", err)
	} else {
		fmt.Println("My IPV6 address is:", ip6)
	}
}
```

In regards to error handling, there are several ways this can fail:

- The ipify service is down (*not likely*), or:
- Your machine is unable to get the request to ipify because of a network error of some sort (*DNS, no internet, etc.*).
- If attempting to get IPV6 or IPV4 addresses only, your machine is unable to either resolve or connect to the API. For instance, if your network connection doesn't support IPV6, you will be unable to use the GetIp6 function.

The library will output an informative error message in the event anything fails.

One thing to keep in mind: regardless of how you decide to handle errors, the ipify library will retry any failed requests 3 times before failing, so if you *do* need to handle errors, remember that retry logic has already been attempted.


## Contributing

If you'd like to improve this library, please send me a pull request.  I'll be happy to review them, and if they seem like a good idea, merge them.

The standard contribution workflow should look something like this:

- Fork this project on Github.
- Make some changes in a feature branch.
- Send a pull request when ready.

Also, if you're making changes, please write tests for your changes. This project has a full test suite you can easily modify / test.

To run the test suite, you can use the `go test` command after forking this repository.


  [ipify]: http://www.ipify.org/ "ipify - A Simple IP Address API"
