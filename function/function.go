package function

import (
	"errors"
	"fmt"
	"math"
)

func Hello(name string, age int64) (string, error) {
	if name == "Asel" && age < 18 {
		return "Hello little girl", nil
	}
	if name == "Dwayne" && age > 50 {
		return "Hi sugar daddy", nil
	}
	return "", errors.New("not found person, by name & age")
}

func Cast(value string) int {
	return 0
}

type Operator func(x float64) float64

// Map applies op to each element of a.
func Map(op Operator, a []float64) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = op(x)
	}
	return res
}

func FuncType() {
	op := math.Abs
	a := []float64{1, -2}
	b := Map(op, a)
	fmt.Println(b) // [1 2]

	c := Map(func(x float64) float64 { return 10 * x }, b)
	fmt.Println(c) // [10, 20]
}
func add(x int, y int) int {

	return x + y
}
func multiply(x int, y int) int {

	return x * y
}
func action(n1 int, n2 int, operation func(int, int) int) {

	result := operation(n1, n2)
	fmt.Println(result)
}
func FuncAsParam() {

	action(10, 25, add)    // 35
	action(5, 6, multiply) // 30
}

// func clousere(value string) func(string) int {
// 	return Cast(value)
// }
