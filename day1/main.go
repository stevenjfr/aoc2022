package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// readLines reads a whole file into memory
func readLines(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var biggestElf, currentElf int
	var threeElves = []int{0}

	for scanner.Scan() {

		i, err := strconv.Atoi((scanner.Text()))

		if err != nil {
			// handle error
			if currentElf >= threeElves[0] {
				fmt.Printf("%s is bigger than %s \n", currentElf, biggestElf)

				threeElves = append(threeElves, currentElf)
				currentElf = 0
				fmt.Printf("BREAK")
			} else {
				fmt.Println("not biggest")
				currentElf = 0
			}
		} else {
			fmt.Printf("%s\n", i)
			currentElf += i
		}

	}
	sort.Ints(threeElves)
	fmt.Println(threeElves)
	return nil
}
func code2() {
	fmt.Println("Starting script")
	err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println("drop into a shell and add 'em up!")
	fmt.Println("finished script")

}
func main() {

	code2()
}
