// Package counter holds my function for counting the number of lines and words
package counter

func GetLineAndWordCount(inputBytes []byte, isInAWord bool) (int, int, bool) {
	lineCount := 0
	wordCount := 0

	for i := range inputBytes {
		switch inputBytes[i] {
		case '\n':
			lineCount++
			isInAWord = false
		case ' ':
			isInAWord = false
		// when end of file is reached break out of for loop as no more counting is needed
		case '\x00':
			break
		default:
			if !isInAWord {
				wordCount++
			}
			isInAWord = true
		}
	}

	return lineCount, wordCount, isInAWord
}
