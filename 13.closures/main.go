package main
import "fmt"

// closure in go is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables;
// in this sense the function is "bound" to the variables.

func counter() func() int {
	var count = 0

	return func() int {
		count++
		return count
	}
}

func main(){
// 	c := counter()
// 	fmt.Println(c())
// 	fmt.Println(c())
// 	fmt.Println(c())
// }
fmt.Println(counter()())


}
