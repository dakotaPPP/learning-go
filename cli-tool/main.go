package main

// END GOAL output
// numLines numWords numBytes fileName or wildcards
// i.e. go run main.go * will go through all files in cwd

// COMPLTED GOALS
// numBytes and fileName
import (
	"fmt"
	"io"

	"example.com/counter"
	"example.com/file"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileNames := []string{"all-space.txt", "empty-test.txt", "normal-test.txt", "too-many-spaces.txt"}

	for _, fileName := range fileNames {
		fileName = "test-inputs/" + fileName
		numLines, numWords, numBytes, finalFileName, err := getWCData(fileName)
		check(err)

		fmt.Printf("%d %d %d %s\n", numLines, numWords, numBytes, finalFileName)
	}
}

func getWCData(fileName string) (int, int, int, string, error) {
	var numLines, numWords, numBytes int
	var offset int64

	for {
		bytesRead, fileContents, err := file.ReadBuffer(fileName, offset)

		if err == io.EOF {
			break
		}

		check(err)

		bufferNumLines, bufferNumWords := counter.GetLineAndWordCount(fileContents)

		numLines += bufferNumLines
		numWords += bufferNumWords
		numBytes += bytesRead

		if bytesRead != file.BufferSize {
			break
		}
		offset += file.BufferSize + 1
	}

	return numLines, numWords, numBytes, fileName, nil
}
