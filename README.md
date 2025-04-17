# SystemsP-Test-2

My TCP Echo Server that listens for client connections, logs all received messages to a file &  handles custom commands.

Link to Youtube demo video: https://youtu.be/vG0LbMo_uls


How to run the server :
To run the  default server use the command: go run main.go 
Or to run a server with a specific port use the command: go run main.go -flagPort=7000 
After running the server, open a new terminal and connect the client using the command: nc localhost 4000 

Commands to run on the Client: 
input: hello     output: Hi there!
input: bye     output:  Goodbye!
input: /time     output: Returns the current time
input: /quit     output: Closed connection.
input: /echo anymessage  output: anymessage
input: An empty message     output: Say somthing...
input: Long text entered     output: Truncated the message, due to its length.
Input: no input (waiting)   otuput: Server disconnects after 30 seconds


Which functionality was the most educationally enriching:

One of the most enriching parts was implementing error handling and file logging. Writing logs to files, handling disconnects and giving useful server feedback was interesting to learn and implement. It also helped me understand how servers manage client connections and input validation, it also  taught me that debugging is important and of course, using proper error handling.


Which functionally required that you do the most research:

The most research went into implementing SetReadDeadline, creating log files, and understanding error handling. Figuring out how to properly set timeouts for inactive clients, generate log files for each user and fix errors at each step was a bit complex. I spent a lot of time figuringout Go documentation and troubleshooting various commands to make it all work smoothly.



