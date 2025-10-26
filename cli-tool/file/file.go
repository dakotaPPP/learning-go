// Package file hosts all the file operations i.e. reading buffers or parsing through a current directory
package file

import (
	"io"
	"os"
)

const BufferSize = 64

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckIfDir(fileName string) bool {
	result, err := os.Stat(fileName)
	check(err)

	return result.IsDir()
}

func ReadBuffer(fileName string, offset int64) (int, []byte, error) {
	file, err := os.Open(fileName)
	check(err)

	// jump to offset in file
	_, err = file.Seek(offset, io.SeekStart)
	check(err)

	data := make([]byte, BufferSize)
	count, err := file.Read(data)

	if err == io.EOF {
		return -1, data, err
	}

	check(err)

	file.Close()
	// I pressume the :count in data is acting as the null terminator point in the files input
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])
	// fmt.Printf("read %d bytes: %v\n", count, data[:count+1])

	// presumtion seems correct as count + 1 showcases a 0 at the end of the output
	// output: [104 101 108 108 111 32 116 104 105 115 32 105 115 32 97 32 102 105 108 101 32 102 111 114 109 97 116 32 108 111 108 111 108 111 108 10 0]

	// so numBytes part of the output is captured for us already in count
	// for calculating the numLines and numWords I'm thinking go through the data slice and increment a counter-
	// whenever we run into a space ascii character or a \n ascii character

	// THERES A BIG MISCONCEPTION I HAD ABOUT READING THE FILE
	// I pressumed the slice would automatically grow in size, which it doesn't
	// So this readFile more so acts as a read buffer that needs to loop till the EOF is reached
	// meaning our current setup for detecting empty files is slighty off but it's not lost
	return count, data, nil
}
