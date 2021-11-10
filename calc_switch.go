package main

import (
	"fmt"
	"math"
)

func main() {

	var operator string
	var operand1 float64
	var operand2 float64
	var result float64

	fmt.Println("Добро пожаловать в программу Калькулятор.Вам необходимо ввести поочередно два числа. После ввода " +
		"задуманного числа, нажмите Enter.\n" +
		"Затем,выберите математическую операцию и введите ее в поле. " +
		"Для получения вычислений нажмите Enter." +
		"Список доступных математических операций:\n" +
		"Сложение - '+'\n" +
		"Вычитание - '-'\n" +
		"Умножение - '*'\n" +
		"Деление '-' /\n" +
		"Квадратный корень - 'sqrt'\n" +
		"Возведение в степень '^'\n" +
		"Введите первое число:")
	fmt.Scanln(&operand1)
	fmt.Println("Введите математическую операцию:")
	fmt.Scanln(&operator)
	if operator=="sqrt" {
		result = math.Sqrt(operand1)
		fmt.Println(result)
		return
	}
	fmt.Println("Введите второе число:")
	fmt.Scanln(&operand2)


	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 + operand2
	case "*":
		result = operand1 + operand2
	case "/":
		result = operand1 + operand2
	case "^":
		result =math.Pow (operand1,operand2)


	}
	fmt.Println(result)
}
