package benchmark

import (
	"NumberMethods/utils"
	"fmt"
	"log"
	"time"
)

type Benchmark struct {
	totalTime   int64
	averageTime float64
}

type Provider interface {
	Operator() func()
	Count() int64
}

func singleBenchmark(operator func()) time.Duration {
	t := time.Now()
	operator()
	return time.Now().Sub(t) / time.Millisecond
}

func multiBenchmark(c chan Benchmark, count int64, operator func()) {
	t := time.Now()
	var total time.Duration = 0
	step := int(count / 100)

	for i := 0; i < int(count); i++ {
		total += singleBenchmark(operator)

		if i%step == 0 {
			utils.ClearConsole()
			fmt.Printf("Benchmark progress: %d%% (%d / %d)\n", int8(float64(i)/float64(count)*100), i, count)
		}
	}

	c <- Benchmark{time.Now().Sub(t).Milliseconds(), float64(total) / float64(count)}
}

func StartBenchmark(provider Provider) {
	c := make(chan Benchmark)
	go multiBenchmark(c, provider.Count(), provider.Operator())

	bench := <-c
	log.Printf("Total time: %v ms, %v ms", bench.totalTime, bench.averageTime)
	log.Println("Benchmark ok. Press any key to exit")
	_, _ = fmt.Scanln()
}
