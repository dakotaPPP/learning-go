# Devlog

## 10-22-2025 commit 3871f23eaea6558dd1888c4e86d00a569170dd33

- used go.dev docs `readFile()` code for taking a file as input

## 10-22-2025 commit 7a59db5c951198dc05b7975c7e0b262d5220f01b

- left comments explaining end goal of cli tool

```
// END GOAL output
// numLimes numWords numBytes fileName or wildcards
// i.e. go run main.go * will go through all files
```

- left comments on how I believed outputs from the previous `readFile()` code snippet works

```
// presumtion seems correct as count + 1 showcases a 0 at the end of the output
// output: [104 101 108 108 111 32 116 104 105 115 32 105 115 32 97 32 102 105 108 101 32 102 111 114 109 97 116 32 108 111 108 111 108 111 108 10 0]
// so numBytes part of the output is captured for us already in count
// for calculating the numLines and numWords I'm thinking go through the data slice and increment a counter-
// whenever we run into a space ascii character or a \n ascii character
```

## 10-22-2025 commit a1071360aee189fdbdfa2ef0a597b7057b429784

- Realized there's a big error in my `readFile()` function
  - I thought it was reading the entire file but it's more a read stream system with a buffer
  - Left comments explaining this
- Also added more test files which exposes gaps in my current wc implementation

```
// THERES A BIG MISCONCEPTION I HAD ABOUT READING THE FILE
// I pressumed the slice would automatically grow in size, which it doesn't
// So this readFile more so acts as a read buffer that needs to loop till the EOF is reached
// meaning our current setup for detecting empty files is slighty off but it's not lost
```

## ~11:30 PM 10-22-2025

I know I could've just used the standard library `ReadFile()` but for learning sake I figured it be a better experience to try down the path of reading a file stream rather than a file as a whole

Resources I've used:

- go.dev/doc/
- [gobyexample.com](https://gobyexample.com)
- general google searches of what functions do (mainly did after
  examining code given to me by gobyexample or go.dev/doc
- gemini for help with why packages weren't working (I just had to
  refresh nvim I followed the tutorial correctly)
- gemini for help with why my for loop variable was undefined (I
  misspelled it haha)
- gemini for help with weird case of no new left assigned variables?
- [gemini chat](https://gemini.google.com/share/113cd9851676)

## 2:40 PM 10-23-2025

- Used [https://gobyexample.com/sorting](https://gobyexample.com/sorting) to learn how to sort my file input
- Misunderstood setting up the offset value (was doing a off by one error)
  - This fix gives me the correct bytes output for each file now
  - Meaning the only thing left for me to get right is the word count output then I need to worry about wildcards and directory scanning
- Added test cases for when input file is some multiple of the file buffer to see if an unexpected error would occur
  - This test case passed, no change was need to fix

## 3:25 PM 10-23-2025 All outputs are correct

- Redesigned my algorithm for counting words
  - First came up with track if the space is continous if so don't add to word count
  - Then came up with keeping track of if I was inside a word
    - Ended up being my solution that worked
    - Noticed that my word count numbers were a bit inflated
      - presummed it's likely due to the count word function losing context of if it was in a word before the previous execution
      - To fix this I added in passing in and outputting the context of if the buffer is still in a word or not
    - Noticed the word count was still slightly larger than expected
      - after printf debugging in `counter.go` I noticed that the null bytes of data in my `inputBytes` slice was triggering my default switch statement case
      - To fix this I added the case `'\00'` which would break out of the for loop and return the final count of accumulated values

### Next step

- Taking in file input parameter
  - will begin slow with just taking in a single cli argument with no wild cards
  - Then will begin working on taking in wildcard inputs

## 9:40 PM 10-23-2025

- Began wanting to use structs for my output of each wc entry for easier type completion
  - Made the code a bit more verbose but feels like good practice of getting the concept in
  - Then I wanted to add together all the results to get the total line that wc outputs when multiple files are inputBytes
  - I looked up if overloading operators was possible by default and found out it wasn't
    - Google ai then gave me this code that I based my add function off of

    ```go
    type Vector struct {
      X, Y float64
    }

    // Add method for Vector type
    func (v Vector) Add(other Vector) Vector {
      return Vector{v.X + other.X, v.Y + other.Y}
    }

    func main() {
      v1 := Vector{1, 2}
      v2 := Vector{3, 4}

      v3 := v1.Add(v2) // Using the custom Add method
      fmt.Println(v3)  // Output: {4 6}
    }
    ```
