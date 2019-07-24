// Copyright 2013-2019 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Get the time the machine has been up
// Synopsis:
//     webboot
package main

import (
	"flag"
	"log"

	"github.com/u-root/webboot/pkg/dhclient"
)

var (
	ifName  = flag.String("interface", "^e.*", "Name of the interface")
	timeout = flag.Int("timeout", 15, "Lease timeout in seconds")
	retry   = flag.Int("retry", 5, "Max number of attempts for DHCP clients to send requests. -1 means infinity")
	verbose = flag.Bool("verbose", false, "Verbose output")
	ipv4    = flag.Bool("ipv4", true, "use IPV4")
	ipv6    = flag.Bool("ipv6", true, "use IPV6")
)

func main() {
	flag.Parse()

	dhclient.Request(*ifName, *timeout, *retry, *verbose, *ipv4, *ipv6)
	log.Println("weboot up")
}
