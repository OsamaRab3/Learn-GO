package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to our calculator")
	for {
		choice := prog()
		if choice == "quit" {
			fmt.Println("Exiting the calculator. Goodbye!")
			break
		}

	
	switch choice {
	case "1":
		x, y := readTwoNumbers()
		fmt.Println("Sum:", sum(x, y))
	case "2":
		x, y := readTwoNumbers()
		fmt.Println("Subtraction:", sub(x, y))
	case "3":
		x, y := readTwoNumbers()
		fmt.Println("Multiplication:", mul(x, y))
	case "4":
		x, y := readTwoNumbers()
		fmt.Println("Division:", div(x, y))
	case "5":
		x, y := readTwoNumbers()
		fmt.Println("Modulus:", mod(float64(x), float64(y)))
	case "6":
		x, y := readTwoNumbers()
		fmt.Println("Power:", power(x, y))
	case "7":
		x := readOneNumber()
		fmt.Println("Exponential:", exp(x))
	case "8":
		x := readOneNumber()
		fmt.Println("Sine:", sin(x))
	case "9":
		x := readOneNumber()
		fmt.Println("Cosine:", cos(x))
	case "10":
		x := readOneNumber()
		fmt.Println("Tangent:", tan(x))
	case "11":
		x := readOneNumber()
		fmt.Println("Arcsine:", asin(x))
	case "12":
		x := readOneNumber()
		fmt.Println("Arccosine:", acos(x))
	case "13":
		x := readOneNumber()
		fmt.Println("Arctangent:", atan(x))
	case "14":
		x := readOneNumber()
		fmt.Println("Square Root:", sqrt(x))
	case "15":
		x := readOneNumber()
		fmt.Println("Cube Root:", cbrt(x))
	case "16":
		x := readOneNumber()
		fmt.Println("Logarithm (base e):", log(x))
	case "17":
		x := readOneNumber()
		fmt.Println("Logarithm (base 10):", log10(x))
	case "18":
		x := readOneNumber()
		fmt.Println("Absolute value:", abs(x))
	case "19":
		x := readOneNumber()
		fmt.Println("Ceiling value:", ceil(x))
	case "20":
		x := readOneNumber()
		fmt.Println("Floor value:", floor(x))
	case "21":
		n := readInt()
		fmt.Println("Factorial:", factorial(n))
	default:
		fmt.Println("Invalid choice")
	}
}
}

func prog() string  {
	fmt.Println("Select an operation:")
	fmt.Println("1 - Sum")
	fmt.Println("2 - Sub")
	fmt.Println("3 - Mul")
	fmt.Println("4 - Div")
	fmt.Println("5 - Mod")
	fmt.Println("6 - Power")
	fmt.Println("7 - Exp")
	fmt.Println("8 - Sin")
	fmt.Println("9 - Cos")
	fmt.Println("10 - Tan")
	fmt.Println("11 - Asin")
	fmt.Println("12 - Acos")
	fmt.Println("13 - Atan")
	fmt.Println("14 - Sqrt")
	fmt.Println("15 - Cbrt")
	fmt.Println("\n For Quit (Enter  quit ) ")

	var choice string 
	fmt.Scan(&choice)
	return choice
}

func readTwoNumbers() (int, int) {
	var x, y int
	fmt.Println("Enter two numbers:")
	fmt.Scan(&x, &y)
	return x, y
}

func readOneNumber() float64 {
	var x float64
	fmt.Println("Enter a number:")
	fmt.Scan(&x)
	return x
}


func readInt() int {
	var x int 
	fmt.Println("Enter a number:")
	fmt.Scan(&x)
	return x
}

func sum(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	if y == 0 {
		fmt.Println("Cannot divide by zero")
		return 0
	}
	return x / y
}

func sub(x, y int) int {
	return x - y
}

func power(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func exp(x float64) float64 {
	return math.Exp(x)
}

func sin(x float64) float64 {
	return math.Sin(x)
}

func cos(x float64) float64 {
	return math.Cos(x)
}

func tan(x float64) float64 {
	return math.Tan(x)
}

func asin(x float64) float64 {
	return math.Asin(x)
}

func acos(x float64) float64 {
	return math.Acos(x)
}

func atan(x float64) float64 {
	return math.Atan(x)
}

func mod(x, y float64) float64 {
	return math.Mod(x, y)
}

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func cbrt(x float64) float64 {
	return math.Cbrt(x)
}

func log(x float64) float64 {
    return math.Log(x)
}

func log10(x float64) float64 {
    return math.Log10(x)
}

func abs(x float64) float64 {
    return math.Abs(x)
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func ceil(x float64) float64 {
    return math.Ceil(x)
}

func floor(x float64) float64 {
    return math.Floor(x)
}

