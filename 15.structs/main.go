package main
//struct in go is used to group related data together to form a record.
// A struct is a composite data type that allows you to define a collection of fields, each with its own name and type.
// Structs are defined using the "type" keyword followed by the struct name and the "struct" keyword.
import "fmt"

type customer struct {
	name string

}

type order struct {
	id     string
	amount int
	status string
	address string
	customer 
}

// func newOrder(id string, amount int, status string)*order {
// 	myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }




// //receiver type
// func (o  *order) changeStatus(status string) {
// o.status = status
// } //we only use pointer if the value are change like here in status but not needed when values arent change  

func main(){
	myOrder:= order{  // we can create more myOrder like instance like myOrder2 ,3 and so on 
	id: "12",
	amount: 200,
	status: "received",
	customer: customer{name: "sachin"},
}
// myOrder.address = "Nepal"
// fmt.Println(myOrder)

// myOrder := newOrder("12", 30000, "done")
// myOrder.changeStatus("confirmed")
myOrder.customer.name = "paudel"
fmt.Println(myOrder)














}