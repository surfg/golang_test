package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func input() (string, string, string) {
	var inputString string

	fmt.Println("Введите цифры от 0 до 10 и выполняемое действие '+' '-' '%' '*' через пробел:")
	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Ошибка: введите ровно два числа и один оператор (+, -, *, /), через пробел %v\n", err)
		return "", "", ""
	}

	inputString = strings.TrimSpace(inputString)

	parts := strings.Fields(inputString)
	if len(parts) != 3 {
		fmt.Println("Ошибка: введите ровно два числа и один оператор (+, -, *, /), через пробел")
		return "", "", ""
	}
	return parts[0], parts[1], parts[2]
}

func checkValues(num1Str, operator, num2Str string) (int, int, string, error) {
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Недопустимый оператор или отсутсвие.Введите ровно два числа и один оператор (+, -, *, /)")
		return 0, 0, "", fmt.Errorf("недопустимый оператор: %s", operator)
	}

	var numType1, numType2 string

	checkNumbers := func(numStr string) (int, string, error) {
		var numType string

		num, err := strconv.Atoi(numStr)
		if err == nil {
			if num < 11 {
				numType = "arabic"
			} else {
				return 0, "", fmt.Errorf("числа должны быть меньше 11")
			}
		} else {
			if _, ok := romanNumerals[numStr]; ok {
				num = romanNumerals[numStr]
				numType = "roman"
			} else {
				return 0, "", fmt.Errorf("числа недопустимы: %s", numStr)
			}
		}
		return num, numType, nil
	}

	num1, numType1, err1 := checkNumbers(num1Str)
	if err1 != nil {
		return 0, 0, "", err1
	}

	num2, numType2, err2 := checkNumbers(num2Str)
	if err2 != nil {
		return 0, 0, "", err2
	}

	if numType1 != numType2 {
		return 0, 0, "", fmt.Errorf("одно число арабское, другое римское: %s и %s", num1Str, num2Str)
	}

	return num1, num2, numType1, nil
}

func arabComputations(num1 int, num2 int, operator string) (string, error) {
	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", fmt.Errorf("деление на ноль")
		}
		result = num1 / num2
	default:
		return "", fmt.Errorf("недопустимый оператор: %s", operator)
	}
	return strconv.Itoa(result), nil

}

func RomanComputation(num1 int, num2 int, operator string) (int, error) {
	var result int

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
		if result < 0 {
			return 0, fmt.Errorf("результат отрицательный, такого не может быть в римской системе")
		}
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		if result < 1 {
			return 0, fmt.Errorf("результат меньше единицы, такого не может быть в римской системе")
		}
	default:
		return 0, fmt.Errorf("недопустимый оператор: %s", operator)
	}

	return result, nil
}

func arabToRoman(number int) string {
	var roman strings.Builder

	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	for _, numeral := range romanNumerals {
		for number >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			number -= numeral.Value
		}
	}

	return roman.String()
}

func handleComputation(num1Str, operator, num2Str string) {
	num1, num2, numType, err := checkValues(num1Str, operator, num2Str)
	if err != nil {
		fmt.Printf("Произошла ошибка на уровне ввода: %v\n", err)
		return
	}

	switch numType {
	case "arabic":
		result, err := arabComputations(num1, num2, operator)
		outputResult(num1, num2, numType, result, err)
	case "roman":
		result, err := RomanComputation(num1, num2, operator)
		resultInRomanSystem := arabToRoman(result)
		outputResult(num1, num2, numType, resultInRomanSystem, err)
	}
}

func outputResult(num1, num2 int, numType string, result string, err error) {
	if err != nil {
		fmt.Printf("Ошибка вычисления: %v\n", err)
	} else {
		fmt.Printf("%s\n", result)
	}
}

func main() {
	num1Str, operator, num2Str := input()
	handleComputation(num1Str, operator, num2Str)
}
