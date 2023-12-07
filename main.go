package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/joshmenden/aocplz/internal/fetch"
	"github.com/joshmenden/aocplz/internal/printit"
	"github.com/joshmenden/aocplz/internal/solve"
)

var (
	now          = time.Now()
	day          = now.Day() + 1
	year         = now.Year()
	requiredVars = []string{
		"AOCPLZ_SESSION_TOKEN",
		"AOCPLZ_ROOT_DIR",
	}
)

func main() {
	err := validateEnvVars()
	if err != nil {
		handleError(err)
	}

	fetchCmd := flag.NewFlagSet("fetch", flag.ExitOnError)
	dayPtr := fetchCmd.Int("day", day, "the aoc day you're solving")
	yearPtr := fetchCmd.Int("year", year, "the aoc year you're solving")
	waitPtr := fetchCmd.Bool("wait", false, "whether or not to wait for the puzzle to be ready")
	dontOpenPtr := fetchCmd.Bool("dont-open", false, "if this flag is included, the puzzle won't automatically open in your browser")

	if len(os.Args) < 2 {
		handleError(fmt.Errorf("expected 'fetch' or 'solve' command"))
	}

	switch os.Args[1] {
	case "fetch":
		fetchCmd.Parse(os.Args[2:])

		puzzleReady, diff := fetch.IsPuzzleReady(*dayPtr, *yearPtr)
		if !puzzleReady {
			msg := fmt.Sprintf("the puzzle for day %v won't be ready until %v from now", *dayPtr, diff)
			if *waitPtr {
				printit.Info(msg)
				err = fetch.WaitForPuzzle(dayPtr, yearPtr, dontOpenPtr)
				if err != nil {
					handleError(err)
				}
			} else {
				handleError(fmt.Errorf(msg))
			}
		} else {
			err := fetch.FetchDayInput(dayPtr, yearPtr, dontOpenPtr)
			if err != nil {
				handleError(err)
			}
		}

		printit.Success("Fetch Complete, Happy Coding!")
	case "solve":
		solve.SolveDayPuzzle()
	default:
		handleError(fmt.Errorf("expected 'fetch' or 'solve' command"))
	}
}

func handleError(err error) {
	printit.Error(err.Error())
	os.Exit(1)
}

func validateEnvVars() error {
	for _, varName := range requiredVars {
		if os.Getenv(varName) == "" {
			return fmt.Errorf("%s must be set", varName)
		}
	}

	return nil
}
