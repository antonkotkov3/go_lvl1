package main

import (
	"fmt"
)

func main() {
	var n int

	slice := []int{}
	fmt.Println("Введите любое количество целых чисел, но не более 1000 шт. \n" +
		" Вводить нужно не по порядку! Программа сама отсортирует введеные числа \n" +
		"Чтобы выйти из режима ввода чисел, введите '9999")
	fmt.Scan(&n)
	// Вводим числа, чтобы заполнить ими динамический массив
	for i := 0; i < 1001; i++ {
		if n != 9999 {
			fmt.Scan(&n)
			slice = append(slice, n) // тут каждая п это очередное новое число массива

		} else {
			var result = slice[0 : len(slice)-1] //удаляю последний элемент массива 9999

			fmt.Println(result)
			break

		}

	}

	fmt.Println("Отлично, выше мы видим введенные числа в ввиде массива, ниже - сортировка по:\n" +

		"От меньшего к большему\n" +
		"Длина массива\n")

	//сортирую от меньшего к большему
	for i := 0; i < len(slice); i++ {
		x := slice[i]
		j := i
		for ; j >= 1 && slice[j-1] > x; j-- {
			slice[j] = slice[j-1]
		}
		slice[j] = x

	}
	sort1 := slice[0 : len(slice)-1]
	fmt.Println(sort1)
	sort2 := len(slice) - 1
	fmt.Println("Длина массива: ", sort2)
}
