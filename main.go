package main

import (
	"fmt"
	"os"

	"github.com/edlingao/advent-of-code-2024-golang/challenges"
)

func main() {
  args := os.Args

  if len(args) < 2 {
    fmt.Println("Please provide a day")
    return
  }
  
  day := args[1]
  switch day {
  case "1":
    challenges.Day1(true)
  case "1.1": 
    challenges.Day1(false)
  default:
    fmt.Println("Invalid day")
  }
}
