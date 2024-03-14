package main

import (
	"fmt"
)

func input() {
	var num1, num2, operator string

	fmt.Println("Введите цифры от 0 до 10 и выполняемое действие '+' '-' '%' '*' через пробел:")
	count, err := fmt.Scanf("%s %s %s", &num1, &operator, &num2)

	if err != nil {
		fmt.Printf("Ошибка при чтении ввода, выполните условие ввода: %v\n", err)
	} else if count != 3 {
		fmt.Printf("Ошибка при чтении ввода, выполните условие ввода")
	} else {
		fmt.Printf("Вы ввели %s %s %s\n", num1, operator, num2)
	}

}

func main() {
	input()
}
