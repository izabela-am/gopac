package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var maze []string

	func renderMaze(file string) error {
		// os.Open() returns a file and an error
		f, err := os.Open(file)

		if err != nil {
			return err
		}
		
		defer f.Close() // puts f.close() in the call stack

		// reads all lines of the txt file and appends them to the maze slice
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			maze = append(maze, line)
		}

		return nil
	}

	for {
		fmt.Println("Hello, Pac Go!")
		break
	}
}
