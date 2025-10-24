package main

import "fmt"

func main() {
	age := 20

	if age < 18 {
		fmt.Println("Minor")
	} else {
		fmt.Println("Adult")
	}
}