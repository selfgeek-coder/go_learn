package main

import "fmt"

type Car struct {
	Brand string
	Year  int
}

// прибавляем +1 год к машине принимая адрес
func updateCarYearPointer(car *Car) {
	car.Year = car.Year + 1
}

func main() {
	car := Car{Brand: "Mersedes", Year: 2011}

	updateCarYearPointer(&car)

	fmt.Println(car)
}