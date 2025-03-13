package game

import (
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type results struct {
	p1Name          string
	p2Name          string
	p1Win           uint
	p2Win           uint
	tie             uint
	p1WinPercentage float32
	p2WinPercentage float32
	tiePercentage   float32
}

type GameParameters struct {
	Runs       uint64
	P1Name     string
	P2Name     string
	P1Sequence []string
	P2Sequence []string
}

func Play(gp *GameParameters) {
	var p1Flips []string
	var p2Flips []string

	results := &results{
		p1Name: gp.P1Name,
		p2Name: gp.P2Name,
	}

	for range gp.Runs {
		for {
			p1Result := flipCoin()
			p2Result := flipCoin()
			p1Flips = addToResult(p1Result, p1Flips, len(gp.P1Sequence))
			p2Flips = addToResult(p2Result, p2Flips, len(gp.P2Sequence))
			winner, ok := winner(p1Flips, p2Flips, gp)

			if ok {
				switch winner {
				case "p1":
					results.p1Win++
				case "p2":
					results.p2Win++
				case "tie":
					results.tie++
				}

				clear(p1Flips)
				clear(p2Flips)
				break
			}
		}
	}

	results.p1WinPercentage = float32(results.p1Win) / float32(gp.Runs)
	results.p2WinPercentage = float32(results.p2Win) / float32(gp.Runs)
	results.tiePercentage = float32(results.tie) / float32(gp.Runs)

	printResults(results)
}

func flipCoin() string {
	if rand.IntN(2) == 0 {
		return "heads"
	}

	return "tails"
}

func addToResult(result string, flips []string, max int) []string {
	if len(flips) == max {
		flips = flips[1:]
	}

	return append(flips, result)
}

func winner(p1Flips []string, p2Flips []string, gp *GameParameters) (string, bool) {
	p1Win := equalSlices(p1Flips, gp.P1Sequence)
	p2Win := equalSlices(p2Flips, gp.P2Sequence)

	if p1Win && p2Win {
		return "tie", true
	}

	if p1Win {
		return "p1", true
	}

	if p2Win {
		return "p2", true
	}

	return "", false
}

func equalSlices(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}

func printResults(results *results) {
	totalWins := results.p1Win + results.p2Win + results.tie
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Win Percentage", "Wins"})
	table.SetFooter([]string{"", "Total", strconv.FormatUint(uint64(totalWins), 10)})

	data := [][]string{
		{results.p1Name, strconv.FormatFloat(float64(results.p1WinPercentage), 'f', 2, 64), strconv.FormatUint(uint64(results.p1Win), 10)},
		{results.p2Name, strconv.FormatFloat(float64(results.p2WinPercentage), 'f', 2, 64), strconv.FormatUint(uint64(results.p2Win), 10)},
		{"Tie", strconv.FormatFloat(float64(results.tiePercentage), 'f', 2, 64), strconv.FormatUint(uint64(results.tie), 10)},
	}

	table.AppendBulk(data)
	table.Render()
}
