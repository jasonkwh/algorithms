package zigzag

import (
	"fmt"
	"strings"
)

func Convert(s string, numRows int) string {
	indexMap := make([]string, numRows)

	calc := 1
	inverse := true
	for i := 0; i < len(s); i++ {
		if calc-1 >= len(indexMap) {
			indexMap[len(indexMap)-1] = fmt.Sprintf("%s%s", indexMap[len(indexMap)-1], string(s[i]))
		} else {
			indexMap[calc-1] = fmt.Sprintf("%s%s", indexMap[calc-1], string(s[i]))
		}

		if calc == numRows && inverse {
			inverse = false
		}
		if calc == 1 && !inverse {
			inverse = true
		}
		if inverse {
			calc++
			continue
		}
		if !inverse {
			calc--
		}
	}

	return strings.Join(indexMap, "")
}
