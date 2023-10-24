package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func scanFloat() float64 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	digit, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return digit
}

func main() {

	fmt.Print("Введите первое число: ")
	digit1 := scanFloat()

	fmt.Print("Введите один из ниже перечисленных знаков\n+ - * / : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)

	}
	sign := strings.TrimSpace(input)
	if sign != "+" && sign != "-" && sign != "*" && sign != "/" {
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		return
	}

	fmt.Print("Введите второе число: ")
	digit2 := scanFloat()

	if sign == "+" {
		fmt.Println("a + b = ", digit1+digit2)
	} else if sign == "-" {
		fmt.Println("a - b = ", digit1-digit2)
	} else if sign == "*" {
		fmt.Println("a * b = ", digit1*digit2)
	} else {
		if digit2 == 0 {
			fmt.Println("деление на ноль невозможно.")
		} else {
			fmt.Println("a / b = ", digit1/digit2)
		}
	}
}
