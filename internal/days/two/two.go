package two

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Number int
	Blue   int
	Green  int
	Red    int
}

func ParseGame(s string) (Game, error) {
	// parse game number
	gameAndTail := strings.Split(s, ":")
	n, err := parseGameNumber(gameAndTail[0])
	if err != nil {
		return Game{}, err
	}

	// parse colors
	blue, green, red := parseColors(gameAndTail[1])

	return Game{
		Number: n,
		Blue:   blue,
		Green:  green,
		Red:    red,
	}, nil
}

func parseGameNumber(s string) (int, error) {
	re := regexp.MustCompile(`Game (\d+)`)
	match := re.FindStringSubmatch(s)
	number, err := strconv.ParseInt(match[1], 10, 32)

	return int(number), err
}

func parseColors(s string) (int, int, int) {
	var blue int
	var green int
	var red int

	for _, roll := range strings.Split(s, ";") {
		for _, c := range []struct {
			Re    *regexp.Regexp
			Color *int
		}{
			{regexp.MustCompile(`(\d+) blue`), &blue},
			{regexp.MustCompile(`(\d+) green`), &green},
			{regexp.MustCompile(`(\d+) red`), &red},
		} {
			swapIfLarger(c.Re, c.Color, roll)
		}
	}

	return blue, green, red
}

func swapIfLarger(re *regexp.Regexp, old *int, s string) {
	if new, ok := parseColor(re, s); ok {
		if *old < new {
			*old = new
		}
	}
}

func parseColor(re *regexp.Regexp, s string) (int, bool) {
	if match := re.FindStringSubmatch(s); len(match) >= 2 {
		new, err := strconv.ParseInt(match[1], 10, 32)
		if err != nil {
			return 0, false
		}
		return int(new), true
	}

	return 0, false
}

func Solution(input *bufio.Scanner, blue, green, red int) (int, int) {
	var winning_sum int
	var power_sum int

	for input.Scan() {
		in := input.Text()
		game, err := ParseGame(in)
		if err != nil {
			panic(err)
		}

		if blue >= game.Blue && green >= game.Green && red >= game.Red {
			winning_sum = winning_sum + game.Number
		}

		power_sum = power_sum + (game.Blue * game.Green * game.Red)
	}

	return winning_sum, power_sum
}
