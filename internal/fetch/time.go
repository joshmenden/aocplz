package fetch

import (
	"fmt"
	"sync"
	"time"

	"github.com/joshmenden/aocplz/internal/printit"
)

func IsPuzzleReady(day, year int) (bool, time.Duration) {
	now := time.Now()
	var dayStr string
	if day < 10 {
		dayStr = fmt.Sprintf("0%v", day)
	} else {
		dayStr = fmt.Sprintf("%v", day)
	}

	aocDay, _ := time.Parse(time.RFC3339, fmt.Sprintf("%v-12-%sT00:00:00-05:00", year, dayStr))

	return now.After(aocDay), aocDay.Sub(now).Round(1 * time.Second)
}

func WaitForPuzzle(day, year *int) (err error) {
	waitTimeSeconds := 5

	printit.Info((fmt.Sprintf("waiting for puzzle, checking every %v seconds...", waitTimeSeconds)))

	var wg sync.WaitGroup
	wg.Add(1) // keep the function open until the puzzle is ready

	ticker := time.NewTicker(time.Duration(waitTimeSeconds) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				ready, diff := IsPuzzleReady(*day, *year)
				if ready {
					printit.Success("puzzle is ready! fetching now...")
					err = FetchDayInput(day, year)
					close(quit)
					wg.Done()
				} else {
					printit.Info(fmt.Sprintf("%v left to wait...", diff))
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	wg.Wait()

	return
}
