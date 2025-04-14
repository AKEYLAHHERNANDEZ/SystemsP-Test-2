# SystemsP-Test-2

My TCP Echo Server that listens for client connections, logs all received messages to a file &  handles custom commands.

Link to Youtube demo video: 


How to run the server :
To run the  default server use the command: go run main.go 
Or run a server with a specific port use the command: go run main.go -flagPort=7000 
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

One of the most enriching functionality was implementing error handling and file logging. Writing logs to files, handling disconnects, and giving useful server feedback was inteesting to leanr and implement, while also learning how  servers manage connections and input validation. It also taught the importance of logging for debugging and ofcourse error handle.



Which functionally required that you do the most research:

The most research went into implementing SetReadDeadline, log files, and the error handling. Understanding how to properly set timeouts for when the client is inactive, creating  log files for each user, and managing errors at each step was complex. I had to understand what go official documentation was and troubleshoot alot of different commands.


