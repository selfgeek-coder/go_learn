package main

import "fmt"

// интерфейс это чертеж который говорит, что если это все есть значит это является это N объектом

// например
// если это крякает как утка, плавает как утка то это утка


type Flyer interface {
	Fly() string
}

func describeFlight(object Flyer) string {
	return object.Fly()
}

type Bird struct {
	species string
}

type Airplane struct {
	model string
}

func (a Airplane) Fly() string {
	return  a.model + " взлетел"
}

func (b Bird) Fly() string {
	return b.species + " взлетела"
}


func main() {
	pigeon := Bird{species: "pigeon"}
	fmt.Println(describeFlight(pigeon))

	boeing := Airplane{model: "212"}
	fmt.Println(describeFlight(boeing))
}