import "slices"

func groupAnagrams(strs []string) [][]string {

	m := make(map[string]int, len(strs))
	res := make([][]string, 0, len(strs))

	for _, s := range strs {
		runes := []rune(s)
		slices.Sort(runes)
		key := string(runes)

		if _, exists := m[key]; !exists {
			m[key] = len(res)
			res = append(res, []string{})
		}


		res[m[key]] = append(res[m[key]], s)

	}

	return res

}
