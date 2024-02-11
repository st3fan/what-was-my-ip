package ipify

import (
	"errors"
)

// A struct to retrieve your IP address in any way you would like.
// You can define your own functions to call that can be used
// to retrieve your IP address.
type Ip struct {
	Ip4f func() (string, error) // Function to retrieve IPV4 address.
	Ip6f func() (string, error) // Function to retrieve IPV6 address.
	Ip4u string                 // URL to use to retrieve IPV4 address.
	Ip6u string                 // URL to use to retrieve IPV6 address.
	Ip4  string                 // IPV4 address once retrieved.
	Ip6  string                 // IPV6 address once retrieved.
	Ip4e error                  // Error value if unable to retrieve IPV4 address.
	Ip6e error                  // Error value if unable to retrieve IPV6 address.
}

// Initialize a default Ip struct with appropriate functions and the URL's necessary to retrieve your IP addresses,
// IPV4 and IPV6.
func DefaultIp() *Ip {
	ips := &Ip{
		Ip4u: API_URI_4,
		Ip6u: API_URI_6,
	}
	ips.Ip4f = func() (string, error) {
		return getIp(ips.Ip4u)
	}
	ips.Ip6f = func() (string, error) {
		return getIp(ips.Ip6u)
	}
	return ips
}

// Retrieve IPV4 and IPV6 addresses, or only one of them,
// depending on which function is nil.
func (i *Ip) Retrieve() error {
	if i.Ip4f == nil && i.Ip6f == nil {
		return errors.New("ip functions are nil, unable to retrieve addresses")
	}
	i.r4()
	i.r6()
	if i.Ip4e != nil && i.Ip6e != nil {
		return errors.New(i.Ip4e.Error() + "\n" + i.Ip6e.Error())
	}
	if i.Ip4e != nil && i.Ip6f == nil {
		return i.Ip4e
	}
	if i.Ip6e != nil && i.Ip4f == nil {
		return i.Ip6e
	}
	return nil
}

func (i *Ip) r4() {
	if i.Ip4f != nil && i.Ip4 == "" && i.Ip4e == nil {
		i.Ip4, i.Ip4e = i.Ip4f()
	}
}

func (i *Ip) r6() {
	if i.Ip6f != nil && i.Ip6 == "" && i.Ip6e == nil {
		i.Ip6, i.Ip6e = i.Ip6f()
	}
}
