package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CardinalDirection Go-style Enum to track a cardinal direction
type CardinalDirection int

const (
	North CardinalDirection = iota
	East
	South
	West
)

func (d CardinalDirection) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

// Constants for Right ("R") and Left ("L")
const (
	Left  = "L"
	Right = "R"
)

// Instruction models a single instruction for moving
type Instruction struct {
	Direction string
	Distance  int
}

type Instructions []Instruction

// Coordinate models the current location as X and Y
type Coordinate struct{ X, Y int }

// loadInput loads the input.txt file into a string array
func loadInput(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load/open file %s: %v", filePath, err)
	}

	content := string(data)

	parts := strings.Split(content, ",")

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return parts, nil
}

// parseInstructions takes the input from the file and parses each into an Instruction
func parseInstructions(unparsedInstructions []string) (Instructions, error) {
	instructions := make(Instructions, 0, len(unparsedInstructions))

	for _, unparsedInstruction := range unparsedInstructions {
		if unparsedInstruction == "" {
			continue
		}
		if len(unparsedInstruction) < 2 {
			return nil, fmt.Errorf("invalid instruction: %s", unparsedInstruction)
		}

		direction := unparsedInstruction[:1]
		distanceStr := unparsedInstruction[1:]

		if direction != "R" && direction != "L" {
			return nil, fmt.Errorf("invalid direction: %s", direction)
		}

		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			return nil, fmt.Errorf("invalid distance: %s", distanceStr)
		}

		instructions = append(instructions, Instruction{Direction: direction, Distance: distance})
	}

	return instructions, nil
}

// changeDirection updates the cardinal direction based on which way we turn
func changeDirection(currentDirection CardinalDirection, direction string) (CardinalDirection, error) {
	switch direction {
	case Left:
		return (currentDirection - 1 + 4) % 4, nil
	case Right:
		return (currentDirection + 1) % 4, nil
	default:
		return 0, fmt.Errorf("invalid direction: %s", direction)
	}
}

// move updates the current location based on our current facing direction and the distance needed to travel
func move(facing CardinalDirection, currentLocation Coordinate, distance int) (Coordinate, error) {
	switch facing {
	case North:
		return Coordinate{currentLocation.X, currentLocation.Y + distance}, nil
	case South:
		return Coordinate{currentLocation.X, currentLocation.Y - distance}, nil
	case West:
		return Coordinate{currentLocation.X - distance, currentLocation.Y}, nil
	case East:
		return Coordinate{currentLocation.X + distance, currentLocation.Y}, nil
	default:
		return Coordinate{}, fmt.Errorf("invalid direction: %s", facing)
	}

	return nil
}

// calculateBlocksAway calculates the total number of blocks away we are from the start to the final location
func calculateBlocksAway(finalLocation Coordinate) int {
	if finalLocation.X < 0 {
		finalLocation.X = -finalLocation.X
	}
	if finalLocation.Y < 0 {
		finalLocation.Y = -finalLocation.Y
	}
	return finalLocation.X + finalLocation.Y
}

func run(part int) error {
	unparsedDirections, err := loadInput("input.txt")
	if err != nil {
		return err
	}

	directions, err := parseInstructions(unparsedDirections)
	if err != nil {
		return err
	}

	if part == 1 {
		facing := North
		currentLocation := Coordinate{X: 0, Y: 0}
		for _, instruction := range directions {
			facing, err = changeDirection(facing, instruction.Direction)
			if err != nil {
				return err
			}

			currentLocation, err = move(facing, currentLocation, instruction.Distance)
			if err != nil {
				return err
			}
		}

		blocksAway := calculateBlocksAway(currentLocation)
		fmt.Printf("Blocks away: %d\n", blocksAway)

	} else if part == 2 {
		return fmt.Errorf("part 2 not implemented")
	} else {
		return fmt.Errorf("invalid part: %d", part)
	}

	return nil
}

func main() {
	part := flag.Int("part", 1, "Run part 1 or 2")
	flag.Parse()

	if err := run(*part); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
