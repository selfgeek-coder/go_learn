// указатели

package main

import "fmt"

type House struct {
    wallColor string
}

// без * буквально: я дам тебе дом, покрась в нем стены
// с *: я дам тебе АДРЕС дома, покрась в нем стены
func changeWallColor(color string, h *House) { // * говорит что функция примет адрес
    h.wallColor = color
}

func main() {
    h := House{wallColor: "White"}
    fmt.Println("Цвет стен до покраски: ", h.wallColor)

    changeWallColor("Green", &h) // &h: с & мы передаем адрес, без & мы передаем экземпляр дома (тоесть копию)
    fmt.Println("Цвет дома после покраски: ", h.wallColor)
}