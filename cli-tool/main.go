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
	cliArgs := os.Args[1:]

	if len(cliArgs) != 1 {
		errString := `Incorrect number of cli inputs!
Please use only one argument
example: 
	./wc file.txt`
		panic(errString)
	}

	userInput := cliArgs[0]
	// fileNames := []string{"all-space.txt", "empty-test.txt", "normal-test.txt", "mid-size.txt", "too-many-spaces.txt", "buffer-size.txt", "buffer-size-2x.txt"}
	// slices.Sort(fileNames)
	var fileNames []string
	fileNames = append(fileNames, userInput)
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
