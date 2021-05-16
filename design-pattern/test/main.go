package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	//var str strings.Builder
	//for i := 0; i < 1000; i++ {
	//	str.WriteString("a")
	//}
	//fmt.Println(str.String())
	var stdOut, stdErr bytes.Buffer
	cmd := exec.Command("sh", "-c", "dstat -n 1 1 | awk 'END {print}' | awk '{print $2}'")
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil {
		fmt.Errorf("error %v", err)
	}
	fmt.Println(stdOut.String())
}
