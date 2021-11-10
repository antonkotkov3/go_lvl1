//С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе
package main
import "fmt"
func main() {
	var number int
	var sot int//количество сотен
	var des int // количество десятков
	var ed int //количество единиц
	fmt.Println("Введите трёхзначное число: ")
	fmt.Scan (&number)
	sot=number/100
	des=number/10
	ed=number
	fmt.Println("Количество сотен:",sot )
	fmt.Println("Количество десятков:",des )
	fmt.Println("Количество единиц:",ed )
}