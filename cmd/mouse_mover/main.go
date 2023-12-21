package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func mouseMover() {
	for {
		fmt.Println("Parriba..")
		robotgo.MoveRelative(-100, -100)
		time.Sleep(30 * time.Second)

		fmt.Println("Pabajo..")
		robotgo.MoveRelative(100, 100)
		time.Sleep(30 * time.Second)
	}
}

func main() {
	fmt.Println("Trae pa' ca' el ratón...")
	mouseMover()
}
