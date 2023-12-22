package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func StayActive() {
	for {
		fmt.Println("Still here..")
		robotgo.KeyToggle("capslock", "down")
		robotgo.MilliSleep(1000)
		robotgo.KeyToggle("capslock", "up")
		robotgo.KeyToggle("capslock", "down")
		robotgo.MilliSleep(60000)
	}
}

func main() {
	fmt.Println("Hey, I'm here..")
	StayActive()
}
