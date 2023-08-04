package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RPS int
type WLD int

const (
	Rock RPS = iota
	Paper
	Scissors
)
const (
	Win WLD = iota
	Lose
	Draw
)

type Round struct {
	mine     RPS
	opponent RPS
}

type OutcomePair struct {
	opponent RPS
	outcome  WLD
}

var xyzWinLoseMap = map[string]WLD{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

var xyzRPSMap = map[string]RPS{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var abcRPSMap = map[string]RPS{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var choiceScore = map[RPS]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var compareMap = map[Round]int{
	{Rock, Rock}:         3,
	{Rock, Paper}:        0,
	{Rock, Scissors}:     6,
	{Paper, Rock}:        6,
	{Paper, Paper}:       3,
	{Paper, Scissors}:    0,
	{Scissors, Rock}:     0,
	{Scissors, Paper}:    6,
	{Scissors, Scissors}: 3,
}

var outcomeMap = map[OutcomePair]RPS{
	{Rock, Win}:      Paper,
	{Rock, Lose}:     Scissors,
	{Rock, Draw}:     Rock,
	{Paper, Win}:     Scissors,
	{Paper, Lose}:    Rock,
	{Paper, Draw}:    Paper,
	{Scissors, Win}:  Rock,
	{Scissors, Lose}: Paper,
	{Scissors, Draw}: Scissors,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculateRoundScore(round Round) int {
	score, exists := compareMap[round]
	if !exists {
		panic("Invalid round")
	}

	return choiceScore[round.mine] + score
}

func main() {
	file, err := os.Open("2022/day2/input.txt")
	check(err)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	score := 0
	altScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		opponent, mine, found := strings.Cut(line, " ")
		if !found {
			panic("Invalid input")
		}

		opponentPlay := abcRPSMap[opponent]
		round := Round{
			mine:     xyzRPSMap[mine],
			opponent: opponentPlay,
		}

		score += calculateRoundScore(round)

		// Part 2
		myOutcome := xyzWinLoseMap[mine]
		altScore += calculateRoundScore(Round{
			mine:     outcomeMap[OutcomePair{opponentPlay, myOutcome}],
			opponent: opponentPlay,
		})
	}

	fmt.Println("Part 1:", score)
	fmt.Println("Part 2:", altScore)

}
