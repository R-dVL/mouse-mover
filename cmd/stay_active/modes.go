package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func Work() {
	ClearTerminal()

	for {
		fmt.Println("Still here..")
		robotgo.KeyTap("capslock")
		robotgo.KeyTap("capslock")
		robotgo.MilliSleep(60000)
	}
}

func Gaming() {
	ClearTerminal()

	for {
		// Generate seed and random number
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		action := rand.Intn(4)

		switch action {
		case 0:
			fmt.Println("Move forward")
			robotgo.KeyToggle("w", "down")
			robotgo.MilliSleep(1000)
			robotgo.KeyToggle("w", "up")
			robotgo.MilliSleep(60000)
		case 1:
			fmt.Println("Move backwards")
			robotgo.KeyToggle("s", "down")
			robotgo.MilliSleep(1000)
			robotgo.KeyToggle("s", "up")
			robotgo.MilliSleep(60000)
		case 2:
			fmt.Println("Action 1")
			robotgo.KeyTap("1")
			robotgo.MilliSleep(60000)
		case 3:
			fmt.Println("Action 2")
			robotgo.KeyTap("2")
			robotgo.MilliSleep(60000)
		default:
			fmt.Println("Not an option.")
		}
	}
}

func MouseMover() {
	ClearTerminal()

	for {
		fmt.Println("Up..")
		robotgo.MoveSmoothRelative(-100, -100, 1.0, 30.0)
		robotgo.MilliSleep(60000)

		fmt.Println("Down..")
		robotgo.MoveSmoothRelative(100, 100, 1.0, 30.0)
		robotgo.MilliSleep(60000)
	}
}
