package main

import (
	"fmt"
	"strings"
)

// run this file using :   go run main.go
func main() {
	fmt.Println("test")

	//variables
	a := "visal"
	var b string = "2nd String"
	const c int = 11222
	var f float32 = 12312.00000
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(int(f))

	var1, var2 := "va1", "---va2"

	fmt.Println(var1 + var2)

	callMe("callig call me Func")

	fmt.Println(addMe(5100, 5100))

	plus, minus := doSomeMath(5, 4)

	fmt.Println("plus: ", plus)
	fmt.Println("minus: ", minus)

	//String comparison

	fmt.Println(doCondition_if("visal"))
	fmt.Println(doCondition_if("visal kumar rao"))
}

func callMe(input string) {
	fmt.Println(input)
}

func addMe(input1 int, input2 int) int {
	return input1 + input2
}

// multiple return value from one function
func doSomeMath(intput1 int, intput2 int) (int, int) {
	return intput1 + intput2, intput1 - intput2
}

func doCondition_if(input string) bool {
	if input == "" {
		return false
	}
	return strings.EqualFold("visal", input)
}
