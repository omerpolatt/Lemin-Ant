package bfs

func Bfs(graph map[string]map[string]bool, start string, end string) [][]string {
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
