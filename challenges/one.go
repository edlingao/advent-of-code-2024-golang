package challenges

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Day1(part1 bool) {
  file, err := os.ReadFile("./challenges/inputs/dayOne.txt")
  if err != nil {
    log.Panic("Error:", err.Error())
  }

  input := string(file)
  lines := strings.Split(input, "\n")
  var left_nums []int
  var right_nums []int

  for _, line := range lines {
    numbers := strings.Split(line, "   ")
    if len(numbers) < 2 {
      continue
    }

    left := numbers[0]
    right := numbers[1]

    left_nums = append(left_nums, transformInt(left))
    right_nums = append(right_nums, transformInt(right))
  }

  if part1 {
    Day1Part1(left_nums, right_nums)
  } else {
    Day1Part2(left_nums, right_nums)
  }

}

func Day1Part1(left_nums []int, right_nums []int) {
  var sum int = 0
  for i := 0; i < 1000; i++ {
    new_left_nums, smallest_left := smallestNum(left_nums)
    new_right_nums, smallest_right := smallestNum(right_nums)
    
    if smallest_left > smallest_right {
      sum += smallest_left - smallest_right
    } else {
      sum += smallest_right - smallest_left
    }

    left_nums = new_left_nums
    right_nums = new_right_nums
  }

  log.Println("Difference: ", sum)
}

func Day1Part2(left_nums []int, right_nums []int) {
  similarity_score := 0
  for _, left_num := range left_nums {
    similarity_score += countInstancesOfNumbers(left_num, right_nums)
  }

  log.Println("Similarity Score: ", similarity_score)
}

func transformInt(num string) int {
  i, err := strconv.Atoi(num) 

  if err != nil {
    log.Panic(err.Error())
  }

  return i
}

func smallestNum(arr []int) ([]int, int) {
  smallest_number := 9999999
  smallest_index := 0
  var new_array []int

  for index, number := range arr {
    if number < smallest_number {
      smallest_number = number
      smallest_index = index
    }
  }

  new_array = append(arr[:smallest_index], arr[smallest_index+1:]...)

  return new_array, smallest_number
}

func countInstancesOfNumbers(number int, arr []int) int {
  ocurrence := 0
  for _, numberOnArray := range arr {
    if numberOnArray == number {
      ocurrence += 1
    }
  }

  return ocurrence * number
}
