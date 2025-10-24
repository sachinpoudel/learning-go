package main

import "fmt"

// func getLang () (string,string, int) {
// 	return "Golang", "Java", 10000
// }
func process(fn func(a int) int) int {
	return fn(600)
}


func main (){
// lang1, lang2, _ := getLang()
// fmt.Println(lang1, lang2)
result := process (func(a int) int {
	return a + 1000
})
fmt.Println(result)}