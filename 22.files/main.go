package main

//files in go are used to read and write files from the filesystem.
// The "os" and "io/ioutil" packages provide functions for working with files.
import (
	"fmt"
	"os"
)
func main(){

f, err:=os.Open("example.txt")
if err != nil {
	panic(err)

}

fileInfo,err:=f.Stats()
if err !=nil {
	panic(err)
}
fmt.Println(fileInfo.Name())
}

