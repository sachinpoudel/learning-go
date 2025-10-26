package main
import "fmt"


//channels in go are used for communication between goroutines.
// A channel is a typed conduit through which you can send and receive values with the channel operator, <-.
// Channels are defined using the "chan" keyword followed by the type of values that will be sent through the channel.

// func processNum(numChan chan int){
// 	fmt.Println("processing", <-numChan)
// }
func emailSender(emailChan chan string, done chan bool){
	defer func() {done <-true}() // here defer is used to ensure that the done channel is closed when the function exits. and func is used to create an anonymous function that closes the done channel.
	//anonymous function is a function that is defined without a name. It is often used as a closure or a callback function.
	for email := range emailChan {
		fmt.Println("Sending email to", email)
	}
}

func main(){
	// messageChan:= make(chan string)
// numChan := make(chan int)
// go processNum(numChan)

// numChan <- 5 // value sent
	// messageChan <- "ping"
	// msg:= <- messageChan
	// fmt.Println(msg)

emailChan := make(chan string, 100)
done:= make(chan bool)
emailChan <- "10@gmail.com" // here we are sending value to channel but no goroutine is receiving it so it will cause deadlock
// emailChan <- "sa@gmail.com"

// fmt.Println(<- emailChan)
// fmt.Println(<- emailChan)

go emailSender(emailChan, done)

}

//deadlock is a situation where two or more goroutines are blocked forever, waiting for each other.
// This can happen when goroutines are trying to send or receive values on channels that are not being read from or written to.
// For example, if a goroutine is trying to send a value on a channel that has no receiver, it will block forever, causing a deadlock.
// To avoid deadlocks, make sure that there is always a corresponding send and receive operation for each channel operation.
// Buffered channels can help prevent deadlocks by allowing a certain number of values to be sent without an immediate corresponding receive.



