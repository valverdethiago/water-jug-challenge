package main

import (
	"fmt"
	"github.com/valverdethiago/water-jug-challenge/code/backend/service"
)

func printSolution(steps []service.State) {
	if steps == nil {
		return
	}
	for _, step := range steps {
		fmt.Printf("%d, %d - %s \n", step.BucketX, step.BucketY, step.Explanation)
	}
}

func main() {
	waterJugService := service.NewWaterJugServiceImpl()
	X, Y, Z := 8, 6, 5 // Example: Jug capacities and target

	solution, found, err := waterJugService.SolveWaterJugProblem(X, Y, Z)
	if err != nil {
		fmt.Printf("Error %s \n", err.Error())
	}
	if found {
		fmt.Println("Solution Path:")
		printSolution(solution)
	} else {
		fmt.Println("No Solution")
	}
}
