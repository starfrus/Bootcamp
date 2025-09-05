package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var s1, oper, s2 string
	var num1, num2 float64
	error_name := "Invalid input"

	err := errors.New(error_name)
	fmt.Println("Input left operand")
	for {
		fmt.Scanf("%s", &s1)
		num1, err = strconv.ParseFloat(s1, 64)
		if err != nil {
			fmt.Println(error_name)
		} else {
			break
		}
	}
	fmt.Println("Input operation [+] [-] [*] [/]")
	for {
		fmt.Scanf("%s", &oper)
		if oper == "+" || oper == "-" || oper == "*" || oper == "/" {
			break
		} else {
			fmt.Println(error_name)
		}
	}
	err = errors.New(error_name)
	fmt.Println("Input right operand")
	for {
		fmt.Scanf("%s", &s2)
		num2, err = strconv.ParseFloat(s2, 64)
		if err != nil {
			fmt.Println(error_name)
		} else {
			break
		}
	}
	switch oper {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Division zero")
			break
		}
		fmt.Printf("%.3f\n", num1/num2)
	}
}
