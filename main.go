package main

import (
	"fmt"
	"runtime"
	"strconv"

	"snail/cmd"
)

var (
	SetCpuCount string
)

func main() {
	if SetCpuCount != "" {
		procsNum, err := strconv.Atoi(SetCpuCount)
		if err == nil {
			runtime.GOMAXPROCS(procsNum)
			fmt.Printf("GOMAXPROCS num set %v\n", procsNum)
		}
	}
	cmd.Execute()
}
