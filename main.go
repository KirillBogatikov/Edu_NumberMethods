package main

import (
	"NumberMethods/benchmark"
	"fmt"
)

func main() {
	var benchmarkCode int64

	fmt.Print("Choose benchmark: \n1 - Gauss Equation System solving\n2 - Matrix reversing. Enter code: ")
	_, e := fmt.Scanln(&benchmarkCode)
	if e != nil {
		panic(e)
	}

	var prov benchmark.Provider
	switch benchmarkCode {
	case 1:
		fmt.Println("Gauss benchmark started")
		prov = new(benchmark.GaussProvider)
		break
	case 2:
		fmt.Println("Matrix benchmark started")
		prov = new(benchmark.MatrixProvider)
		break
	}

	benchmark.StartBenchmark(prov)
}
