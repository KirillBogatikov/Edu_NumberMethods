package main

import (
	"NumberMethods/benchmark"
	"NumberMethods/div2"
	"NumberMethods/equation"
	_ "NumberMethods/seidel"
	"fmt"
	"github.com/shopspring/decimal"
)

const (
	Verbose = false
)

func main() {
	//runBenchmark()
	eq := equation.NewMemberEquation(
		equation.NewMemberFromInt(-2, 3),
		equation.NewMemberFromInt(-6, 2),
		equation.NewMemberFromInt(7, 1),
		equation.NewMemberFromFloat(-5, 0))
	eq.Verbose(Verbose)

	fmt.Println(div2.Solve(Verbose, decimal.NewFromFloat(0.00001), eq))
	fmt.Println("Log:\n" + eq.Log().String())
}

func runBenchmark() {
	var benchmarkCode int64

	fmt.Print("Choose benchmark: \n1 - Gauss Equation System solving\n2 - Matrix reversing\n3 - Seidel Equation System solving. Enter code: ")
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
	case 3:
		fmt.Println("Seidel benchmark started")
		prov = new(benchmark.SeidelProvider)
		break
	}

	benchmark.StartBenchmark(prov)
}
