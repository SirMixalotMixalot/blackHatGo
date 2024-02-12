package main

import (
	"fmt"
	"net"
)

func main() {
	validPorts := make([]int, 10)
	for port := 1; port <= 1024; port++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", port);
		conn, err := net.Dial("tcp", address);
		if err != nil { // port closed or filtered
			continue
		}
		conn.Close() // do not use defer in these cases!
		//defer would wait until main is done or a panic occurs
		// if you want to use defer, enclose this all in a seperate function
		validPorts = append(validPorts, port)
		fmt.Println(port, " is a valid port")

	}
	fmt.Println(validPorts)
}