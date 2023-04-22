package _go

func checkMagazine(magazine []string, note []string) string {
	magazineWordsMap := make(map[string]int)
	for _, word := range magazine {
		magazineWordsMap[word] = magazineWordsMap[word] + 1
	}
	noteWordsMap := make(map[string]int)
	for _, word := range note {
		noteWordsMap[word] = noteWordsMap[word] + 1
	}
	for word, times := range noteWordsMap {
		if magazineWordsMap[word] < times {
			return "No"
		}
	}
	return "Yes"
}