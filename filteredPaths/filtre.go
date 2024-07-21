package filtre

func FilterPaths(paths [][]string) [][]string {
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
