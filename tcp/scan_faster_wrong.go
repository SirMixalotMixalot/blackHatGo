package main

import (
	"fmt"
	"net"
)
func isValidPort(port int, validPortChannel chan int) {
	address := fmt.Sprintf("scanme.nmap.org:%d", port);
	conn, err := net.Dial("tcp", address);
	if err != nil { // port closed or filtered
		return
	}
	conn.Close() // do not use defer in these cases!
	//defer would wait until main is done or a panic occurs
	// if you want to use defer, enclose this all in a seperate function
	validPortChannel <- port;

}
func main() {
	portChannel := make(chan int)
	for port := 1; port <= 1024; port++ {
		go isValidPort(port, portChannel)
	}

	for port := range portChannel { // the channel is never closed so this doesn't work and hangs forever

		fmt.Println(port, " is a valid port")
	}

}