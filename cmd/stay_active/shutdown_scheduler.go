package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Multiplatform shutdown scheduler function
func ShutdownScheduler() {
	ClearTerminal()

out:
	for {
		var option string

		fmt.Print("Do you want to schedule a computer shutdown?(y/n): ")
		fmt.Scan(&option)

		switch option {
		case "y":
			var time int

			fmt.Print("Enter minutes until shutdown: ")
			fmt.Scan(&time)

			// Minutes to seconds
			time = time * 60

			var cmd *exec.Cmd

			switch runtime.GOOS {
			case "darwin":
				cmd = exec.Command(fmt.Sprintf("sudo shutdown -h %d", time))
			case "linux":
				cmd = exec.Command(fmt.Sprintf("sudo shutdown -h %d", time))
			case "windows":
				cmd = exec.Command("cmd", "/c", fmt.Sprintf("shutdown /s /t %d", time))
			default:
				fmt.Println("Not supported platform.")
			}

			cmd.Stdout = os.Stdout
			cmd.Run()
			break out

		case "n":
			break out

		default:
			fmt.Println("Enter a valid option.")
		}
	}
}
