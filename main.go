package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Variables for paths we want to read
	var firstFile string
	var secondFile string

	//Channels for fgoroutines
	line1 := make(chan string)
	line2 := make(chan string)
	done1 := make(chan bool)
	done2 := make(chan bool)

	//Simple user input for paths
	fmt.Print("[!] Enter path to first file: ")
	fmt.Scanln(&firstFile)
	fmt.Print("[!] Enter path to second file: ")
	fmt.Scanln(&secondFile)

	//Start go routines
	go readFromFile(firstFile, line1, done1)
	go readFromFile(secondFile, line2, done2)
	go compareLines(line1, line2)

	<-done1
	<-done2
	close(line1)
	close(line2)
}

// Function for reading line in text one by one
func readFromFile(filename string, lines chan string, done chan bool) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}

	done <- true
}

// Compare if lines are identical and if they are print it out
func compareLines(lines1, lines2 chan string) {
	var lineNumber int
	for {
		line1, ok1 := <-lines1
		line2, ok2 := <-lines2

		if !ok1 || !ok2 {
			break
		}

		lineNumber++

		if line1 == line2 {
			fmt.Printf("[*] Found matching text: \"%s\" on line %d\n", line1, lineNumber)
		}
	}
}
