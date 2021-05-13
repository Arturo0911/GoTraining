package main

import (
	"flag"
	"fmt"
)

var (
	host = flag.String("host", "192.168.100.215", "host or ip to scann")
	port = flag.String("range", "1-65535", "range ports to scann")
)

func main() {

	flag.Parse()
	fmt.Printf("\n[*] Scanning host %s (ports: %s)\n\n", *host, *port)

}
