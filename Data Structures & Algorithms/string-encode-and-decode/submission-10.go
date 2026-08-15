

type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	
	sizes := make([]string, 0, len(strs))
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}

	return strings.Join(sizes, ",") + ",#" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	sizes := make([]int,0, len(encoded) + 1)
	
	size := ""
	for i, r := range encoded {
		
		if r == '#' {
			encoded = encoded[i + 1:]
			break
		}

		if r == ',' {
			wordSize, _ := strconv.Atoi(size)
			sizes = append(sizes, wordSize)
			size = ""
			continue
		}

		size += string(r)	
	}

	
	strs := make([]string, 0, len(sizes))
	j := 0 
	for i := range sizes {
		length := sizes[i]
		strs = append(strs, encoded[j:j+length])
		j += sizes[i]

	}



	return strs

}
