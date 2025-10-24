package main

import (
	"os"

	"example.com/wc"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// [1:] is just so we don't grab the program path
	fileNames := os.Args[1:]

	totalOfAllEntries := wc.WcEntry{FileName: "total"}

	for _, fileName := range fileNames {
		entry, err := wc.GetWCData(fileName)
		check(err)
		wc.PrintWCEntry(entry)
		totalOfAllEntries = totalOfAllEntries.Add(entry)
	}

	if len(fileNames) > 1 {
		wc.PrintWCEntry(totalOfAllEntries)
	}
}
