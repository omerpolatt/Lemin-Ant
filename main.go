package main

import (
	bfs "ant/allPaths"
	errorslemin "ant/errorlemin"
	events "ant/fileEvent"
	filtre "ant/filteredPaths"
	"ant/simulate"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")

	// MergeProcess fonksiyonunu çağırarak dosya içeriğindeki hataları kontrol et ( Check for errors in the file content by calling the MergeProcess function )
	if errorOutput := errorslemin.Mergeprocess(lines); errorOutput != "" {
		fmt.Println(errorOutput)
		return
	}

	// Dosya içeriği uygunsa, dosya ayrıştırma işlemini gerçekleştirir ( Performs file parsing if the file content is appropriate  )
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
