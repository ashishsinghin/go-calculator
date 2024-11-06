package parser

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ParseExpression parses an infix expression and converts it to a postfix expression using the shunting yard algorithm
func ParseExpression(expression string) ([]string, error) {
	var outputQueue []string
	var operatorStack []string
	tokens := tokenize(expression)

	for _, token := range tokens {
		if isNumber(token) {
			outputQueue = append(outputQueue, token)
		} else if isFunction(token) {
			operatorStack = append(operatorStack, token)
		} else if token == "(" {
			operatorStack = append(operatorStack, token)
		} else if token == ")" {
			// Pop from the stack to the output queue until we find a left parenthesis
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
				outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			if len(operatorStack) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			// Pop the left parenthesis from the stack
			operatorStack = operatorStack[:len(operatorStack)-1]

			// If the token at the top of the stack is a function, pop it to the output queue
			if len(operatorStack) > 0 && isFunction(operatorStack[len(operatorStack)-1]) {
				outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
		} else if isOperator(token) {
			for len(operatorStack) > 0 && precedence(operatorStack[len(operatorStack)-1]) >= precedence(token) {
				outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		} else {
			return nil, errors.New("invalid token: " + token)
		}
	}

	// Pop the remaining operators from the stack to the output queue
	for len(operatorStack) > 0 {
		if operatorStack[len(operatorStack)-1] == "(" || operatorStack[len(operatorStack)-1] == ")" {
			return nil, errors.New("mismatched parentheses")
		}
		outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return outputQueue, nil
}

// tokenize splits the expression into tokens
func tokenize(expression string) []string {
	var tokens []string
	var currentToken strings.Builder

	for _, char := range expression {
		switch {
		case unicode.IsDigit(char) || char == '.':
			currentToken.WriteRune(char)
		case unicode.IsLetter(char):
			currentToken.WriteRune(char)
		case char == ' ':
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
		case isOperatorChar(char) || char == '(' || char == ')':
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		default:
			// Ignore any unrecognized characters
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}

// isNumber checks if a token is a valid number
func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

// isFunction checks if a token is a recognized mathematical function
func isFunction(token string) bool {
	functions := map[string]bool{
		"sin": true,
		"cos": true,
		"tan": true,
	}
	return functions[strings.ToLower(token)]
}

// isOperator checks if a token is an operator
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

// isOperatorChar checks if a rune is an operator character
func isOperatorChar(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

// precedence returns the precedence of an operator
func precedence(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}
