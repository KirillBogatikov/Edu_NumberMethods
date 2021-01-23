package main

import (
	"NumberMethods/benchmark"
)

func main() {
	//prov := new(benchmark.GaussProvider)
	prov := new(benchmark.MatrixProvider)
	benchmark.StartBenchmark(prov)
}
