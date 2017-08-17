// Package stairs provides functions to randomly select items from a weighted array.
package stairs

import (
	"errors"
	"math"
	"math/rand"
	"sort"
	"time"
)

// EPSILON is an arbitrarily small floating-point
// number used for equality comparison when
// searching through a floating-point CDF.
const EPSILON = 0.00001

// WeightedItem contains the weight for the item
// and the index it represents in the original array.
type WeightedItem struct {
	// The relative weight assigned to the item
	Weight int
	// Index is the location in the original array
	// for the item
	Index int
}

// WeightedItems is an array of WeightedItem interfaces.
type WeightedItems []WeightedItem

// WeightedItemFloat contains the floating-point weight
// for the item and the index it represents in the
// original array.
type WeightedItemFloat struct {
	// The relative weight assigned to the item
	Weight float64
	// Index is the location in the original array
	// for the item
	Index int
}

// WeightedItemsFloat is an array of WeightedItemFloat interfaces.
type WeightedItemsFloat []WeightedItemFloat

// Sort interface implementation
// sort.Sort will sort by weight ascending
func (s WeightedItems) Len() int {
	return len(s)
}

func (s WeightedItems) Less(i, j int) bool {
	return s[i].Weight < s[j].Weight
}

func (s WeightedItems) Swap(i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

// Sort interface implementation
// sort.Sort will sort by weight ascending
func (s WeightedItemsFloat) Len() int {
	return len(s)
}

func (s WeightedItemsFloat) Less(i, j int) bool {
	return s[i].Weight < s[j].Weight
}

func (s WeightedItemsFloat) Swap(i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

const tooShortErr = "Array of items must be longer than 0."
const zeroWeightErr = "All items must have a positive weight."

// BuildCDF converts a weighted array into a function that will return
// random elements from it, when called.
func (s WeightedItems) BuildCDF() (func() int, error) {
	// Reject empty arrays
	if len(s) <= 0 {
		return nil, errors.New(tooShortErr)
	}

	// Sort the array ascending by weight
	sort.Sort(s)

	// Make sure first item has positive weight
	if s[0].Weight <= 0 {
		return nil, errors.New(zeroWeightErr)
	}

	// Accumulate the weights
	for i := 1; i < len(s); i++ {
		// Make sure all items have positive weight
		if s[i].Weight <= 0 {
			return nil, errors.New(zeroWeightErr)
		}

		s[i].Weight += s[i-1].Weight
	}

	// Initialize random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	searchCDF := func() int {
		// Picking a random number in the range [1, max weight + 1)
		num := r.Intn(s[len(s)-1].Weight) + 1

		// Binary search! Look for the number generated.
		// Right and left are the bounds for the binary search
		right := len(s) - 1
		left := 0

		for {
			// check the middle of the bounds
			m := (left + right) / 2 // m stands for middle
			valm := s[m].Weight

			if valm == num { // exact match
				return s[m].Index
			} else if valm < num {
				// Middle item is less than number

				if m == len(s)-1 {
					// only option is rightmost item
					return s[m].Index
				} else if s[m+1].Weight > num {
					// return the right item when
					// the search is finished
					// and left between two items.
					return s[m+1].Index
				}
				// bring left bound to the middle
				left = m + 1
			} else {
				// Middle item is more than number

				if m == 0 || s[m-1].Weight <= num {
					// Can't move left, so return the middle.
					return s[m].Index
				}
				// bring right bound to the middle
				right = m - 1
			}
		}
	}
	return searchCDF, nil
}

// BuildCDF converts a weighted array into a function that will return
// random elements from it, when called.
// Allows for use of floating-point weights.
func (s WeightedItemsFloat) BuildCDF() (func() int, error) {
	// Reject empty arrays
	if len(s) <= 0 {
		return nil, errors.New(tooShortErr)
	}

	// Sort the array ascending by weight
	sort.Sort(s)

	// Make sure first item has positive weight
	if s[0].Weight <= 0 {
		return nil, errors.New(zeroWeightErr)
	}

	// Accumulate the weights
	for i := 1; i < len(s); i++ {
		// Make sure all items have positive weight
		if s[0].Weight <= 0 {
			return nil, errors.New(zeroWeightErr)
		}

		s[i].Weight += s[i-1].Weight
	}

	// Initialize random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	searchCDF := func() int {
		// Picking a random number in the range [1, max weight + 1)
		num := r.Float64()*(s[len(s)-1].Weight-1) + 1

		// Binary search! Look for the number generated.
		// Right and left are the bounds for the binary search
		right := len(s) - 1
		left := 0

		for {
			// check the middle of the bounds
			m := (left + right) / 2 // m stands for middle
			valm := s[m].Weight

			if math.Abs(valm-num) <= EPSILON { // exact match
				return s[m].Index
			} else if valm < num {
				// Middle item is less than number

				if m == len(s)-1 {
					// only option is rightmost item
					return s[m].Index
				} else if s[m+1].Weight > num {
					// return the right item when
					// the search is finished
					// and left between two items.
					return s[m+1].Index
				}
				// bring left bound to the middle
				left = m + 1
			} else {
				// Middle item is more than number

				if m == 0 || s[m-1].Weight <= num {
					// Can't move left, so return the middle.
					return s[m].Index
				}
				// bring right bound to the middle
				right = m - 1
			}
		}
	}
	return searchCDF, nil
}
