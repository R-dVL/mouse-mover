package main

import (
	"os"
	"os/exec"
	"runtime"
)

// Multiplatform clear terminal function
func ClearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("clear")

	case "linux":
		cmd = exec.Command("clear")

	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")

	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
