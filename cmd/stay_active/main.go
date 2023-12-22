package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func stayActive() {
	for {
		fmt.Println("Parriba..")
		x1, y1 := robotgo.Location()
		robotgo.DragSmooth(x1 - 100, y1 - 100)
		time.Sleep(30 * time.Second)

		fmt.Println("Pabajo..")
		x2, y2 := robotgo.Location()
		robotgo.DragSmooth(x2 + 100, y2 + 100)
		time.Sleep(30 * time.Second)
	}
}

func main() {
	fmt.Println("Trae pa' ca' el ratón...")
	stayActive()
}
