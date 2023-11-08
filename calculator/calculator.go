package calculator

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number> <operator> <number>")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator := os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid numbers provided.")
		return
	}

	result, err := calculate(num1, operator, num2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}

func calculate(num1 float64, operator string, num2 float64) (float64, error) {
	var result float64
	var err error

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			err = fmt.Errorf("Division by zero is not allowed")
		} else {
			result = num1 / num2
		}
	default:
		err = fmt.Errorf("Invalid operator: %s", operator)
	}

	return result, err
}
