package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	day2()

}

func day2() {
	contents, err := os.Open("day2/day2.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return

	}
	scanner := bufio.NewScanner(contents)
	total := 0.0
	for scanner.Scan() {
		line := scanner.Text()
		arr := findMin(line)
		total += multiplyArrayElements(arr)

	}
	fmt.Println(total)
}

func findMin(line string) map[string]float64 {
	min := map[string]float64{"red": 0, "green": 0, "blue": 0}
	initalSplit := strings.Split(line, ":")
	secondSplit := strings.Split(initalSplit[1], ";")
	for _, element := range secondSplit {
		re := regexp.MustCompile("(\\d+)\\s*(red|green|blue)")
		items := re.FindAllString(element, -1)
		for _, item := range items {
			split := strings.Split(item, " ")
			value, colour := split[0], split[1]
			number, err := strconv.ParseFloat(value, 64)
			if err == nil {
				min[colour] = math.Max(min[colour], number)
			}
		}

	}
	return min

}

func multiplyArrayElements(arr map[string]float64) float64 {
	result := 1.0
	for _, v := range arr {
		result *= v
	}
	return result
}

func day1() {
	contents, err := os.Open("day2/day2.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return

	}
	scanner := bufio.NewScanner(contents)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if isPossible(line) {
			total += getId(line)
		}

	}
	fmt.Println(total)

}

func getId(line string) int {
	re := regexp.MustCompile("[0-9]+")
	number := re.FindString(line)
	ans, err := strconv.Atoi(number)
	if err == nil {
		return ans
	}

	fmt.Println("Couldn't process the number ", ans)
	return 0

}

func isPossible(line string) bool {
	check := map[string]int{"red": 12, "green": 13, "blue": 14}
	max := 39
	initalSplit := strings.Split(line, ":")
	secondSplit := strings.Split(initalSplit[1], ";")
	for _, element := range secondSplit {
		re := regexp.MustCompile("(\\d+)\\s*(red|green|blue)")
		items := re.FindAllString(element, -1)
		total := 0
		for _, item := range items {
			split := strings.Split(item, " ")
			value, colour := split[0], split[1]
			number, err := strconv.Atoi(value)
			if err == nil && check[colour] < number {
				return false
			}
			total += number
		}
		if total > max {
			return false
		}

	}
	return true
}
