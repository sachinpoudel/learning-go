package types


type Student struct{
	Id int 
	Name string `validate-required`

	Email int `validate-required`
	Age int `validate-required`
}