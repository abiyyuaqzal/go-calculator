package main

import "fmt"

func main() {

	fmt.Println("GOLANG CALCULATOR")

	var a, b float64
	var operator string

	fmt.Println("Enter your first number : ")
	fmt.Scanln(&a)

	fmt.Println("Enter your second number : ")
	fmt.Scanln(&b)

	fmt.Println("Enter the operator (+, -, *, /) : ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%.3f %s %.3f = %.3f", a, operator, b, a+b)
	case "-":
		fmt.Printf("%.3f %s %.3f = %.3f", a, operator, b, a-b)
	case "*":
		fmt.Printf("%.3f %s %.3f = %.3f", a, operator, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Divided by zero situation")
		} else {
			fmt.Printf("%.3f %s %.3f = %.3f", a, operator, b, a/b)
		}
	default:
		fmt.Println("Invalid Operator")
	}

}
