package cmd

import (
	"fmt"
	"strings"

	"github.com/lanebryer/coinsim/game"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short:   "A simple coin flip simulator with a variable number of iterations",
	PreRunE: validateInputs,
	Run:     run,
}

var runs uint64
var p1Name string
var p2Name string
var p1Sequence []string
var p2Sequence []string

func init() {
	rootCmd.Flags().Uint64Var(&runs, "runs", 10_000, "Sets the number of total games to play")
	rootCmd.Flags().StringVar(&p1Name, "p1name", "John", "The first player's name")
	rootCmd.Flags().StringVar(&p2Name, "p2name", "Adam", "The second player's name")
	rootCmd.Flags().StringSliceVar(&p1Sequence, "p1sequence", []string{"heads", "tails", "heads"}, "An unspaced, comma-delimited list of the winning sequence for player 1")
	rootCmd.Flags().StringSliceVar(&p2Sequence, "p2sequence", []string{"heads", "tails", "tails"}, "An unspaced, comma-delimited list of the winning sequence for player 2")
}

func Execute() error {
	return rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	gameParameters := &game.GameParameters{
		P1Name:     p1Name,
		P2Name:     p2Name,
		P1Sequence: p1Sequence,
		P2Sequence: p2Sequence,
		Runs:       runs,
	}

	game.Play(gameParameters)
}

func validateInputs(_ *cobra.Command, _ []string) error {
	if len(p1Sequence) != len(p2Sequence) {
		return fmt.Errorf("the player sequences must be of equal length - play fair")
	}

	if err := validateSequence(p1Sequence); err != nil {
		return err
	}

	if err := validateSequence(p2Sequence); err != nil {
		return err
	}

	return nil
}

func validateSequence(seq []string) error {
	for i, v := range seq {
		seq[i] = strings.ToLower(v)
		if seq[i] != "heads" && seq[i] != "tails" {
			return fmt.Errorf("coin sequences must only consist of heads or tails")
		}
	}

	return nil
}
