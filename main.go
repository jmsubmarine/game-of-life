package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Alive = 'O'
	Dead  = ' '
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialize a generator for random numbers
	size, generationSteps := inputParameters()

	currentGeneration := generateFirstGeneration(size)

	for i := 0; i < generationSteps; i++ {
		clearConsole()
		printState(i+1, currentGeneration)
		currentGeneration = generateNextGeneration(currentGeneration)
		time.Sleep(500 * time.Millisecond)
	}
}

// Gets inputs with parameters from a user
func inputParameters() (int, int) {
	size := inputPositiveInt("Enter the size of the universe: ")
	generationSteps := inputPositiveInt("Enter a number of generation steps: ")
	return size, generationSteps
}

// Checks for positive parameters
func inputPositiveInt(prompt string) int {
	var value int
	for {
		fmt.Print(prompt)
		fmt.Scan(&value)
		if value > 0 {
			break
		}
		fmt.Println("Invalid input. Please enter a positive number.")
	}
	return value
}

// Creates an initial universe
func generateFirstGeneration(size int) [][]rune {
	firstGeneration := make([][]rune, size)

	for i := range firstGeneration {
		firstGeneration[i] = make([]rune, size)
		for j := range firstGeneration[i] {
			if rand.Intn(2) == 0 {
				firstGeneration[i][j] = Dead
			} else {
				firstGeneration[i][j] = Alive
			}
		}
	}
	return firstGeneration
}

// Counts alive neighbors for every cell
func countAliveNeighbors(i, j int, universe [][]rune) int {
	n := len(universe)
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}
	count := 0

	for _, dir := range directions {
		ni := (i + dir[0] + n) % n
		nj := (j + dir[1] + n) % n

		if universe[ni][nj] == Alive {
			count++
		}
	}
	return count
}

// Calculates the total number of alive cells in the universe
func countAliveCells(universe [][]rune) int {
	n := len(universe)
	liveCells := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if universe[i][j] == Alive {
				liveCells++
			}
		}
	}
	return liveCells
}

// Computes the next state of the universe
func generateNextGeneration(currentGeneration [][]rune) [][]rune {
	n := len(currentGeneration)
	nextGeneration := make([][]rune, n)

	for i := 0; i < n; i++ {
		nextGeneration[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			liveNeighbors := countAliveNeighbors(i, j, currentGeneration)
			if currentGeneration[i][j] == Alive {
				if liveNeighbors == 2 || liveNeighbors == 3 {
					nextGeneration[i][j] = Alive
				} else {
					nextGeneration[i][j] = Dead
				}
			} else {
				if liveNeighbors == 3 {
					nextGeneration[i][j] = Alive
				} else {
					nextGeneration[i][j] = Dead
				}
			}
		}
	}

	return nextGeneration
}

// Displays the current generation number, alive cell count and the universe
func printState(generation int, universe [][]rune) {
	fmt.Printf("Generation: #%d\n", generation)
	fmt.Printf("Alive: %d\n", countAliveCells(universe))
	printUniverse(universe)
}

// Prints the current state of the universe
func printUniverse(generation [][]rune) {
	for _, row := range generation {
		fmt.Println(string(row))
	}
}

// Clears the console after every generation of the universe
func clearConsole() {
	fmt.Print("\033[H\033[2J")
}
