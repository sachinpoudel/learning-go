package main
 // interfaces in go are used to define a set of method signatures that a type must implement to satisfy the interface.
 // An interface type is defined using the "type" keyword followed by the interface name and the "interface" keyword.
 // in simple terms, interfaces allow us to define behavior that can be shared across different types without requiring them to be related through a common ancestor.

import "fmt"

type paymenter interface {
	pay(amount float32)
}
type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32){
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32){
	fmt.Println("makingpayment through razorpay", amount)
} // here we dotn explicitly say we are implementing paymenter interface


func main(){


}
