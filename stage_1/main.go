package main

import (
	"fmt"
)

func calculate(a int, b int, operation string) (int, string) {
	var result int
	var err string
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	case "%":
		result = a % b
	case "**":
		result = 1
		if b == 0 {
			result = 0
			break
		} else if b == 1 {
			result = a
			break
		}
		for range b {
			result = result * a
		}
	default:
		err = "Неизвестная операция"
	}
	return result, err
}

func main() {
	var i, j int
	var op string
	for {
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&i)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&j)
		fmt.Print("Введите операцию: ")
		fmt.Scanln(&op)
		r, err := calculate(i, j, op)
		if err != "" {
			fmt.Println(err, op)
			continue
		}
		fmt.Printf("Результат: %d %s %d = %d\n", i, op, j, r)
	}
}
