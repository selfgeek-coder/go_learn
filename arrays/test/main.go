package main

import "fmt"

type Laptop struct {
	Brand string
	Price int
}

func main() {
	laptops := []Laptop{
		{Brand: "Asus", Price: 1000},
		{Brand: "Acer", Price: 700},
	}

	for index, laptop := range laptops {
		fmt.Printf("%d. %s, %d\n", index, laptop.Brand, laptop.Price)
	}
}