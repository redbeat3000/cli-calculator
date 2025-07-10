package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run calculator.go [add|sub|mul|div] num1 num2")
		return
	}

	op := os.Args[1]
	a, err1 := strconv.ParseFloat(os.Args[2], 64)
	b, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Both arguments must be numbers.")
		return
	}

	switch op {
	case "add":
		fmt.Printf("Result: %.2f\n", a+b)
	case "sub":
		fmt.Printf("Result: %.2f\n", a-b)
	case "mul":
		fmt.Printf("Result: %.2f\n", a*b)
	case "div":
		if b == 0 {
			fmt.Println("Error: Cannot divide by zero.")
		} else {
			fmt.Printf("Result: %.2f\n", a/b)
		}
	default:
		fmt.Println("Unknown operation. Use: add, sub, mul, or div.")
	}
}
