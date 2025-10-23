package main

// END GOAL output
// numLimes numWords numBytes fileName or wildcards
// i.e. go run main.go * will go through all files in cwd
import (
	"fmt"
	"log"
	"os"
)

func main() {
	readFile("file.txt")
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 64)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// I pressume the :count in data is acting as the null terminator point in the files input
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	fmt.Printf("read %d bytes: %v\n", count, data[:count+1])

	// presumtion seems correct as count + 1 showcases a 0 at the end of the output
	// output: [104 101 108 108 111 32 116 104 105 115 32 105 115 32 97 32 102 105 108 101 32 102 111 114 109 97 116 32 108 111 108 111 108 111 108 10 0]

	// so numBytes part of the output is captured for us already in count
	// for calculating the numLines and numWords I'm thinking go through the data slice and increment a counter-
	// whenever we run into a space ascii character or a \n ascii character
}
