package main

import "fmt"
//generics in go are used to write flexible and reusable code that can work with different data types.
// Generics allow you to define functions, types, and data structures that can operate on different types without sacrificing type safety.
// Generics are defined using type parameters enclosed in square brackets [].
func printSlice[T any](items []T){
	for _ , item:= range items {
		fmt.Println(item)
	}
}
func main(){	
	printSlice([]int{1,2,3})
	printSlice([]string{"apple", "banana", "cherry"})
}