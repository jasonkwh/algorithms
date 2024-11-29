package others

import (
	"runtime"
	"sync"
	"sync/atomic"
)

func CountBalancingElements(arr []int32) int32 {
	count := int32(0)
	var wg sync.WaitGroup
	//var lock sync.Mutex

	for i := 0; i < len(arr); i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			addCount := false
			if index == 0 {
				addCount = isBalanced2(arr[index+1:])
			} else if index == len(arr)-1 {
				addCount = isBalanced2(arr[:len(arr)-1])
			} else {
				// isBalanced version 1
				// testArr := appendSlice(index, arr)
				// addCount = isBalanced(testArr)

				// isBalanced version 2
				// lock.Lock()
				// testVal := arr[index]
				// arr[index] = -1
				// addCount = isBalanced2(arr)
				// arr[index] = testVal
				// lock.Unlock()

				// this method introduces problems...
				// what happened if arr[index] changed to -1,
				// while the another goroutine (index+1) took this -1
				// in isBalanced2 calculation
				// calculation will be incorrect
				// testVal := atomic.LoadInt32(&arr[index])
				// atomic.StoreInt32(&arr[index], -1)
				// addCount = isBalanced2(arr)
				// atomic.StoreInt32(&arr[index], testVal)

				// isBalanced version 3
				addCount = isBalanced3(arr, index)
			}

			if addCount {
				atomic.AddInt32(&count, 1)
			}
		}(i)
	}
	wg.Wait()

	return count
}

// sharding method: distribute work loads to the threads
func CountBalancingElementsSharding(arr []int32) int32 {
	count := int32(0)

	var wg sync.WaitGroup

	if len(arr) > runtime.NumCPU() {
		for i := 0; i < runtime.NumCPU(); i++ {
			wg.Add(1)
			if i != runtime.NumCPU()-1 {
				go processChunks(arr[len(arr)/runtime.NumCPU()*i:len(arr)/runtime.NumCPU()*(i+1)], &count, &wg)
				continue
			}
			go processChunks(arr[len(arr)/runtime.NumCPU()*i:], &count, &wg)
		}
	} else {
		wg.Add(1)
		go processChunks(arr, &count, &wg)
	}

	wg.Wait()

	return count
}

func processChunks(arr []int32, count *int32, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < len(arr); i++ {
		addCount := false
		if i == 0 {
			addCount = isBalanced(arr[i+1:])
		} else if i == len(arr)-1 {
			addCount = isBalanced(arr[:len(arr)-1])
		} else {
			testArr := appendSlice(i, arr)
			addCount = isBalanced(testArr)
		}

		if addCount {
			atomic.AddInt32(count, 1)
		}
	}
}

// method 1: create new slice
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

// method 2: use existing slice
// space complexity is better compare to method 1
func isBalanced2(arr []int32) bool {
	if len(arr) == 0 {
		return false
	}

	skipped := false
	sumeven, sumodd := int32(0), int32(0)
	for i := 0; i < len(arr); i++ {
		index := i
		if arr[i] == -1 {
			skipped = true
			continue
		}
		if skipped {
			index = index - 1
		}
		if index%2 == 0 {
			sumeven += arr[i]
			continue
		}
		sumodd += arr[i]
	}
	return sumeven == sumodd
}

// method 3
// based on method 2, skip the provided index
func isBalanced3(arr []int32, skipIndex int) bool {
	if len(arr) == 0 {
		return false
	}

	skipped := false
	sumeven, sumodd := int32(0), int32(0)
	for i := 0; i < len(arr); i++ {
		index := i
		if i == skipIndex {
			skipped = true
			continue
		}
		if skipped {
			index = index - 1
		}
		if index%2 == 0 {
			sumeven += arr[i]
			continue
		}
		sumodd += arr[i]
	}
	return sumeven == sumodd
}
