package main

import (
	"NumberMethods/benchmark"
	"NumberMethods/gauss"
	"NumberMethods/matrix"
	"NumberMethods/seidel"
	_ "NumberMethods/seidel"
	"NumberMethods/utils"
	"fmt"
)

func main() {
	//runBenchmark()

	gauss_system := gauss.NewSystem(matrix.NewMatrix(4, 3, [][]float64{
		{1.06, -0.28, 0.84, 0.57},
		{0.43, 0.62, -0.35, 0.66},
		{0.37, -0.75, -0.64, -0.38},
	}))

	gauss_sol := gauss_system.Solve()
	utils.PrintDecimalArray(gauss_sol)

	equals, accuracy := gauss_system.CheckSolution(gauss_sol)
	fmt.Printf("Solution is correct: %v\n", equals)
	utils.PrintDecimalArray(accuracy)
	fmt.Println()

	system := seidel.NewSystem(matrix.NewMatrix(4, 3, [][]float64{
		{1.06, -0.28, 0.84, 0.57},
		{0.43, 0.62, -0.35, 0.66},
		{0.37, -0.75, -0.64, -0.38},
	}), 0.001, 0.001)

	sol, count := system.Solve()
	fmt.Printf("Iterations: %d\n", count)
	utils.PrintDecimalArray(sol)

	equals, accuracy = gauss_system.CheckSolution(gauss_sol)
	fmt.Printf("Solution is correct: %v\n", equals)
	utils.PrintDecimalArray(accuracy)
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
