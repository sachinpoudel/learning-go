package main
import "fmt"
func main(){


	var arr[5]int = [5]int{1,2,3,4,5}
	fmt.Println("Array:",arr)
	arr[4] = 20
	fmt.Println("Updated Array:",arr)


names := [4]string{"Alice","Bob","Charlie","Alex"}
for i:=1; i<len(names); i++{
	fmt.Println("Name",i,":",names[i])	

}
}