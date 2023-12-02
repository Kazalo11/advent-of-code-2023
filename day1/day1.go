package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return

	}
	scanner := bufio.NewScanner(contents)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(findLineNumber(line))
		total += findLineNumber(line)
	}
	fmt.Println(total)

}

func findLineNumber(line string) int {
	wordNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digitNumbers := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	firstWord := ""
	lastWord := ""
	start := 1000000
	end := -1
	lineNumber := ""
	firstConv := false
	lastConv := false
	firstElement := 0
	lastElement := 0
	for index, element := range wordNumbers {
		firstIndex := strings.Index(line, element)
		lastIndex := strings.LastIndex(line, element)
		if firstIndex < start && firstIndex != -1 {
			start = firstIndex
			firstWord = element
			firstConv = true
			firstElement = index
		}
		if lastIndex > end && lastIndex != -1 {
			end = lastIndex
			lastWord = element
			lastConv = true
			lastElement = index
		}
	}
	for _, element := range digitNumbers {
		firstIndex := strings.Index(line, element)
		lastIndex := strings.LastIndex(line, element)
		if firstIndex < start && firstIndex != -1 {
			start = firstIndex
			firstWord = element
			firstConv = false
		}
		if lastIndex > end && lastIndex != -1 {
			end = lastIndex
			lastWord = element
			lastConv = false
		}
	}
	if firstConv {
		lineNumber += digitNumbers[firstElement]
	} else {
		lineNumber += firstWord
	}
	if lastConv {
		lineNumber += digitNumbers[lastElement]

	} else {
		lineNumber += lastWord

	}
	ans, err := strconv.Atoi(lineNumber)
	if err == nil {
		fmt.Println(ans)
		return ans
	}
	fmt.Println("Couldn't process the number ", lineNumber)
	return 0

	// lineNumber := ""
	// for _, char := range line {
	// 	if unicode.IsDigit(char) {
	// 		lineNumber += string(char)
	// 		break
	// 	}

	// }

	// for i := len(line) - 1; i >= 0; i-- {
	// 	if unicode.IsDigit(rune(line[i])) {
	// 		lineNumber += string(rune(line[i]))
	// 		break
	// 	}
	// }
	// ans, err := strconv.Atoi(lineNumber)
	// if err == nil {
	// 	return ans
	// }
	// fmt.Println("Couldn't process the number ", lineNumber)
	// return 0

}
