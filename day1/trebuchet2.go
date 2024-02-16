package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// This works fine in testing but not in practice?
func stringToNumber(input string) string {
	_, err := strconv.Atoi(input)
	if err != nil {
		switch input {
		case "one":
			return "1"
		case "two":
			return "2"
		case "three":
			return "3"
		case "four":
			return "4"
		case "five":
			return "5"
		case "six":
			return "6"
		case "seven":
			return "7"
		case "eight":
			return "8"
		case "nine":
			return "9"
		}
	}
	return input
}

func part2(input ...string) {
	numbers := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
	var result = 0
	for _, str := range input {
		match := numbers.FindAllStringSubmatch(str, -1)
		var nums strings.Builder
		nums.WriteString(stringToNumber(match[0][0]) + stringToNumber(match[len(match)-1][0]))

		numsAsInt, err := strconv.Atoi(nums.String())
		if err != nil {
			println("Issue with converting to an int: %q", err)
			return
		}
		result = result + numsAsInt
	}
	fmt.Printf("\nFinal Result:\n%d", result)
}

func Trebuchet2() {
	fmt.Println("Input text for Trebuchet 2")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	part2(lines...)
}
