package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  fmt.Printf("Day 1 - Puzzle 2\n")
  INPUT_LINES := 959;
  currentFrequency := 0;
  seenTotalFrequencies := make([]int, INPUT_LINES)
  currentLineCount := 0

  for true {
    currentIterationEntries := make([]int, INPUT_LINES)
    if currentLineCount > 0 {
      // Learning: append works nicely if you actually assign it to the variable!!
      seenTotalFrequencies = append(seenTotalFrequencies, currentIterationEntries...)
    }

    file, err := os.Open("inputs/day1-puzzle2.txt")
    check(err)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    // Learning: scanner won't do it if we don't open the file again.. hah!
    for scanner.Scan() {
      i, err := strconv.Atoi(scanner.Text())
      check(err)

      currentFrequency += i
      isDuplicate := intInSlice(currentFrequency, seenTotalFrequencies);
      if (isDuplicate) {
        fmt.Println("Found first duplicate:", currentFrequency);
        return;
      }


      seenTotalFrequencies[currentLineCount] = currentFrequency
      currentLineCount++;
    }
  }
}

func intInSlice(a int, list []int) bool {
  for _, b := range list {
    if b == a {
        return true
    }
  }
  return false
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}