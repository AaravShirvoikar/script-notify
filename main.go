package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./script-notify <script> <desired_output>")
		os.Exit(1)
	}

	script := os.Args[1]
	desiredOutput := os.Args[2]

	monitor(script, desiredOutput)
}

func monitor(script string, desiredOutput string) {
	fmt.Printf("Monitoring script: %s\n", script)

	cmd := exec.Command(script)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if strings.Contains(line, desiredOutput) {
			notify("Desired output found: " + line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
}

func notify(message string) {
	cmd := exec.Command("notify-send", "SCRIPT-NOTIFY", message)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
