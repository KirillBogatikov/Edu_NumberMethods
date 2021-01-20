package main

import (
	"NumberMethods/matrix"
	"fmt"
	"log"
	"time"
)

type Benchmark struct {
	totalTime   time.Duration
	averageTime float64
}

func main() {
	log.Printf("Total time: %d ms", singleBenchmark())
}

var (
	matricies = map[int][][]float64 {
		0: {{1, 2, 3, 4}, {2, 3, 4, 1}, {3, 4, 1, 2}, {4, 1, 2, 3}},
		1: {{1, 4, 1, 3}, {0, -1, 3, -1}, {3, 1, 0, 2}, {1, -2, 5, 1}},
		2: {{1, 1, 1, 1}, {1, 4, 2, 3}, {1, 10, 3, 6}, {6, 10, 1, 4}},
		3: {{1, 2, 3, -2}, {2, -1, -2, -3}, {3, 2, -1, 2}, {2, -3, 2, 1}},
		4: {{-2, 2, 1, 0}, {1, -3, 3, 7}, {2, -1, 2, -3}, {-5, 4, -1, 2}},
		5: {{1, -1, -1, 1}, {-1, 2, 2, 0}, {0, -1, 1, 4}, {1, 1, -1, -1.5}},
		6: {{3, -2, 2, 0}, {2, 1, 1, -2}, {3, -1, 2, 1}, {1, 2, -1, -1}},
		7: {{5, -4, 0, 2}, { -1, 1, 1, -1}, {2, 3, 1, -6}, {1, 0, 2, -1}},
		8: {{4, -1, 0, 1}, {3, 2, -1, 2}, {0, 2, 2, 1}, {-1, 1, -3, -1}},
		9: {{1, 2, 0, 1}, {-1, -3, 3, -1}, {0, 4, -10, 2}, {1, -1, 2, -1}},
		10: {{1, 4, -3, 0}, {0, 4, 1, 2}, {-1, 2, 4, 1}, {1, 0, -1, 5}},
		11: {{2, -1, 1, 2}, {1, 2, -1, 1}, {3, 0, -1, -3}, {1, -1, 1, 3}},
		12: {{2, 1, 2, 0}, {-1, -3, 3, -1}, {1, 3, -8, 1}, {1, -1, 2, -1}},
		13: {{2, 3, 0, 1}, {-1, 1, 3, 0}, {0, 2, -1, 1}, {3, -1, 1, -2}},
		14: {{0, 3, 3, 2}, {-2, 2, 2, 1}, {0, 2, 2, 0}, {-1, 3, 3, 3}},
		15: {{1, 0, 3, 24}, {0, 1, 5, 6}, {-3, 4, 10, 6}, {0, -6, 0, -6}},
	}

	Variant = 10
)

func singleBenchmark() time.Duration {
	t := time.Now()
	m := matrix.NewMatrix(4, 4, matricies[Variant])
	r := m.Reverse()

	fmt.Printf("Reversed matrix: %s\n", "\n"+r.ToString())
	fmt.Printf("Check matrix: %s\n", "\n"+m.MulMatrix(r).ToString())

	return time.Now().Sub(t) / time.Millisecond
}

const (
	OperationCount = 10_000
)

func multiBenchmark(c chan Benchmark) {
	t := time.Now()
	var total time.Duration = 0

	for i := 0; i < OperationCount; i++ {
		total += singleBenchmark()

		if i%1_000 == 0 {
			fmt.Printf("%v %", float64(i)/float64(OperationCount))
		}
	}

	c <- Benchmark{time.Now().Sub(t) / time.Millisecond, float64(total) / float64(OperationCount)}
}
