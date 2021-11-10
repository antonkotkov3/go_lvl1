package main

import (
	"fmt"

)

func fibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}
func main() {
	var number uint


	fmt.Printf("Введите число для получения результата. Мoжно вводить поочередно неограниченное количество чисел.\n" +
		"Чтобы выйти из программы - введите 000\n")
	_, _ = fmt.Scan(&number)
	i := 0
	mapa := map[uint]uint{}
	for i = 0; ; i++ {
		if number != 000 {
			fmt.Println("Результат: ", fibbonachi(number))
			mapa[number] = fibbonachi(number)

			fmt.Println("Следующее число : ")
			_, _ = fmt.Scan(&number)
		} else if number == 000 {
			fmt.Println("Ниже Вы можете видеть сохраненные результаты Ваших вычислений: ")
			for k, v := range mapa {
				fmt.Println("Число:", k, "Результат:", v)
			}


			fmt.Println("Конец программы!")
			break
		}

		}

	}
