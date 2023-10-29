package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var maze []string

func initialise()  {
	cdTerm:=exec.Command("stty","cbreak","-echo")
	cdTerm.Stdin=os.Stdin

	err:=cdTerm.Run()
	if err != nil {
		log.Fatalln("unable to active cbreak mode: "err)
	}
}

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
