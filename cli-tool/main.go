package main

// END GOAL output
// numLines numWords numBytes fileName or wildcards
// i.e. go run main.go * will go through all files in cwd

// COMPLTED GOALS
// numBytes and fileName
import (
	"fmt"
	"log"
	"os"

	"example.com/counter"
)

func main() {
	fileName := "file.txt"
	numBytes, fileContents, err := readFile(fileName)
	if err != nil {
		panic(err)
	}

	var numLines, numWords int
	numLines, numWords = counter.GetLineAndWordCount(fileContents)
	fmt.Printf("%d %d %d %s\n", numLines, numWords, numBytes, fileName)
}

func readFile(fileName string) (int, []byte, error) {
	file, err := os.Open(fileName)
	data := make([]byte, 64)

	if err != nil {
		log.Fatal(err)
		return -1, data, err
	}

	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
		return -1, data, err
	}
	// I pressume the :count in data is acting as the null terminator point in the files input
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])
	// fmt.Printf("read %d bytes: %v\n", count, data[:count+1])

	// presumtion seems correct as count + 1 showcases a 0 at the end of the output
	// output: [104 101 108 108 111 32 116 104 105 115 32 105 115 32 97 32 102 105 108 101 32 102 111 114 109 97 116 32 108 111 108 111 108 111 108 10 0]

	// so numBytes part of the output is captured for us already in count
	// for calculating the numLines and numWords I'm thinking go through the data slice and increment a counter-
	// whenever we run into a space ascii character or a \n ascii character
	return count, data, nil
}
