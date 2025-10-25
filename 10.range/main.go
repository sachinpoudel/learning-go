package main

import "fmt"
//ranges in go are used to iterate over elements in a variety of data structures such as arrays, slices, maps, strings, and channels.
func main(){
	nums:= []int{1,2,3,4,5,6,7,8,9,10}

	

	for i, num:= range nums{
		
		fmt.Println(i,num)
	}
}
