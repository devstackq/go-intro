package main

import (
	"awesome-project/datatypes"

	"fmt"
)

func init() {
	fmt.Println("first run")
}

func main() {
	//main thread
	fmt.Println("second run")
	//varibales
	datatypes.Primitive()
	// datatypes.Aggregate()
	// datatypes.Reference()

	// fmt.Println(function.Hello("Asel", 6))
	// fmt.Println(function.Hello("Polina", 76))
	// fmt.Println(function.Hello("Dwayne", 69))

	// function.FuncType()
	// function.FuncAsParam()

	// h1 := methods.BornHuman("Asya", "boy", "nigga", "Sudan", "sudanec")
	// newID := uuid.New()
	// h1.SetDocumentID(newID)

	// log.Println("get incapsulate data Human class", h1.GetDocumentID())

	// company := methods.NewCompany("Kwaaka", "Smart Point", "IT")

	// builder := methods.NewBuilder(h1, company)
	// //builder.SetCompany()

	// builder.Work()

	//OOP example
	//Mongo example
	//libs case

}
