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
	flagPort := flag.String("flagPort", "4000","Listens for connections on port: ")
	flag.Parse()
	ListenON := fmt.Sprintf(":%s", *flagPort)
	fmt.Println("Server is listening on", ListenON)

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
	

	Path := strings.ReplaceAll(ADR, ":", "_") + ".log"
	records, err := os.OpenFile(Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error: ", err)
		return
	}
	defer records.Close()

	timeout := 30 * time.Second 
	inputread := bufio.NewScanner(conn)
	Handle  := 1024    
	for {
	conerr := conn.SetReadDeadline(time.Now().Add(timeout))
	if conerr != nil{
	fmt.Printf("Cannot set deadline, connection timeout: ", conerr)
	return
	}

	if !inputread.Scan() {
	Printlog(fmt.Sprintf("Connection disconnected: %s",ADR))
	return
	}

	input := strings.TrimSpace(inputread.Text())
	records.WriteString(fmt.Sprintf("%s\n", input))	

	if len(input) > Handle {
	conn.Write([]byte("Truncated the message, due to its length "))
	input = input[:Handle]
	}

	if input == "" {
	conn.Write([]byte("Say something...\n"))
	}else if input == "hello" {
	conn.Write([]byte("Hi there!\n"))
	} else if input == "bye" {
	conn.Write([]byte("Goodbye!\n"))
	return
	}


	
	}
	
	
	



}







//run the function using:
// using main.go and then open in web
//open a new terminal and run nc localhost 4000


//task: 
//- main (code - completed), (comments - incomplete)
//- Hanlde function (code - incomplete), (comments - incomplete)
