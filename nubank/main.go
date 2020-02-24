package main

import (
	"exec"
)

func main() {
	exec.Command("echo", "hello")
}
