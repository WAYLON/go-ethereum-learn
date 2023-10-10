package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	for {
		format := time.Now().Format("15:04:05")
		time.Sleep(100 * time.Millisecond)
		fmt.Println(format)
		if format == "14:59:59" {
			time.Sleep(800 * time.Millisecond)
			cmd := exec.Command("/bin/bash", "-c", "/Users/waylon/Desktop/dy.sh")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			e := cmd.Run()
			CheckError(e)
			break
		}
	}
	time.Sleep(10 * time.Second)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
