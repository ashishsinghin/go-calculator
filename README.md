# Go Calculator

A simple calculator built in Go that can handle basic arithmetic and more complex mathematical expressions, including functions like `sin`, `cos`, and `tan`.

## Features

- **Basic Arithmetic**: `+`, `-`, `*`, `/`
- **Order of Operations**: Handles parentheses and operator precedence
- **Mathematical Functions**: Supports `sin`, `cos`, and `tan` functions
- **Error Handling**: Detects mismatched parentheses, division by zero, and other invalid expressions

## Getting Started

### Prerequisites

- [Go](https://golang.org/) installed on your system

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/siashish/go-calculator.git
   cd go-calculator
   ```

2. Build the project:
   ```bash
   go build
   ```

### Usage

You can run the calculator from the command line:

```bash
go run main.go "2 * (3 + 4) / 2"
```

Output:

```
Result: 7
```

### Example Expressions

- Basic arithmetic: `"1 + 2 * 3"`
- Parentheses for precedence: `"(1 + 2) * 3"`
- Trigonometric functions: `"sin(0)"`, `"cos(0)"`, `"tan(0)"`

### Running Tests

To run the unit tests, use:

```bash
go test ./...
```

---

## Project Structure

```
calculator/
│
├── main.go
├── parser/
│   ├── parser.go
│   └── shunting_yard.go
└── evaluator/
    ├── evaluator.go
    └── stack.go
```

- **main.go**: Entry point for the application
- **parser/**: Handles parsing the expression and converting to postfix notation
- **evaluator/**: Evaluates the postfix expression using a stack

---

## Extending the Calculator

You can extend the calculator by adding more mathematical functions like `sqrt`, `log`, or constants like `π` and `e`. Here's how to add a new function:

1. Update `isFunction()` in `evaluator/evaluator.go`.
2. Add the function's logic in the `switch` statement.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request for any features or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/siashish/go-calculator/blob/main/LICENSE) file for details.

## Acknowledgments

- Inspired by the classic [Shunting Yard algorithm](https://en.wikipedia.org/wiki/Shunting_yard_algorithm#The_algorithm_in_detail) by Edsger Dijkstra
- Go's `math` package for trigonometric functions