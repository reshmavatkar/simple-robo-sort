package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sort(width, height, length, mass int) string {
	volume := width * height * length
	isBulky := volume >= 1000000 || width >= 150 || height >= 150 || length >= 150
	isHeavy := mass >= 20

	switch {
	case isBulky && isHeavy:
		return "REJECTED"
	case isBulky || isHeavy:
		return "SPECIAL"
	default:
		return "STANDARD"
	}
}

func readIntInput(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting.")
			os.Exit(0)
		}

		val, err := strconv.Atoi(input)
		if err != nil || val < 0 {
			fmt.Println("Invalid input. Please enter a non-negative integer.")
			continue
		}
		return val
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Package Sorter (type 'exit' at any prompt to quit)")

	for {
		width := readIntInput(reader, "Width (cm): ")
		height := readIntInput(reader, "Height (cm): ")
		length := readIntInput(reader, "Length (cm): ")
		mass := readIntInput(reader, "Mass (kg): ")

		result := sort(width, height, length, mass)
		fmt.Printf("Dispatch to: %s stack\n\n", result)
	}
}
