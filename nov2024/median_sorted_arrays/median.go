package mediansortedarrays

import (
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)

	if len(nums1)%2 != 0 {
		return float64(nums1[len(nums1)/2])
	} else {
		lowIndex := int(len(nums1)/2) - 1
		highIndex := lowIndex + 1
		return float64(nums1[lowIndex]+nums1[highIndex]) / 2
	}
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	for i := 0; i < len(nums2); i++ {
		for j := 0; j < len(nums1); j++ {
			if nums1[j] < nums2[i] {
				if j+1 < len(nums1) && nums1[j+1] < nums2[i] {
					continue
				}

				nums1 = append(nums1[:j+1], append([]int{nums2[i]}, nums1[j+1:]...)...)
				break
			}
		}
	}

	if len(nums1)%2 != 0 {
		return float64(nums1[len(nums1)/2])
	} else {
		lowIndex := int(len(nums1)/2) - 1
		highIndex := lowIndex + 1
		return float64(nums1[lowIndex]+nums1[highIndex]) / 2
	}
}
