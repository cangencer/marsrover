package main

import (
	"./marsrover"
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func readNextInt(scanner *bufio.Scanner) (int, error) {
	if scanner.Scan() {
		return strconv.Atoi(scanner.Text())
	}
	return 0, errors.New("No more text to read.")
}

func readNextString(scanner *bufio.Scanner) (string, error) {
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", errors.New("No more text to read.")
}

func main() {
	const input = `5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	x, _ := readNextInt(scanner)
	y, _ := readNextInt(scanner)

	planet := marsrover.NewPlanet(x, y)

	for true {
		startX, err := readNextInt(scanner)
		if err != nil {
			break
		}
		startY, _ := readNextInt(scanner)
		startDirection, _ := readNextString(scanner)

		robot := marsrover.NewRobot(planet, startX, startY, marsrover.Direction(startDirection))
		commands, _ := readNextString(scanner)

		for i := 0; i < len(commands); i++ {
			robot.ExecuteCommand(commands[i : i+1])
		}
		fmt.Println(robot.Location())
	}
}
