package main

import (
	"fmt"
	"os"

	"example.com/wc"
)

func main() {
	// [1:] is just so we don't grab the program path
	fileNames := os.Args[1:]

	totalOfAllEntries := wc.WcEntry{FileName: "total"}

	for _, fileName := range fileNames {
		entry, isDirectory := wc.GetWCData(fileName)

		wc.PrintWCEntry(entry)
		if !isDirectory {
			totalOfAllEntries = totalOfAllEntries.Add(entry)
		} else {
			fmt.Printf("wc: %s: Is a directory\n", fileName)
		}
	}

	if len(fileNames) > 1 {
		wc.PrintWCEntry(totalOfAllEntries)
	}
}
