package main

/**
 * @author Arturo Negreiros
 * @origin Black hat with Go
 */

/*
The first step in creating the port scanner is understanding how
to initiate a connection from a client to a server. Throughout
this example, you’ll be connecting to and scanning
scanme.nmap.org, a service run by the Nmap project. To do
this, you’ll use Go’s net package: net.Dial(network, address string).
The first argument is a string that identifies the kind of
connection to initiate. This is because Dial isn’t just for TCP; it
can be used for creating connections that use Unix sockets,
UDP, and Layer 4 protocols that exist only in your head (the
authors have been down this road, and suffice it to say, TCP is
very good). There are a few strings you can provide, but for
the sake of brevity, you’ll use the string tcp.
The second argument tells Dial(network, address string) the host
to which you wish to connect. Notice it’s a single string, not a
string and an int. For IPv4/TCP connections, this string will take
the form of host:port. For example, if you wanted to connect to
scanme.nmap.org on TCP port 80, you would supply
*/

import (
	"fmt"
	"net"

	//"net"
	"sync"
)

func workers(ports chan int, wg *sync.WaitGroup) {

	for p := range ports {

		fmt.Println(p)
		wg.Done()
	}
}

func VeryFastScanner(host string) {

	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("192.168.100.33:%d", j)
			connection, err := net.Dial("tcp", address)

			if err != nil {
				return
			}
			connection.Close()
			fmt.Printf("%d\n open", j)
		}(i)

		wg.Wait()
	}
}

func main() {

	ports := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		go workers(ports, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
