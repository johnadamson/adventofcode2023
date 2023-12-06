package one

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var numberWords = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func ParseDigits(in string) uint8 {
	var numbers []string

	for i, c := range in {
		if unicode.IsDigit(c) {
			numbers = append(numbers, string(c))
			continue
		}
		for k, v := range numberWords {
			if strings.Index(in[i:], k) == 0 {
				numbers = append(numbers, fmt.Sprint(v))
			}
		}
	}

	result, err := strconv.ParseUint(fmt.Sprintf(`%s%s`, numbers[0], numbers[len(numbers)-1]), 10, 8)
	if err != nil {
		panic(err)
	}

	return uint8(result)
}

func Solution(input *bufio.Scanner) uint {
	var sum uint

	for input.Scan() {
		sum = sum + uint(ParseDigits(input.Text()))

	}

	return sum
}
