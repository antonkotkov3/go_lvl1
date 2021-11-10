//Простейший калькулятор без валидации данных
package main

import (
	"fmt"
	"math"
)
func main() {
	var number1 float64
	var number2 float64
	var operation string
	var result float64
fmt.Println("Добро пожаловать в программу Калькулятор.Вам необходимо ввести поочередно два числа. После ввода " +
	"задуманного числа нажмите Enter.\n" +
	"Затем,выберите математическую операцию и введите ее в поле. " +
	"Для получения вычислений нажмите Enter." +
	"Список доступных математических операций:\n" +
	"Сложение - sum\n" +
	"Вычитание - min\n" +
	"Умножение - umn\n" +
	"Деление - del\n" +
	"Квадратный корень - sqrt" )
	fmt.Println("Введите первое число")
fmt.Scan(&number1)
fmt.Println("Введите математическую операцию")
fmt.Scan(&operation)
	if operation == "sqrt" {
		result=math.Sqrt(number1)
		fmt.Println("Результат", result)
		return
	}
fmt.Println("Введите второе число")
fmt.Scan(&number2)

if operation == "sum" {
	result = number2+number1
	fmt.Println("Результат ",result)
} else if operation=="min" {
	result = number1-number2
	fmt.Println("Результат ",result)
	} else if operation=="umn" {
	result = number2*number1
	fmt.Println("Результат ",result)
	} else if operation=="del" {
	result = number1/number2
	fmt.Println("Результат ",result)
	}else {
		fmt.Println("Что то пошло не так, попробуйте снова.")
		return
}
}
