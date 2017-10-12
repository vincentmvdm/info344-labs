package main

import (
	"net/http"
	"os"

	"github.com/info344-a17/info344-labs/lab01-testing/handlers"
)

// This program scans a text file for swears
func main() {

	//get the value of the ADDR environment variable
	//and use that as the address this server will listen on
	addr := os.Getenv("ADDR")
	//if not set, default to ":443", which means listen for
	//all requests to all hosts on port 443
	if len(addr) == 0 {
		addr = "localhost:4242"
	}

	mux := http.NewServeMux()

	// TODO: read known swears from a file
	knownSwears := loadSwears("test.txt")
	//create a new handlers.SwearsHandler struct
	//since that is in a different package, use the
	//package name as a prefix, and import the package above
	swearsHandler := &handlers.SwearsHandler{
		KnownSwears: knownSwears,
	}

	mux.Handle("/swears", swearsHandler)
}

func loadSwears(fileName string) []string {
	// // get the text file to scan
	// fd := os.Args[1]
	// // b, _ := ioutil.ReadFile(f)

	// f, _ := os.Open(fd)
	// scanner := bufio.NewScanner(f)

	// // Set the Split method to ScanWords.
	// scanner.Split(bufio.ScanWords)

	// // Scan all words from the file.
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// }
	knownSwears := []string{"fuck", "shit", "damn", "hell"}
	return knownSwears
}
