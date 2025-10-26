/* Package wc holds everything related to the
*formatting and aggregating of wc data */
package wc

import (
	"fmt"
	"io"

	"example.com/counter"
	"example.com/file"
)

type WcEntry struct {
	numLines int
	numWords int
	numBytes int
	FileName string
}

func PrintWCEntry(entry WcEntry) {
	fmt.Printf("%6d %6d %6d %s\n", entry.numLines, entry.numWords, entry.numBytes, entry.FileName)
}

/* Probably could make this have a side effect of updating the input e's
* value but seems like bad practice as maybe this isn't wanted */
func (e WcEntry) Add(other WcEntry) WcEntry {
	// first entry's FileName is kept
	return WcEntry{e.numLines + other.numLines, e.numWords + other.numWords, e.numBytes + other.numBytes, e.FileName}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetWCData(fileName string) (WcEntry, bool) {
	entry := WcEntry{FileName: fileName}
	isDir := false

	if file.CheckIfDir(fileName) {
		isDir = true
		return entry, isDir
	}

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

	return entry, isDir
}
