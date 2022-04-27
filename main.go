package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Main function to start, pause or stop the pomodoro.

	var instruction = flag.String(
		"instruction",
		"start",
		"Instruction to start or stop a pomodoro",
	)
	flag.Parse()
	var instruction_value = *instruction
	var no_valid_option = errors.New(
		"No instruction provided, please use start or stop")

	if instruction_value != " " {
		instruction_value = strings.ToLower(instruction_value)
	}

	switch instruction_value {
	case "start":
		fmt.Println("Start Pomodoro")
	case "stop":
		fmt.Println("Stop Pomodoro")
	default:
		fmt.Print(no_valid_option)
	}
}
