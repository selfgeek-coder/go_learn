package main

import "fmt"

type Movable interface {
	Move() string
}

type Car struct {
	Brand string
}

type Person struct {
	Name string
}

func (c Car) Move() string {
	return "Машина " + c.Brand + " едет со скоростью X км/ч"
}

func (p Person) Move() string {
	return "Человек " + p.Name + " идёт со скоростью Y км/ч"
}

func printMovement(m Movable) string {
	return m.Move()
}

func main() {
	car := Car{Brand: "Mercedes"}
	fmt.Println(printMovement(car))

	person := Person{Name: "Oleg"}
	fmt.Println(printMovement(person))
}