package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name        string
	Coordinates []int
}

type Connection struct {
	From string
	To   string
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File Not Read", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var antCount int
	var start Room
	var end Room
	var rooms []Room
	var connections []Connection

	if scanner.Scan() { // Read the first line to get the ant count
		antCount, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Ant Count Not Read", err)
			return
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##comment") || strings.HasPrefix(line, "#comment") || strings.HasPrefix(line, "#another comment") {
			continue
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			connections = append(connections, Connection{From: parts[0], To: parts[1]})
		} else {
			if strings.HasPrefix(line, "##start") {
				scanner.Scan()
				start = parseRoom(scanner.Text())
				rooms = append(rooms, start)
			} else if strings.HasPrefix(line, "##end") {
				scanner.Scan()
				end = parseRoom(scanner.Text())
				rooms = append(rooms, end)
			} else {
				rooms = append(rooms, parseRoom(line))
			}
		}
	}

	fmt.Printf("Ant Count: %d\n", antCount)
	graph := buildGraph(rooms, connections)

	startName := start.Name
	endName := end.Name

	allPaths := bfs(graph, startName, endName)
	fmt.Println(allPaths)

	truePath := filterPaths(allPaths)
	fmt.Println("Filtered Paths:", truePath)

	movements := simulateAntMovement(truePath, antCount, startName, endName, truePath[0])
	printOutput(movements)
}

func parseRoom(line string) Room {
	parts := strings.Fields(line)
	name := parts[0]
	var coordinates []int
	for _, part := range parts[1:] {
		coordinate, _ := strconv.Atoi(part)
		coordinates = append(coordinates, coordinate)
	}
	return Room{Name: name, Coordinates: coordinates}
}

func buildGraph(rooms []Room, connections []Connection) map[string]map[string]bool {
	graph := make(map[string]map[string]bool)
	for _, room := range rooms {
		graph[room.Name] = make(map[string]bool)
	}
	for _, conn := range connections {
		graph[conn.From][conn.To] = true
		graph[conn.To][conn.From] = true
	}
	return graph
}

func bfs(graph map[string]map[string]bool, start string, end string) [][]string {
	var allPaths [][]string      // Array to hold all alternative paths
	queue := [][]string{{start}} // The queue always starts with the starting node

	for len(queue) > 0 { // Continue as long as there are elements in the queue
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1] // The node to be processed is the last element in the queue

		if node == end { // If the node being checked is equal to the end, add the path to allPaths
			allPaths = append(allPaths, path)
			continue

		} else { // If the node being checked is not equal to the end, continue from here

			for neighbor := range graph[node] { // Create a graph with the nodes connected to the last node

				// Check if the neighboring node is already in the path
				if !isNeighborInPath(neighbor, path) { // If the neighbors are not in the queue
					// Create a new path by copying the existing path and adding the neighboring node
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, neighbor)

					// Add the new path to the queue
					queue = append(queue, newPath)
				}
			}
		}

	}

	return allPaths
}

func isNeighborInPath(neighbor string, path []string) bool {
	for _, node := range path {
		if node == neighbor {
			return true
		}
	}
	return false
}

func filterPaths(paths [][]string) [][]string {
	var maxPaths [][]string
	var currentPaths [][]string
	usedNodes := make(map[string]bool)

	var backtrack func(int)
	backtrack = func(start int) {
		if len(currentPaths) > len(maxPaths) {
			maxPaths = make([][]string, len(currentPaths))
			copy(maxPaths, currentPaths)
		}

		for i := start; i < len(paths); i++ {
			path := paths[i]
			keepPath := true

			for _, node := range path[1 : len(path)-1] {
				if usedNodes[node] {
					keepPath = false
					break
				}
			}

			if keepPath {
				currentPaths = append(currentPaths, path)
				for _, node := range path[1 : len(path)-1] {
					usedNodes[node] = true
				}

				backtrack(i + 1)

				// Backtrack
				currentPaths = currentPaths[:len(currentPaths)-1]
				for _, node := range path[1 : len(path)-1] {
					delete(usedNodes, node)
				}
			}
		}
	}

	backtrack(0)
	return maxPaths
}

func simulateAntMovement(paths [][]string, antCount int, start, end string, shortestPath []string) []string {
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

func printOutput(movements []string) {
	for _, movement := range movements {
		fmt.Println(movement)
	}
}
