package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) { // wait groups are noCopy so you need to use a pointer
	for port := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port

	}
}

func main() {
	ports := make(chan int, 100) // buffered channel, doesn't block until it gets 100 items
	results := make(chan int)
	var openPorts []int 

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func ()  {
		for i := 1; i <= 1024; i++ {
			ports <- i;
		}
	}()
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)

	for _, port := range openPorts {
		fmt.Println(port, "open")
	}

 }