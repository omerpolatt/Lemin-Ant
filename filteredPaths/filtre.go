package filtre

// FilterPaths fonksiyonu, verilen yollar arasından çakışmayan en uzun yollar kümesini bulur ve döndürür.
func FilterPaths(paths [][]string) [][]string {
	var maxPaths [][]string            // Şimdiye kadar bulunan en iyi çakışmayan yollar kümesini tutar.
	var currentPaths [][]string        // Geçici olarak mevcut denenen yolları tutar.
	usedNodes := make(map[string]bool) // Mevcut yollarda kullanılan düğümleri işaretlemek için kullanılır.

	// Tüm kombinasyonları keşfetmek için geri dönüş (backtrack) fonksiyonu.
	var backtrack func(int)
	backtrack = func(start int) {
		// Eğer mevcut yollar kümesi, şimdiye kadar bulunandan daha uzunsa, maxPaths güncellenir.
		if len(currentPaths) > len(maxPaths) {
			maxPaths = make([][]string, len(currentPaths))
			copy(maxPaths, currentPaths)
		}

		// 'start' indeksinden başlayarak tüm yolları dener.
		for i := start; i < len(paths); i++ {
			path := paths[i]
			keepPath := true // Mevcut yolu eklemek için çakışma olup olmadığını kontrol eder.

			// Yolun orta düğümleri zaten kullanılmış mı diye kontrol eder.
			for _, node := range path[1 : len(path)-1] {
				if usedNodes[node] {
					keepPath = false // Eğer düğüm kullanılmışsa, flag false olarak ayarlanır ve döngüden çıkılır.
					break
				}
			}

			// Eğer yol çakışma içermiyorsa, mevcut yollara eklenir.
			if keepPath {
				currentPaths = append(currentPaths, path)
				// Bu yolun düğümlerini kullanıldı olarak işaretler.
				for _, node := range path[1 : len(path)-1] {
					usedNodes[node] = true
				}

				// Bir sonraki başlangıç indeksi ile geri dönüşü çağırır.
				backtrack(i + 1)

				// Geri dönüş: Eklenen son yolu kaldırır ve düğümlerinin işaretlerini siler.
				currentPaths = currentPaths[:len(currentPaths)-1]
				for _, node := range path[1 : len(path)-1] {
					delete(usedNodes, node)
				}
			}
		}
	}

	// Rekürsif fonksiyonu 0 indeksinden başlatır.
	backtrack(0)
	// Bulunan en iyi çakışmayan yollar kümesini döndürür.
	return maxPaths
}
