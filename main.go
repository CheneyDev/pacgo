package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var maze []string

func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	return nil
}

func printScreen() {
	for _, line := range maze {
		fmt.Println(line)
	}
}

func main() {

	//load resources
	err := loadMaze("maze01.txt")
	if err != nil {
		log.Println("failed to load maze:", err)
		return
	}

	//game loop
	for {
		printScreen()

		break
	}
}
