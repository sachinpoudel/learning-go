package main
import "fmt"
//maps in go are used to store key-value pairs.
// A map is an unordered collection of key-value pairs where each key is unique.
// Maps are defined using the "map" keyword followed by the key type in square brackets [] and the value type.
func main (){

//make func is used to create maps, slices, and channels in Go.
 maps:= make(map[string]int) // map gives key value pair
maps["alice"] = 25 // like map[alice:25]
maps["bob"] = 30
maps["charlie"] = 35

clear(maps)
 fmt.Println(maps)






}
