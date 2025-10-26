package main

import (
	"fmt"

	"github.com/sachinpoudel/learning_go/23.package/auth"
	"github.com/sachinpoudel/learning_go/23.package/user"
)

func main(){
	auth.LoginWithCredentials("sachin", "123")
	session:= auth.GetSession()
	fmt.Println(session)


userr:= user.User{
	Email: "sachin@gmail.com",
	Name: "sachin",
}
fmt.Println(userr)
}
