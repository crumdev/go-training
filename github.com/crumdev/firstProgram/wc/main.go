package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// TODO: main contents
	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count lines")

	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type tp words (default is split by lines)
	if !countLines {
		if !countBytes {
			scanner.Split(bufio.ScanWords)
		} else {
			scanner.Split(bufio.ScanBytes)
		}
	}

	// defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	//return the total
	return wc
}
