package ipify

import (
	"fmt"
	"runtime"
)

// The maximum amount of tries to attempt when making API calls.
const MAX_TRIES = 3

// This is the ipify service base URI for IPV4 resolution.
const API_URI_4 = "https://api.ipify.org"

// This is the ipify service base URI for IPV4 or IPV6 resolution, whichever comes first.
const API_URI_64 = "https://api64.ipify.org"

// This is the ipify service base URI for IPV6 resolution.
const API_URI_6 = "https://api6.ipify.org"

// USER_AGENT string is provided so that I can (eventually) keep track of
// what libraries to support over time.  EG: Maybe the service is used
// primarily by Windows developers, and I should invest more time in improving
// those integrations.
var USER_AGENT = fmt.Sprintf("ipify/%s go/%s %s", VERSION, runtime.Version()[2:], runtime.GOOS)
