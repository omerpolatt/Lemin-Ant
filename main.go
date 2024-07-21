package main

import (
	bfs "ant/allPaths"
	events "ant/fileEvent"
	filtre "ant/filteredPaths"
	"ant/simulate"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	result, err := events.ParseInputFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Ant Count: %d\n", result.AntCount)

	graph := events.BuildGraph(result.Rooms, result.Connections)

	startName := result.StartRoom.Name
	endName := result.EndRoom.Name

	allPaths := bfs.Bfs(graph, startName, endName)
	fmt.Println("ALL PATHS :", allPaths)

	filtrepaths := filtre.FilterPaths(allPaths)
	fmt.Println("FİLTERED PATHS :", filtrepaths)

	movements := simulate.SimulateAntMovement(filtrepaths, result.AntCount, startName, endName, filtrepaths[0])
	simulate.PrintOutput(movements)
}
