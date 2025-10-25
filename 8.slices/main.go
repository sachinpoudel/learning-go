package main
import "fmt"
import "slices"

//slices in go are used to represent a sequence of elements of the same type.
// A slice is a dynamically-sized, flexible view into the elements of an array.
// Slices are more powerful than arrays because they can grow and shrink in size as needed.
// Slices are defined using square brackets [] followed by the type of elements they contain.
func main() {
	var fruits []string
	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana")
	fruits = append(fruits, "Cherry")
	fmt.Println("Fruits:", fruits)
	fmt.Println(len(fruits))





var nums = make([]int, 3,6) //length 3 makes slices with 3 elements and capacity 6
nums[0] = 10
nums[1] = 20
nums[2] = 30
fmt.Println("Numbers:", nums)

fmt.Println(cap((nums))) //capacity
nums = append(nums, 40)
nums = append(nums, 60)
nums = append(nums, 70)
nums = append(nums, 80)
fmt.Println("Numbers after append:", nums)
fmt.Println(cap((nums))) //capacity

//slice operator 
mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
subSlice := mySlice[2:6] //elements from index 2 to 5
fmt.Println("Sub-slice:", subSlice)



slice := []int{1,2,3}
slice2 := []int{1,2,3}



fmt.Println(slices.Equal(slice,slice2))

} //slice = dynamic array
