/*
Требования:

Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b.
Данные передаются в одну строку (смотри пример ниже).
Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.

Калькулятор умеет работать как с арабскими (1,2,3,4,5…), так и с римскими (I,II,III,IV,V…) числами.
Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.
Калькулятор умеет работать только с целыми числами.
Калькулятор умеет работать только с арабскими или римскими цифрами одновременно,
при вводе пользователем строки вроде 3 + II калькулятор должен указать на ошибку и прекратить работу.

При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.
При вводе пользователем не подходящих чисел приложение выводит ошибку в терминал и завершает работу.
При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выводит ошибку и завершает работу.
Результатом операции деления является целое число, остаток отбрасывается.
Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль.
Результатом работы калькулятора с римскими числами могут быть только положительные числа,
если результат работы меньше единицы, программа должна указать на исключение.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RazMap map[string]string
type AvailableRomanMap map[string]int

var firstRaz RazMap = RazMap{
	"0": "0",
	"1": "I",
	"2": "II",
	"3": "III",
	"4": "IV",
	"5": "V",
	"6": "VI",
	"7": "VII",
	"8": "VIII",
	"9": "IX",
}

var secondRaz RazMap = RazMap{
	"1":  "X",
	"2":  "XX",
	"3":  "XXX",
	"4":  "XL",
	"5":  "L",
	"6":  "LX",
	"7":  "LXX",
	"8":  "LXXX",
	"9":  "XC",
	"10": "C",
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var romanMap AvailableRomanMap
	romanMap = make(AvailableRomanMap, 10)
	if romanMap != nil {
		romanMap["I"] = 1
		romanMap["II"] = 2
		romanMap["III"] = 3
		romanMap["IV"] = 4
		romanMap["V"] = 5
		romanMap["VI"] = 6
		romanMap["VII"] = 7
		romanMap["VIII"] = 8
		romanMap["IX"] = 9
		romanMap["X"] = 10
	}

	for {
	repeat:
		fmt.Print("Введите выражение с арабскими (1,2,3,4,5…) или с римскими (I,II,III,IV,V…) числами от 1 до 10 \n a + b, a - b, a * b, a / b :  ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		var expression []string = strings.Split(text, " ")
		convertedToRomanResult := ""
		errText := ""

		if len(expression) != 3 {
			expression = strings.Split(text, "")
			if len(expression) < 3 {
				fmt.Println("Выражение не валидно\n")
				goto repeat
			}
			if len(expression) != 3 {
				fmt.Println("Строка должна содержать пробелы\n")
				goto repeat
			}
		}

		// Вычисление для римских чисел
		errText = calculateRomanExpression(expression, &convertedToRomanResult, romanMap)
		if len(errText) > 0 {
			fmt.Println(errText)
			errText = ""
			goto repeat
		}
		if len(convertedToRomanResult) > 0 {
			goto repeat
		}

		// Вычисление для арабских чисел
		errText = calculateArabicExpression(expression)
		if len(errText) > 0 {
			fmt.Println(errText)
			errText = ""
		}
	}
}

func convertResult(result int) string {
	splitedResult := strings.Split(strconv.Itoa(result), "")
	switch len(splitedResult) {
	case 3:
		return "C"
	case 2:
		if firstRaz[splitedResult[1]] != "0" {
			return secondRaz[splitedResult[0]] + firstRaz[splitedResult[1]]
		}
		return secondRaz[splitedResult[0]]
	case 1:
		return firstRaz[splitedResult[0]]
	default:
		return ""
	}
}

func calculateRomanExpression(expression []string, convertedToRomanResult *string, romanMap AvailableRomanMap) string {
	operation := strings.TrimSpace(expression[1])
	valRim1, isRom1 := romanMap[strings.ToUpper(expression[0])]
	valRim2, isRom2 := romanMap[strings.ToUpper(expression[2])]
	if (!isRom1 && isRom2) || (isRom1 && !isRom2) {
		return "Введите выражение только с арабскими (1,2,3,4,5…) или только с римскими (I,II,III,IV,V…) числами\n"
	}

	if isRom1 && isRom2 {
		if valRim1 < 1 || valRim1 > 10 {
			return "Число " + strconv.Itoa(valRim2) + " не валидно\n"
		}
		if valRim2 < 1 || valRim2 > 10 {
			return "Число " + strconv.Itoa(valRim2) + " не валидно\n"
		}
		result := 0

		switch operation {
		case "+":
			result = valRim1 + valRim2
		case "-":
			result = valRim1 - valRim2
			if result < 0 {
				return "В римской системе нет отрицательных чисел, выражение не валидно\n"
			}
		case "*":
			result = valRim1 * valRim2
		case "/":
			result = valRim1 / valRim2
		default:
			return "Операция " + operation + " не валидна\n"
		}

		*convertedToRomanResult = convertResult(result)
		if len(*convertedToRomanResult) > 0 {
			fmt.Println("\nРезультат: " + strings.ToUpper(expression[0]) + " " + operation + " " + strings.ToUpper(expression[2]) + " = " + *convertedToRomanResult + "\n")
		}
	}
	return ""
}

func calculateArabicExpression(expression []string) string {
	operation := strings.TrimSpace(expression[1])
	num1, err1 := strconv.Atoi(expression[0])
	num2, err2 := strconv.Atoi(expression[2])

	if err1 != nil || err2 != nil || len(expression) != 3 {
		return "Выражение: \"" + expression[0] + " " + operation + " " + expression[2] + "\" - не валидно \n"
	}

	if num1 < 1 || num1 > 10 {
		return "Число " + strconv.Itoa(num1) + " не валидно\n"
	}

	if num2 < 1 || num2 > 10 {
		return "Число " + strconv.Itoa(num2) + " не валидно\n"
	}

	if err1 != nil || err2 != nil {
		return "Число не валидно\n"
	}

	result := 0
	isResult := true

	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		isResult = false
		return "Операция " + operation + " не валидна\n"
	}

	if isResult {
		fmt.Println("\nРезультат: " + strconv.Itoa(num1) + " " + operation + " " + strconv.Itoa(num2) + " = " + strconv.Itoa(result) + "\n")
	}
	return ""
}
