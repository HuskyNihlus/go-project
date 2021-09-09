package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readValue() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	return s
}

func readOperation() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	return s
}

func convert(s string) float64 {
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	value1 := readValue()
	v1 := convert(value1)
	value2 := readValue()
	v2 := convert(value2)
	operation := readOperation()
	switch operation {
	case "+":
		fmt.Printf("%.4f", v1 + v2)
	case "-":
		fmt.Printf("%.4f", v1 - v2)
	case "*":
		fmt.Printf("%.4f", v1 * v2)
	case "/":
		fmt.Printf("%.4f", v1 / v2)
	default:
		fmt.Println("Неизвестная операция")
	}
}