package main

import "fmt"

/*
мап - набор пар "ключ - значение"
например , меню ресторана
1. чай - 15р
2. пиво - 100р
позиция - значение т.е цена
*/


func main() {
	// пустая мапа
	// [ключ]значение
	menu := make(map[string]int)

	menu["Пиво"] = 100
	menu["Чай"] = 15
	menu["Сыр"] = 50

	// проверка, есть ли в меню
	// val, ok := menu["Сыр"]
	// fmt.Println(val, ok) // 50 true

	// delete(menu, "Сыр") // так можно удалять

	for item, price := range menu {
		fmt.Println(item, price)
	}

	// можно еще так
	// menu2 := map[string]float64 {
	// 	"Пиво": 100.00,
	// 	"Чай": 15.00,
	// 	"Орефки": 80.00,
	// }

	// menu["Пиво"] - значение 100.00
	// fmt.Println("Цена пива:", menu["Пиво"])
} 