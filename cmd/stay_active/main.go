package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		var mode int = Menu()

		switch mode {
		case 1:
			fmt.Println("Work selected..")
			time.Sleep(1 * time.Second)
			ShutdownScheduler()
			Work()

		case 2:
			fmt.Println("Gaming selected..")
			time.Sleep(1 * time.Second)
			ShutdownScheduler()
			Gaming()

		case 3:
			fmt.Println("Mouse mover selected..")
			time.Sleep(1 * time.Second)
			ShutdownScheduler()
			MouseMover()

		case 4:
			fmt.Println("Exiting..")
			os.Exit(0)

		default:
			fmt.Println("Not an option.")
			time.Sleep(1 * time.Second)
			ClearTerminal()
		}
	}
}
