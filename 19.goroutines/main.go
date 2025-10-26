package main

import (
	"fmt"
	"sync"
)

// routines in go are lightweight threads managed by the Go runtime.
// Goroutines are functions or methods that run concurrently with other functions or methods.
// They are created using the "go" keyword followed by a function call.

func task(id int, w *sync.WaitGroup){
	defer w.Done()
	fmt.Println(id)
}
// wait group is used to wait for a collection of goroutines to finish executing.						
func main(){

	var wg sync.WaitGroup

	for i:= 1; i<20; i++{
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}
// goroutine runs concurrently means independently with other goroutines. 
// that is if a goroutine is blocked or waiting, other goroutines can continue to run.
//means if a value is printed randomly before or after another value, it is because of goroutine scheduling. mean multiple thread of loop run concurrently so the output may vary in order.
// A goroutine is a lightweight thread managed by the Go runtime. 
// To create a goroutine, use the "go" keyword followed by a function call. 
// This starts a new goroutine that runs concurrently with the calling goroutine.
// For example:
// go myFunction() 
// This will start a new goroutine that executes myFunction() concurrently.	