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
	flagPort := flag.String("flagPort", "4000", "Listens for connections on port: ")
	flag.Parse()

	ListenON := fmt.Sprintf(":%s", *flagPort)
	fmt.Printf("Server is listening on %s\n", ListenON)

	listener, err := net.Listen("tcp", ListenON)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n ", ListenON)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error! Can't connect to network: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
} //infinite loop

func handleConnection(conn net.Conn) {
	defer conn.Close()

	ADR := conn.RemoteAddr().String()
	Print(fmt.Sprintf("A new connection:%s - %s", ADR, time.Now().Format("2006-01-02 15:04:05")))

	Path := strings.ReplaceAll(ADR, ":", "_") + ".log"
	records, err := os.OpenFile(Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error! Can't open the log file: %v", err)
		return
	}
	defer records.Close()

	timeout := 30 * time.Second
	inputread := bufio.NewScanner(conn)

	for {
		conerr := conn.SetReadDeadline(time.Now().Add(timeout))
		if conerr != nil {
			fmt.Printf("Cannot set deadline: %s\n ", conerr)
			return
		}

		if !inputread.Scan() {
			if err := inputread.Err(); err != nil {
				Print(fmt.Sprintf("Read timeout or error from %s: %v", ADR, err))
			} else {
				Print(fmt.Sprintf("Connection disconnected: %s", ADR))
			}
			return
		}

		input := strings.TrimSpace(inputread.Text())
		if _, err := records.WriteString(input + "\n"); err != nil {
			fmt.Printf("Error! Can't log message : %v\n", err)
		}
		if len(input) > 1024 {
			input = input[:1024]
			conn.Write([]byte("Truncated the message, due to its length.\n"))
		}

		switch {
		case input == "":
			conn.Write([]byte("Say something...\n"))

		case input == "hello":
			conn.Write([]byte("Hi there!\n"))

		case input == "bye":
			conn.Write([]byte("Goodbye!\n"))
			Print(fmt.Sprintf("Connection from %s closed with 'bye'", ADR))
			conn.Close()
			return

		case input == "/time":
			conn.Write([]byte(time.Now().Format(time.RFC1123) + "\n"))

		case input == "/quit":
			conn.Write([]byte("Closed connection.\n"))
			Print(fmt.Sprintf("Connection from %s closed with '/quit'", ADR))
			conn.Close()
			return

		case strings.HasPrefix(input, "/echo"):
			conn.Write([]byte(input[6:] + "\n"))

		default:
			conn.Write([]byte(input + "\n"))
		}
	}
}

func Print(msg string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
}



//task:
//- main (code - completed), (comments - incomplete)
//- Hanlde function (code - complete), (comments - incomplete)
