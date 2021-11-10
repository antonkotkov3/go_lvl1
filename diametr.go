//Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.
package main

import (
	"fmt"
	"math"
)
func main() {
var sCircle float64 //площадь круга, вводимая вручную
var l float64 //длина окружности
var d float64 // диаметр окружности
	fmt.Println("Введите площадь круга:")
	fmt.Scan(&sCircle)
l=math.Sqrt(sCircle*4*math.Pi)
fmt.Println("Длина окружности равна: ", l )
d=l/math.Pi
fmt.Println("Диаметр равен :",d)

}