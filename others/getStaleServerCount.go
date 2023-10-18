package others

func GetStaleServerCount(n int32, log_data [][]int32, query []int32, X int32) []int32 {
	staleServers := make([]int32, len(query))

	for k, q := range query {
		// get time interval
		intervalMin := q - X
		intervalMax := q
		interactMap := make(map[int32]bool, 0)

		for _, data := range log_data {
			if data[1] >= intervalMin && data[1] <= intervalMax {
				interactMap[data[0]] = true
			}
		}
		staleServers[k] = n - int32(len(interactMap))
	}
	return staleServers
}
