package main

import "fmt"

// START OMIT
type Appliance interface { // 介面型別（interface types）：只描述這個值能夠做什麼
	TurnOn()
}
type Fan string // 具象型別（concrete types）：指明它的值可以做什麼，也代表它們是什麼

func (f Fan) TurnOn() {
	fmt.Println(f, "- Spinning")
}

type CoffeePot string // 具象型別（concrete types）

func (c CoffeePot) TurnOn() {
	fmt.Println(c, "- Powering up")
}
func (c CoffeePot) Brew() {
	fmt.Println(c, "- Heating Up")
}
func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device = CoffeePot("LuxBrew")
	device.TurnOn()
}

// END OMIT
