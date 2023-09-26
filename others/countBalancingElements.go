package others

import (
	"sync"
	"sync/atomic"
)

func CountBalancingElements(arr []int32) int32 {
	count := int32(0)
	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			addCount := false
			if index == 0 {
				addCount = isBalanced(arr[index+1:])
			} else if index == len(arr)-1 {
				addCount = isBalanced(arr[:len(arr)-1])
			} else {
				testArr := appendSlice(index, arr)
				addCount = isBalanced(testArr)
			}

			if addCount {
				atomic.AddInt32(&count, 1)
			}
		}(i)
	}
	wg.Wait()

	return count
}

func appendSlice(index int, arr []int32) []int32 {
	testArr := make([]int32, len(arr))
	for i := 0; i < len(arr[:index]); i++ {
		testArr[i] = arr[i]
	}
	for i := index + 1; i < len(arr); i++ {
		testArr[i-1] = arr[i]
	}
	return testArr
}

func isBalanced(arr []int32) bool {
	if len(arr) == 0 {
		return false
	}

	sumeven, sumodd := int32(0), int32(0)
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			sumeven += arr[i]
			continue
		}
		sumodd += arr[i]
	}
	return sumeven == sumodd
}
