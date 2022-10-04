package datatypes

import (
	"fmt"
)

func Primitive() {
	var (
		str                = "Hello Alem!"
		doubleDeclaration1 = 12.1
	)

	numberDeclaration1 := 0 //init int 0 value
	var numberDeclaration2 uint64 = 12

	var arrayBytes []byte
	runes := []rune{} // 4byte
	charA := 'A'
	russianChar := 'Б'

	runes = append(runes, charA, russianChar)

	b := byte(1)
	b2 := byte(252)

	arrayBytes = append(arrayBytes, b, b2)

	var numberDeclaration3 complex128 //0
	var doubleDeclaration2 float64 = 153213123.21

	fmt.Printf("string : %s \n  float values : %f \n  numbers: %v \n  runes: %v \n bytes : %v \n", str, doubleDeclaration1, doubleDeclaration2, numberDeclaration1, numberDeclaration2, numberDeclaration3, runes, arrayBytes)

}

func Aggregate() {
	array := [7]rune{65, 'b', 67, 'c', 69, 'e', 70} // fixed
	//struct/object/class case
	type Person struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	p := Person{
		ID:   "123",
		Name: "Layer",
	}
	fmt.Printf(" Array: %v \n  struct: %v \n", array, p)
}

func Reference() {
	type Car struct {
		ID    string
		Model string
		Brand string
	}

	sliceStr := []string{}

	for _, letter := range "abcde" {
		sliceStr = append(sliceStr, string(letter))
	}
	//slice of struct
	sliceStruct := []Car{
		Car{
			ID:    "123",
			Model: "Nismo",
			Brand: "Toyota",
		},
		Car{
			ID:    "456",
			Model: "Gt-R",
			Brand: "Nissan",
		},
	}

	fmt.Printf(" Slice string: %v \n slice struct %v", sliceStr, sliceStruct)

	//map
	m := make(map[Car]interface{}) //init map
	car := Car{
		ID:    "123",
		Model: "Brabus",
		Brand: "Mercedes",
	}

	m[car] = "now value string type" //any value; then cast; key - struct - value string

	fmt.Printf("%v map value \n", m) // by pointer

	chanel := make(chan int, 3) //buf channel, 3 capacity;
	fmt.Printf("%v chan value ", <-chanel)

	//pointer case

	var n *int
	temp := 20
	n = &temp
	temp = 10

	fmt.Printf("%v pointer value \n %v temp value ", n, temp)

}
