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
	if d < North || d > West {
		return fmt.Sprintf("Unknown Cardinal Direction: CardinalDirection(%d)", d)
	}
	return [...]string{"N", "E", "S", "W"}[d]
}

// Turn represents L or R
type Turn byte

const (
	Left  Turn = 'L'
	Right Turn = 'R'
)

// Instruction models a single instruction for moving
type Instruction struct {
	Direction Turn
	Distance  int
}

type Instructions []Instruction

// Coordinate models the current location as X and Y
type Coordinate struct{ X, Y int }

type VisitedLocations map[Coordinate]bool

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

		direction := Turn(unparsedInstruction[0])
		distanceStr := unparsedInstruction[1:]

		if direction != Right && direction != Left {
			return nil, fmt.Errorf("invalid direction: %q", direction)
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
func changeDirection(currentDirection CardinalDirection, direction Turn) (CardinalDirection, error) {
	switch direction {
	case Left:
		return (currentDirection - 1 + 4) % 4, nil
	case Right:
		return (currentDirection + 1) % 4, nil
	default:
		return 0, fmt.Errorf("invalid direction: %q", direction)
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
		return Coordinate{}, fmt.Errorf("invalid direction: %v", facing)
	}
}

// calculateBlocksAway calculates the total number of blocks away we are from the start to the final location
func calculateBlocksAway(finalLocation Coordinate) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	return abs(finalLocation.X) + abs(finalLocation.Y)
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

	} else {
		facing := North
		startingLocation := Coordinate{X: 0, Y: 0}
		visited := make(VisitedLocations)
		visited[startingLocation] = true

		currentLocation := startingLocation
		for _, instruction := range directions {
			facing, err = changeDirection(facing, instruction.Direction)
			if err != nil {
				return err
			}

			for i := 1; i <= instruction.Distance; i++ {
				currentLocation, err = move(facing, currentLocation, 1)
				if err != nil {
					return err
				}

				_, ok := visited[currentLocation]
				if ok {
					blocksAway := calculateBlocksAway(currentLocation)
					fmt.Printf("Location %v has been visited prior\n", currentLocation)
					fmt.Printf("Blocks away: %d\n", blocksAway)
					return nil
				}

				visited[currentLocation] = true
			}
		}

		return nil
	}

	return nil
}

func main() {
	part := flag.Int("part", 1, "Run Part 1 or 2")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: go run main.go [--part=1|2]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *part != 1 && *part != 2 {
		fmt.Fprintf(os.Stderr, "invalid --part %d; must be 1 or 2\n", *part)
		flag.Usage()
		os.Exit(2)
	}

	fmt.Printf("Running Day 01 Part %d\n", *part)
	if err := run(*part); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
