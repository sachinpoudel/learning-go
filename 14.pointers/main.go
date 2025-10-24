package main

import "fmt"
func changenum(num *int){
	*num=10
	fmt.Println("Inside changenum:",*num)
}
func main() {

	num := 1
	changenum(&num)
	fmt.Println(&num)
fmt.Println("Inside main:",num)
}
