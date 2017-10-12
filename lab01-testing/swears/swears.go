package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type KnownSwears map[string]bool
type SwearCounts map[string]int

// This program scans a text file for swears
func main() {
	// read known swears from a file
	knownSwears := loadKnownSwears("known_swears.txt")
	log.Printf("loaded %d swear words", len(knownSwears))

	words := loadWords() // load the
	log.Printf("loaded %d words to test", len(words))

	swearCounts := countSwears(knownSwears, words)

	for k, v := range swearCounts {
		fmt.Printf("%s %d\n", k, v)
	}

}

func countSwears(ks KnownSwears, w []string) SwearCounts {
	swearCounts := SwearCounts{}

	// check each word to see if it is a swear word
	for _, word := range w {
		if ks[word] {
			swearCounts[word]++
		}
	}

	return swearCounts
}

func loadKnownSwears(fd string) KnownSwears {

	// make a map to store all the swears (faster lookup than array)
	knownSwears := make(KnownSwears)
	scanner := openFile(fd)
	// Scan all words from the file.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		knownSwears[line] = true
	}

	return knownSwears
}

func loadWords() []string {
	words := make([]string, 3) // hmm
	scanner := openFile(os.Args[1])

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)
	}

	return words
}

func openFile(fd string) *bufio.Scanner {
	f, _ := os.Open(fd) // hmmm
	scanner := bufio.NewScanner(f)

	// Set the Split method to ScanWords.
	scanner.Split(bufio.ScanWords)
	return scanner
}
