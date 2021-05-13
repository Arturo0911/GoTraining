package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	host    = flag.String("host", "192.168.100.215", "host or ip to scann")
	port    = flag.String("range", "1-65535", "range ports to scann")
	threads = flag.Int("threads", 1000, "threads to use in scanner")
	timeout = flag.Duration("timeout", 1*time.Second, "seconds per port")
)

func processRange(ctx context.Context, r string) chan int {
	c := make(chan int)
	done := ctx.Done()

	go func() {
		defer close(c)
		blocks := strings.Split(r, ",")

		for _, block := range blocks {
			rg := strings.Split(block, "-") // if i found a - then split it
			var minPort, maxPort int
			var err error

			minPort, err = strconv.Atoi(rg[0])
			if err != nil {
				log.Print("Its not possible convert the port: \n", block)
				continue
			}

			if len(rg) == 1 {
				maxPort = minPort
			} else {
				maxPort, err = strconv.Atoi(rg[1])
				if err != nil {
					log.Print("It's not possible parse the range: \n", block)
					continue
				}
			}

			for port := minPort; port <= maxPort; port++ {
				select {
				// declaration
				// allow a goroutine in multiple operations
				case c <- port:
				case <-done:
					return
				}
			}
		}
	}()
	return c
}

func scanPorts(ctx context.Context, in <-chan int) chan string {

	out := make(chan string)
	done := ctx.Done()

	var wg sync.WaitGroup
	wg.Add(*threads)

	for i := 0; i < *threads; i++ {
		go func() {
			defer wg.Done() // To close the threads
			for {
				select {
				case port, ok := <-in:
					if !ok {
						return
					}
					s := scanPort(port)
					select {
					case out <- s:
					case <-done:
						return
					}
				case <-done:
					return
				}
			}
		}()

		go func() {
			wg.Wait()
			close(out)
		}()
	}
	return out
}

func scanPort(port int) string {
	addr := fmt.Sprintf("%s:%d", *host, port)
	conn, err := net.DialTimeout("tcp", addr, *timeout)
	if err != nil {
		return fmt.Sprintf("%d: %s", port, err.Error())
	}

	conn.Close()

	return fmt.Sprintf("%d: Open!", port)

}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // seting context

	defer cancel()
	flag.Parse()
	fmt.Printf("\n[*] Scanning host %s (ports: %s)\n\n", *host, *port)

	pR := processRange(ctx, *port)
	sP := scanPorts(ctx, pR)

	for port := range sP {
		if strings.HasSuffix(port, ": Open!") {
			fmt.Println(port)
		}
	}
}
