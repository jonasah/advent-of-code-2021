package day2

import (
	"strconv"
	"strings"

	"github.com/jonasah/advent-of-code-2021/lib/common"
)

type command struct {
	direction string
	steps     int
}

func Part1(input string) int {
	commands := parseCommands(common.GetLines(input))

	horizontal := 0
	depth := 0
	for _, command := range commands {
		switch command.direction {
		case "forward":
			horizontal += command.steps
		case "up":
			depth -= command.steps
		case "down":
			depth += command.steps
		}
	}

	return horizontal * depth
}

func Part2(input string) int {
	commands := parseCommands(common.GetLines(input))

	horizontal := 0
	depth := 0
	aim := 0
	for _, command := range commands {
		switch command.direction {
		case "forward":
			horizontal += command.steps
			depth += aim * command.steps
		case "up":
			aim -= command.steps
		case "down":
			aim += command.steps
		}
	}

	return horizontal * depth
}

func parseCommands(lines []string) []command {
	commands := make([]command, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		commands[i] = command{direction, steps}
	}
	return commands
}
