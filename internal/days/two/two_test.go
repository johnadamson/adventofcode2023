package two_test

import (
	"adventofcode/internal/days/two"
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestParseLine(t *testing.T) {
	t.Run("parse into game number", func(t *testing.T) {
		// given
		line := "Game 1: 4 red, 5 blue, 9 green; 7 green, 7 blue, 3 red; 16 red, 7 blue, 3 green; 11 green, 11 blue, 6 red; 12 red, 14 blue"
		var expected int = 1

		// when
		result, _ := two.ParseGame(line)

		// then
		if expected != result.Number {
			t.Errorf("expected result.Number = %d, result was %d", expected, result.Number)
		}
	})

	t.Run("returns the max", func(t *testing.T) {
		// given
		line := "Game 3: 8 blue, 5 green, 2 red; 5 blue, 5 green, 7 red; 7 blue, 1 green, 7 red; 8 green, 14 blue, 7 red; 8 green, 14 blue; 8 blue, 2 green, 8 red"

		// when
		result, _ := two.ParseGame(line)

		// then
		for _, r := range []struct {
			Color    string
			Expected int
			Result   int
		}{
			{Color: "blue", Expected: 14, Result: result.Blue},
			{Color: "green", Expected: 8, Result: result.Green},
			{Color: "red", Expected: 8, Result: result.Red},
		} {
			if r.Expected != r.Result {
				t.Errorf("expected %s %d, result was %d", r.Color, r.Expected, r.Result)
			}
		}
	})
}

func TestDayTwoRunner(t *testing.T) {
	// given
	blue := 14
	green := 13
	red := 12

	// in
	input, _ := os.Open("input")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	// out
	output, _ := os.Create("output")
	defer output.Close()
	writer := bufio.NewWriter(output)

	// when
	winning_sum, power_sum := two.Solution(scanner, blue, green, red)
	fmt.Fprintln(writer, winning_sum, power_sum)

	// cleanup
	writer.Flush()
}
