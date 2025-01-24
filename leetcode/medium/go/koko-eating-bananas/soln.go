package kokoeatingbananas

import (
	"math"
	"slices"
)

// started at 1:30
// ended at 2:10

func minEatingSpeed(piles []int, h int) int {
	small := 0
	large := slices.Max(piles)
	prevMid := -1
	if h == len(piles) {
		return large
	}

	for small != large {
		mid := int(math.Ceil(float64(small) + float64(large-small)/2))
		if mid == prevMid {
			if isValid(small, h, piles) {
				return large
			}
			return large
		}

		if isValid(mid, h, piles) {
			large = mid
		} else {
			small = mid
		}
		prevMid = mid
	}
	return small
}

func isValid(rate, h int, piles []int) bool {
	cumH := 0
	for _, nBananas := range piles {
		cumH += int(math.Ceil(float64(nBananas) / float64(rate)))
	}

	return cumH <= h
}
