package main

// in variadic functions, we can pass variable number of arguments to a function but only of same type
import "fmt"
func sum(nums ...int) int  {
	total:=0
	for _, num:= range nums{
		total += num
	}
	return total
}


func main(){
	slice:= []int{1,2,3,4,5,6}
	fmt.Println(sum(slice...)) //unpacking the slice
}