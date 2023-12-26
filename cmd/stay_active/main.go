package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		var option int = Menu()

		switch option {
		case 1:
			fmt.Println("Work selected..")
			time.Sleep(1 * time.Second)
			ClearTerminal()
			Work()
		case 2:
			fmt.Println("Gaming selected..")
			time.Sleep(1 * time.Second)
			ClearTerminal()
			Gaming()
		case 3:
			fmt.Println("Exiting..")
			os.Exit(0)
		default:
			fmt.Println("Not an option.")
			time.Sleep(1 * time.Second)
			ClearTerminal()
		}
	}
}
