package main

import "fmt"

func main() {
	menu := make(map[string]float64)
	menu["Макароны"] = 125.00

	_, ok := menu["Салат"]

	if !ok {
		menu["Салат"] = 300.00
	}

	delete(menu, "Макароны")

	for item, value := range menu {
		fmt.Println(item, value)
	}
}