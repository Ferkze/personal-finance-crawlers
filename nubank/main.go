package main

import (
	"os/exec"
	"fmt"
)

func main() {
	fmt.Println(exec.Command("echo", "hello").CombinedOutput())
}
