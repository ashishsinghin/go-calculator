package parser

import (
	"errors"
)

func shuntingYard(tokens []string) ([]string, error) {
	var outputQueue []string
	var operatorStack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	isOperator := func(token string) bool {
		_, exists := precedence[token]
		return exists
	}

	for _, token := range tokens {
		if token == "(" {
			operatorStack = append(operatorStack, token)
		} else if token == ")" {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
				outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			if len(operatorStack) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			operatorStack = operatorStack[:len(operatorStack)-1] // Pop the "("
		} else if isOperator(token) {
			for len(operatorStack) > 0 && precedence[operatorStack[len(operatorStack)-1]] >= precedence[token] {
				outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		} else {
			outputQueue = append(outputQueue, token)
		}
	}

	// Pop remaining operators onto the output queue
	for len(operatorStack) > 0 {
		if operatorStack[len(operatorStack)-1] == "(" {
			return nil, errors.New("mismatched parentheses")
		}
		outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return outputQueue, nil
}
