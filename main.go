package main

import (
	"NumberMethods/gauss"
	"NumberMethods/matrix"
	"NumberMethods/utils"
	"fmt"
	"log"
	"time"
)

type Benchmark struct {
	totalTime   int64
	averageTime float64
}

func main() {
	/* log.Printf("Total time: %d ms", singleBenchmark(func() {
		m := matrix.NewMatrix(4, 4, matricies[Variant])
		r := m.Reverse()

		fmt.Printf("Reversed matrix: %s\n", "\n"+r.ToString())
		fmt.Printf("Check matrix: %s\n", "\n"+m.MulMatrix(r).ToString())
	})) */
	c := make(chan Benchmark)
	go multiBenchmark(c, func() {
		_ = gauss.NewEquationSystem(matrix.NewMatrix(4, 3, [][]float64{
			{5.4, -2.3, 3.4, -3.5},
			{4.2, 1.7, -2.3, 2.7},
			{3.4, 2.4, 7.4, 1.9},
		}))

		//fmt.Printf("Result: %s\n", e.Solve())
	})

	bench := <-c
	log.Printf("Total time: %v ms, %v ms", bench.totalTime, bench.averageTime)
	log.Println("Benchmark ok. Press any key to exit")
	_, _ = fmt.Scanln()
}

var (
	matricies = map[int][][]float64{
		0:  {{1, 2, 3, 4}, {2, 3, 4, 1}, {3, 4, 1, 2}, {4, 1, 2, 3}},
		1:  {{1, 4, 1, 3}, {0, -1, 3, -1}, {3, 1, 0, 2}, {1, -2, 5, 1}},
		2:  {{1, 1, 1, 1}, {1, 4, 2, 3}, {1, 10, 3, 6}, {6, 10, 1, 4}},
		3:  {{1, 2, 3, -2}, {2, -1, -2, -3}, {3, 2, -1, 2}, {2, -3, 2, 1}},
		4:  {{-2, 2, 1, 0}, {1, -3, 3, 7}, {2, -1, 2, -3}, {-5, 4, -1, 2}},
		5:  {{1, -1, -1, 1}, {-1, 2, 2, 0}, {0, -1, 1, 4}, {1, 1, -1, -1.5}},
		6:  {{3, -2, 2, 0}, {2, 1, 1, -2}, {3, -1, 2, 1}, {1, 2, -1, -1}},
		7:  {{5, -4, 0, 2}, {-1, 1, 1, -1}, {2, 3, 1, -6}, {1, 0, 2, -1}},
		8:  {{4, -1, 0, 1}, {3, 2, -1, 2}, {0, 2, 2, 1}, {-1, 1, -3, -1}},
		9:  {{1, 2, 0, 1}, {-1, -3, 3, -1}, {0, 4, -10, 2}, {1, -1, 2, -1}},
		10: {{1, 4, -3, 0}, {0, 4, 1, 2}, {-1, 2, 4, 1}, {1, 0, -1, 5}},
		11: {{2, -1, 1, 2}, {1, 2, -1, 1}, {3, 0, -1, -3}, {1, -1, 1, 3}},
		12: {{2, 1, 2, 0}, {-1, -3, 3, -1}, {1, 3, -8, 1}, {1, -1, 2, -1}},
		13: {{2, 3, 0, 1}, {-1, 1, 3, 0}, {0, 2, -1, 1}, {3, -1, 1, -2}},
		14: {{0, 3, 3, 2}, {-2, 2, 2, 1}, {0, 2, 2, 0}, {-1, 3, 3, 3}},
		15: {{1, 0, 3, 24}, {0, 1, 5, 6}, {-3, 4, 10, 6}, {0, -6, 0, -6}},
	}

	Variant = 10
)

func singleBenchmark(operator func()) time.Duration {
	t := time.Now()
	operator()
	return time.Now().Sub(t) / time.Millisecond
}

const (
	OperationCount = 1_000_000
)

func multiBenchmark(c chan Benchmark, operator func()) {
	t := time.Now()
	var total time.Duration = 0
	const step = OperationCount / 100

	for i := 0; i < OperationCount; i++ {
		total += singleBenchmark(operator)

		if i%step == 0 {
			utils.ClearConsole()
			fmt.Printf("Benchmark progress: %d%% (%d / %d)\n", int8(float64(i)/float64(OperationCount)*100), i, OperationCount)
		}
	}

	c <- Benchmark{time.Now().Sub(t).Milliseconds(), float64(total) / float64(OperationCount)}
}
