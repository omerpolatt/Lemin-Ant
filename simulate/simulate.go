package simulate

import (
	"fmt"
	"strings"
)

func SimulateAntMovement(paths [][]string, antCount int, start, end string, shortestPath []string) []string {
	var movements []string
	antPositions := make(map[int]int)
	antReached := make(map[int]bool)
	antPaths := make(map[int][]string)
	activeAnts := antCount

	for i := 1; i <= antCount; i++ {
		if i == antCount {
			antPaths[i] = shortestPath
		} else {
			antPaths[i] = paths[(i-1)%len(paths)]
		}
		antPositions[i] = 0
		antReached[i] = false
	}

	round := 0
	for activeAnts > 0 {
		round++
		var roundMovements []string
		tunnelUsage := make(map[string]bool)

		for i := 1; i <= antCount; i++ {
			if antReached[i] {
				continue
			}

			currentRoom := antPaths[i][antPositions[i]]
			nextRoom := antPaths[i][antPositions[i]+1]
			tunnel := fmt.Sprintf("%s-%s", currentRoom, nextRoom)
			reverseTunnel := fmt.Sprintf("%s-%s", nextRoom, currentRoom)

			if !tunnelUsage[tunnel] && !tunnelUsage[reverseTunnel] {
				roundMovements = append(roundMovements, fmt.Sprintf("L%d-%s", i, nextRoom))
				tunnelUsage[tunnel] = true
				tunnelUsage[reverseTunnel] = true
				antPositions[i]++
				if nextRoom == end {
					antReached[i] = true
					activeAnts--
				}
			}
		}

		if len(roundMovements) > 0 {
			movements = append(movements, strings.Join(roundMovements, " "))
		} else {
			break
		}
	}
	return movements
}

func PrintOutput(movements []string) {
	fmt.Println("Movement Details:")
	for round, movement := range movements {
		fmt.Printf("Tour %d: %s\n", round+1, movement)
	}
	fmt.Printf("Completed in a total of %d rounds.\n", len(movements))
}
