package evaluator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func EvaluatePostfix(postfix []string) (float64, error) {
	var stack Stack

	for _, token := range postfix {
		// Check if the token is a number
		if value, err := strconv.ParseFloat(token, 64); err == nil {
			stack.Push(value)
		} else if isFunction(token) {
			// It's a function like sin, cos, or tan
			arg, ok := stack.Pop()
			if !ok {
				return 0, errors.New("invalid expression")
			}
			var result float64
			switch strings.ToLower(token) {
			case "sin":
				result = math.Sin(arg)
			case "cos":
				result = math.Cos(arg)
			case "tan":
				result = math.Tan(arg)
			default:
				return 0, errors.New("unknown function")
			}
			stack.Push(result)
		} else {
			// It's an operator
			b, ok := stack.Pop()
			if !ok {
				return 0, errors.New("invalid expression")
			}
			a, ok := stack.Pop()
			if !ok {
				return 0, errors.New("invalid expression")
			}

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				result = a / b
			default:
				return 0, errors.New("unknown operator")
			}
			stack.Push(result)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}
	result, _ := stack.Pop() // Use both values, but ignore the boolean in this context
	return result, nil
}

// Helper function to check if the token is a mathematical function
func isFunction(token string) bool {
	functions := map[string]bool{
		"sin": true,
		"cos": true,
		"tan": true,
	}
	return functions[strings.ToLower(token)]
}
