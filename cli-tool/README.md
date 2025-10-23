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
