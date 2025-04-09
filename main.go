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
		go handleConnection(conn)//addded go routine
	}
} //infinite loop


func handleConnection(conn net.Conn) {
	defer conn.Close()
	//buf := make([]byte, 1024)
	//add the log connection for the timestamp

	// for {
	// 	n, err := conn.Read(buf)
	// 	if err != nil {
	// 		fmt.Println("Error reading from client:", err)
	// 		return
	// 	}
	// 	_, err = conn.Write(buf[:n])
	// 	if err != nil {
	// 		fmt.Println("Error writing to client:", err)
	// 	}
	// } remove these and implement mathods-2-9
}

//run the function using:
// using main.go and then open in web
//open a new terminal and run nc localhost 4000
