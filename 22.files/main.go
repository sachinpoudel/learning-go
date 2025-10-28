package main

import (
	"fmt"
	"os"
)


func main(){
f, err := os.Open("example.txt")
 if err != nil {
	panic(err)
 }
// fileInfo , err := f.Stat()
// if err != nil {
// 	panicrr)
// }

// fmt.Println("file name", fileInfo.Name())
// fmt.Println("file name", fileInfo.IsDir())
// fmt.Println("file name", fileInfo.Size())
// fmt.Println("file modified at", fileInfo.ModTime())

defer f.Close()

// buf:= make([]byte, 10)
// d, err:= f.Read(buf)

// if err != nil {
// 	panic(err)
// }

data , err := os.ReadFile("example.txt")

if err != nil {
	panic(err)
}

fmt.Println(string(data))





}