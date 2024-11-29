package others

func MinimalOperations(words []string) []int32 {
	ops := make([]int32, len(words))

	for index, word := range words {
		ops[index] = minimalOperation(word)
	}
	return ops
}

func minimalOperation(word string) int32 {
	var count int32
	bWord := []byte(word)
	bmap := make(map[byte]bool)

	for i := 0; i < len(bWord); i++ {
		if bmap[bWord[i]] {
			continue
		}
		if i-1 >= 0 {
			if bWord[i] == bWord[i-1] {
				bmap[bWord[i]] = true
				count++
			} else {
				bmap = make(map[byte]bool)
			}
		}
	}

	return count
}

func minimalOperation2(word string) int32 {
	bWord := []byte(word)
	count := 0
	totalCount := 0

	for i := 0; i < len(bWord); i++ {
		if i < len(bWord)-1 {
			if bWord[i] == bWord[i+1] {
				count++
			} else {
				if count >= 1 {
					totalCount++
				}
				count = 0
			}
		} else {
			if count >= 1 {
				totalCount++
			}
		}
	}

	return int32(totalCount)
}
