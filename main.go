// creating a tcp echo server
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":4000")//change to an address
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server listening on :4000")//change to an address
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		go handleConnection(conn)
	}
} //infinite loop


func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	ADR := conn.RemoteAddr().String()
	Printlog(fmt.Sprintf("A new connection:%s - %s", ADR, time.Now().Format("2025-04-09 13:04:10")))
	

	inputread := bufio.NewScanner(conn)
	if !inputread.Scan() {
	Printlog(fmt.Sprintf("Connection disconnected: %s",ADR))
	}
}







//run the function using:
// using main.go and then open in web
//open a new terminal and run nc localhost 4000
