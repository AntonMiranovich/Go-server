package main

import "fmt"

func main() {
	// 1. Проверка возраста
	age := 25
	if age < 18 {
		fmt.Println("Вы несовершеннолетний")
	} else if age >= 18 && age <= 60 {
		fmt.Println("Вы взрослый")
	} else {
		fmt.Println("Вы пожилой человек")
	}


		// 2. Оценка успеваемости
		grade := 85
		switch {
		case grade >= 90 && grade <= 100:
			fmt.Println("Отлично")
		case grade >= 75 && grade <= 89:
			fmt.Println("Хорошо")
		case grade >= 50 && grade <= 74:
			fmt.Println("Удовлетворительно")
		default:
			fmt.Println("Неудовлетворительно")
		}
	
		// 3. День недели по номеру
		day := 3
		switch day {
		case 1:
			fmt.Println("Понедельник")
		case 2:
			fmt.Println("Вторник")
		case 3:
			fmt.Println("Среда")
		case 4:
			fmt.Println("Четверг")
		case 5:
			fmt.Println("Пятница")
		case 6:
			fmt.Println("Суббота")
		case 7:
			fmt.Println("Воскресенье")
		default:
			fmt.Println("Неверный номер дня")
		}
	
		// 4. Четное/нечетное
		num := 7
		if num%2 == 0 {
			fmt.Println("Четное")
		} else {
			fmt.Println("Нечетное")
		}

}
