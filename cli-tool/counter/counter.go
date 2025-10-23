// Package counter holds my function for counting the number of lines and words
package counter

func GetLineAndWordCount(inputBytes []byte) (int, int) {
	lineCount := 1
	wordCount := 0
	for i := range inputBytes {
		switch inputBytes[i] {
		case '\n':
			lineCount++
		case ' ':
			// know this logic isn't quite right will return to it later
			wordCount++
		}
	}

	return lineCount, wordCount
}
