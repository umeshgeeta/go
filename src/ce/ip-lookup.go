// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"net"
	"os"
)

var names = []string{
	"google.com",
	"cnn.com",
	"comcast.com",
	"google.com",
}

var ipsMap map[string][]net.IP = make(map[string][]net.IP)

func main() {
	for _, name := range names {

		found, ok := ipsMap[name]

		if ok {
			fmt.Fprintf(os.Stderr, "Found ips %v for %s \n", found, name)
		} else {
			// we will have to do the DNS Lookup
			ips, err := net.LookupIP(name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
				os.Exit(1)
			}
			// there is no error here, we would have found entries
			ipsMap[name] = ips
			for _, ip := range ips {
				fmt.Printf("%s. IN A %s\n", name, ip.String())
			}
		}

	}
}
