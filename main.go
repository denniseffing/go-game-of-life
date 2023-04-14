package main

import (
	"os"
	"os/exec"
	"time"
)

func main() {
	slice := make([][]cell, 20)
	for i := range slice {
		slice[i] = make([]cell, 100)
	}

	slice[0][1].alive = true
	slice[1][2].alive = true
	slice[2][0].alive = true
	slice[2][1].alive = true
	slice[2][2].alive = true

	world := world{slice}

	for {
		print(world.String())
		world = world.age()
		time.Sleep(100 * time.Millisecond)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
