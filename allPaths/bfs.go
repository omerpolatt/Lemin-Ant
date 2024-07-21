package bfs

func Bfs(graph map[string]map[string]bool, start string, end string) [][]string {
	var allPaths [][]string // Olası yolların tamamını bulunduran dizi ( Sequence with all possible pathways )
	queue := [][]string{{start}}

	for len(queue) > 0 { // kuyrukta eleman olduğu sürece yol bulmaya devam edeceğiz ( we will continue as long as there are nodes in the queue )
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		if node == end { // son düğüm end noktasına eşit ise o yolu dizi içine ekleriz  ( if the last node is equal to the end point, we add that path to the array )
			allPaths = append(allPaths, path)
			continue
		}

		for neighbor := range graph[node] { // son düğüm end noktasına eşit değil ise komşuluk listesinden mevcut düğümün ziyaret edilmemiş her bir komşusuna gidilerek yol seçeneklerini else ederiz
			if !isNeighborInPath(neighbor, path) { // ( if the last node is not equal to the end point, we go to each unvisited neighbor of the current node from the neighborhood list and choose the path options else )
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return allPaths
}

func isNeighborInPath(neighbor string, path []string) bool { // komşuluk listemizi oluşturduğumuz fonksiyonumuz ( the function where we create our neighborhood list )
	for _, node := range path {
		if node == neighbor {
			return true
		}
	}
	return false
}
