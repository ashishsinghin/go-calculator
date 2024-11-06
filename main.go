package main

import (
	"fmt"
	"go-calculator/evaluator"
	"go-calculator/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: calc 'expression'")
		return
	}
	expression := os.Args[1]

	// Parse the expression to postfix notation
	postfix, err := parser.ParseExpression(expression)
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}

	// Evaluate the postfix expression
	result, err := evaluator.EvaluatePostfix(postfix)
	if err != nil {
		fmt.Println("Error evaluating expression:", err)
		return
	}

	fmt.Printf("Result: %v\n", result)
}
