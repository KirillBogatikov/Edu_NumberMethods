package main

import (
	"NumberMethods/benchmark"
	"NumberMethods/gauss"
	"NumberMethods/matrix"
	_ "NumberMethods/seidel"
	"fmt"
	"log"
)

func main() {
	//runBenchmark()

	system := gauss.NewSystem(matrix.NewMatrix(4, 3, [][]float64{
		{5.4, -2.3, 3.4, -3.5},
		{4.2, 1.7, -2.3, 2.7},
		{3.4, 2.4, 7.4, 1.9},
	}))

	sol := system.Solve()
	log.Println(sol)
	log.Println(system.CheckSolution(sol))
}

func runBenchmark() {
	var benchmarkCode int64

	fmt.Print("Choose benchmark: \n1 - Gauss Equation System solving\n2 - Matrix reversing\n 3 - Seidel Equation System solving. Enter code: ")
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
