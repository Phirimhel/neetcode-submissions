import "slices"

func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string, len(strs))

	for _, s := range strs {
		runes := []rune(s)
		slices.Sort(runes)
		key := string(runes)

		if _, exists := m[key]; !exists {
			m[key] = []string{}
		}


		m[key] = append(m[key], s)


	}

	sl := make([][]string, 0 , len(m))

	for _, strs := range m {
		sl = append(sl, strs)
	}

	return sl

}
