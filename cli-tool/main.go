package main

// END GOAL output
// numLines numWords numBytes fileName or wildcards
// i.e. go run main.go * will go through all files in cwd

// COMPLTED GOALS
// numBytes and fileName
import (
	"fmt"
	"io"
	"slices"

	"example.com/counter"
	"example.com/file"
)

type wcEntry struct {
	numLines int
	numWords int
	numBytes int
	fileName string
}

func printWCEntry(entry wcEntry) {
	fmt.Printf("%d %d %d %s\n", entry.numLines, entry.numWords, entry.numBytes, entry.fileName)
}

/* Probably could make this have a side effect of updating the input e's
* value but seems like bad practice as maybe this isn't wanted */
func (e wcEntry) Add(other wcEntry) wcEntry {
	// first entry's fileName is kept
	return wcEntry{e.numLines + other.numLines, e.numWords + other.numWords, e.numBytes + other.numBytes, e.fileName}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileNames := []string{"all-space.txt", "empty-test.txt", "normal-test.txt", "mid-size.txt", "too-many-spaces.txt", "buffer-size.txt", "buffer-size-2x.txt"}
	slices.Sort(fileNames)
	totalOfAllEntries := wcEntry{fileName: "total"}

	for _, fileName := range fileNames {
		fileName = "test-inputs/" + fileName
		entry, err := getWCData(fileName)
		check(err)
		printWCEntry(entry)
		totalOfAllEntries = totalOfAllEntries.Add(entry)
	}

	if len(fileNames) > 1 {
		printWCEntry(totalOfAllEntries)
	}
}

func getWCData(fileName string) (wcEntry, error) {
	entry := wcEntry{fileName: fileName}

	var offset int64
	var isInAWordTemp bool

	for {
		bytesRead, fileContents, err := file.ReadBuffer(fileName, offset)

		if err == io.EOF {
			break
		}

		check(err)

		var bufferNumLines, bufferNumWords int

		bufferNumLines, bufferNumWords, isInAWordTemp = counter.GetLineAndWordCount(fileContents, isInAWordTemp)

		entry.numLines += bufferNumLines
		entry.numWords += bufferNumWords
		entry.numBytes += bytesRead

		if bytesRead != file.BufferSize {
			break
		}
		offset += file.BufferSize
	}

	return entry, nil
}
