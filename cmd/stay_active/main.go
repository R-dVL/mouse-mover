package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func StayActive() {
	for {
		robotgo.KeyToggle("capslock", "down")
	}
}

func main() {
	fmt.Println("Active...")
	StayActive()
}
