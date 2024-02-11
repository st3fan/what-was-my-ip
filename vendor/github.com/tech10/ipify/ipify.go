// Package ipify provides multiple methods for retrieving your computer's
// public IP address from the ipify service: http://www.ipify.org
// While it uses ipify by default, you are free to use any service you like.
package ipify

// GetIp queries the ipify service (http://www.ipify.org) to retrieve this
// machine's public IP address.  Returns your public IP address as a string, and
// any error encountered.  By default, this function will run using exponential
// backoff -- if this function fails for any reason, the request will be retried
// up to 3 times.
//
// This function will return either the IPV4 or IPV6 address, which ever resolves first.
//
// Usage:
//
//		package main
//
//		import (
//			"fmt"
//			"github.com/tech10/ipify"
//		)
//
//		func main() {
//			ip, err := ipify.GetIp()
//			if err != nil {
//				fmt.Println("Couldn't get my IP address:", err)
//			} else {
//				fmt.Println("My IP address is:", ip)
//			}
//		}
func GetIp() (string, error) {
	return getIp(API_URI_64)
}

// GetIp4 returns the IPV4 address of your computer.
// Use this as you would use the GetIp() function.
func GetIp4() (string, error) {
	return getIp(API_URI_4)
}

// GetIp6 returns the IPV6 address of your computer.
// Use this as you would use the GetIp() function.
func GetIp6() (string, error) {
	return getIp(API_URI_6)
}
