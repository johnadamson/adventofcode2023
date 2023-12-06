package one_test

import (
	"adventofcode/internal/days/one"
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestDayOneParseDigits(t *testing.T) {
	// given
	cases := []struct {
		Input  string
		Output uint8
	}{
		{Input: "1abc2", Output: 12},
		{Input: "pqr3stu8vwx", Output: 38},
		{Input: "a1b2c3d4e5f", Output: 15},
		{Input: "treb7uchet", Output: 77},
		{Input: "eightwo", Output: 82},
		{Input: "oneight", Output: 18},
		{Input: "two1nine", Output: 29},
		{Input: "eightwothree", Output: 83},
		{Input: "abcone2threexyz", Output: 13},
		{Input: "xtwone3four", Output: 24},
		{Input: "4nineeightseven2", Output: 42},
		{Input: "zoneight234", Output: 14},
		{Input: "7pqrstsixteen", Output: 76},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			// when
			result := one.ParseDigits(c.Input)

			// then
			if result != c.Output {
				t.Fatalf(`expected %d, result was %d`, c.Output, result)
			}
		})
	}
}

func TestDayOneRunner(t *testing.T) {
	// given
	// in
	input, _ := os.Open("input")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	// out
	output, _ := os.Create("output")
	defer output.Close()
	writer := bufio.NewWriter(output)

	// when
	fmt.Fprintln(writer, one.Solution(scanner))

	// cleanup
	writer.Flush()
}
