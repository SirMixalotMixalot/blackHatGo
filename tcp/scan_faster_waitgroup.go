package main

import (
	"fmt"
	"net"
	"sync"
)

//technically correct but opening those connections all the time leads to errors that skew our results
//can't open 1024 concurrent connections to 1024 hosts
func main() {
	var wg sync.WaitGroup
	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", port);
			conn, err := net.Dial("tcp", address);
			if err != nil { // port closed or filtered
				return
			}
			conn.Close() // do not use defer in these cases!
			//defer would wait until main is done or a panic occurs
			// if you want to use defer, enclose this all in a seperate function
			fmt.Println(port, " is a valid port")
		}(port)
	}
	wg.Wait()



}