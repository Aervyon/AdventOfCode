package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func problem(input ...string) {
	numbers := regexp.MustCompile("[0-9]")
	var result = 0
	for _, str := range input {
		match := numbers.FindAllStringSubmatch(str, -1)
		var nums strings.Builder
		nums.WriteString(match[0][0] + match[len(match)-1][0])

		numsAsInt, err := strconv.Atoi(nums.String())
		if err != nil {
			fmt.Printf("Issue with converting to an int: %q\n", err)
			return
		}
		result = result + numsAsInt
	}
	fmt.Printf("\nFinal Result:\n%d", result)
}

func Trebuchet() {
	fmt.Println("Input text for Trebuchet")
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

	problem(lines...)
}
