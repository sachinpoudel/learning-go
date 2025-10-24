package main
import "fmt"

func main (){


 maps:= make(map[string]int) // map gives key value pair
maps["alice"] = 25 // like map[alice:25]
maps["bob"] = 30
maps["charlie"] = 35

clear(maps)
 fmt.Println(maps)






}
