package main

import (
	"awesome-project/datatypes"
	"awesome-project/function"
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
	datatypes.Aggregate()
	datatypes.Reference()

	fmt.Println(function.Hello("Asel", 6))
	fmt.Println(function.Hello("Polina", 76))
	fmt.Println(function.Hello("Dwayne", 69))

	function.FuncType()
	function.FuncAsParam()
	//OOP example
	//Mongo example
	//libs case

}
