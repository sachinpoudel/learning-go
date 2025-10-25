package main

// enums in go are used to define a set of named constants that represent a specific type.
// Go does not have a built-in enum type like some other programming languages, but we can achieve similar functionality using the "iota" keyword.
// The "iota" keyword is used to create a sequence of related constants that are automatically incremented.

import "fmt"
type orderstatus string

const (
	Received orderstatus = "done"
	Prepared = "prepared"
)
func changeOrderStatus(status orderstatus) {
	fmt.Println(status)
}

func main(){ 
	changeOrderStatus(Prepared)
}
