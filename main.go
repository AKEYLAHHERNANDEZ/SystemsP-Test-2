//creating a tcp echo server 
package main

import(
	"fmt"
	"net"
)

func main(){
	listener, err := net.Listen("tcp",":4000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

fmt.Println("Server listening on :4000")
for {
	conn,err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err)
		continue
	}
	handleConnection(conn)
}
}//infinite loop

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
_, err = conn.Write(buf[:n])
if err !=nil {
	fmt.Println("Error writing to client:",err)
}
	}
}

//run the function using:
// using main.go and then open in web
//open a new terminal and run nc localhost 4000