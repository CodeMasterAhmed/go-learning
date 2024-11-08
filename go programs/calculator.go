package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Go Calculator!")

	for {
		var a, b int
		var op string

		fmt.Print("Enter the first number: ")
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Println("Invalid input for the first number.")
			continue
		}

		fmt.Print("Enter an operation (+, -, *, /, %): ")
		fmt.Scan(&op)

		fmt.Print("Enter the second number: ")
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Println("Invalid input for the second number.")
			continue
		}

		result, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("The result of %d %s %d is: %d\n", a, op, b, result)
		}

		var choice string
		fmt.Print("Would you like to perform another calculation? (y/n): ")
		fmt.Scan(&choice)
		if strings.ToLower(choice) != "y" {
			fmt.Println("Thank you for using the calculator. Goodbye!")
			break
		}
	}
}

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero is not allowed")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, errors.New("modulus by zero is not allowed")
		}
		return a % b, nil
	default:
		return 0, errors.New("unsupported operation")
	}
}
