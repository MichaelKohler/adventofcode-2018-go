package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  fmt.Printf("Day 1 - Puzzle 1\n")
  file, err := os.Open("inputs/day1-puzzle1.txt")
  check(err)

  defer file.Close()

  var totalFrequency int = 0;

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      i, err := strconv.Atoi(scanner.Text())
      check(err)

      totalFrequency += i
      fmt.Printf("Frequency after iteration: ")
      fmt.Println(totalFrequency)
  }
}

func check(e error) {
  if e != nil {
      panic(e)
  }
}