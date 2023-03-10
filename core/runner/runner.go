package runner

import (
	"fmt"

	"github.com/HungOnBlog/thorr/core/executors"
	"github.com/HungOnBlog/thorr/core/models"
)

type ThorRunner struct{}

func NewThorRunner() *ThorRunner {
	return &ThorRunner{}
}

func (r *ThorRunner) Run(spawn int, suits []models.TestSuit) error {
	fmt.Println("Running suits with spawn: ", spawn)
	fmt.Println("Number of suits: ", len(suits))
	switch spawn {
	case 1:
		return r.runSequential(suits)
	case -1:
		return r.runParallel(suits)
	default:
		return r.runSpecific(suits, spawn)
	}
}

func (r *ThorRunner) runSequential(suits []models.TestSuit) error {
	var textSuitExecutor = executors.NewTestSuitExecutor()
	for _, suit := range suits {
		fmt.Println("Running suit: ", suit.Name)
		err := textSuitExecutor.Execute(&suit)
		if err != nil {
			fmt.Println("❌ Error: ", err)
		}
	}

	return nil
}

// Run all suits in parallel
func (r *ThorRunner) runParallel(suits []models.TestSuit) error {
	err := make(chan error)
	for _, suit := range suits {
		go func(suit models.TestSuit) {
			var textSuitExecutor = executors.NewTestSuitExecutor()
			fmt.Println("Running suit: ", suit.Name)
			err <- textSuitExecutor.Execute(&suit)
		}(suit)
	}

	for range suits {
		if err := <-err; err != nil {
			fmt.Println("❌ Error: ", err)
		}
	}

	return nil
}

func (r *ThorRunner) runSpecific(suits []models.TestSuit, spawn int) error {
	suitsSpawn := r.planner(suits, spawn)
	err := make(chan error)
	for _, suits := range suitsSpawn {
		go func(suits []models.TestSuit) {
			var textSuitExecutor = executors.NewTestSuitExecutor()
			for _, suit := range suits {
				fmt.Println("Running suit: ", suit.Name)
				err <- textSuitExecutor.Execute(&suit)
			}
		}(suits)
	}

	for range suits {
		if err := <-err; err != nil {
			fmt.Println("❌ Error: ", err)
		}
	}

	return nil
}

// To plan the suits to run on specific number of goroutines
// Simple strategy: using Round Robin algorithm (https://en.wikipedia.org/wiki/Round-robin_scheduling)
func (r *ThorRunner) planner(suits []models.TestSuit, spawn int) [][]models.TestSuit {
	suitsSpawn := make([][]models.TestSuit, spawn)
	for i := 0; i < spawn; i++ {
		suitsSpawn[i] = make([]models.TestSuit, 0)
	}

	for i, suit := range suits {
		suitsSpawn[i%spawn] = append(suitsSpawn[i%spawn], suit)
	}

	return suitsSpawn
}
