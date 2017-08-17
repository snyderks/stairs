package stairs

import "testing"

type testWeighted struct {
	name   string
	weight int
}

type testWeights []testWeighted

type testFloatWeighted struct {
	name   string
	weight float64
}

type testFloatWeights []testFloatWeighted

func buildWeightedArray() WeightedItems {
	a := make(testWeights, 0)

	a = append(a, testWeighted{"str", 1})
	a = append(a, testWeighted{"str2", 2})
	a = append(a, testWeighted{"str3", 5})

	var w WeightedItems

	for i, val := range a {
		w = append(w, WeightedItem{val.weight, i})
	}

	return w
}

// TestBuild checks to see if a basic array can
// be correctly processed by the builder function.
func TestBuild(t *testing.T) {
	w := buildWeightedArray()

	_, err := w.BuildCDF()

	if err != nil {
		t.Fail()
	}
}

// TestSelectItemsFrom checks to see if the
// function returned after building correctly
// returns indices in the correct range.
func TestSelectItemsFrom(t *testing.T) {
	w := buildWeightedArray()

	f, err := w.BuildCDF()

	if err != nil {
		t.Fail()
	}

	// Check the returned function many times
	for i := 0; i < 100; i++ {
		index := f()
		if index < 0 || index > len(w) {
			t.Fail()
		}
	}
}

// TestEmpty checks to make sure that a CDF can't be built
// for an empty array.
func TestEmpty(t *testing.T) {
	var w WeightedItems

	_, err := w.BuildCDF()

	if err == nil {
		t.Fail()
	}
}

// TestEmptyFloat checks to make sure that a CDF can't be built
// for an empty array.
func TestEmptyFloat(t *testing.T) {
	var w WeightedItemsFloat

	_, err := w.BuildCDF()

	if err == nil {
		t.Fail()
	}
}

// TestZeroWeight checks to make sure that a CDF
// can't be built with any zero-weight items.
func TestZeroWeight(t *testing.T) {
	var w WeightedItems

	w = append(w, WeightedItem{5, 0})
	w = append(w, WeightedItem{0, 1})
	w = append(w, WeightedItem{3, 2})

	_, err := w.BuildCDF()

	if err == nil {
		t.Fail()
	}
}

// TestZeroWeightFloat checks to make sure that a CDF
// can't be built with any zero-weight items.
func TestZeroWeightFloat(t *testing.T) {
	var w WeightedItemsFloat

	w = append(w, WeightedItemFloat{3.7473, 0})
	w = append(w, WeightedItemFloat{0, 1})
	w = append(w, WeightedItemFloat{1.373, 2})

	_, err := w.BuildCDF()

	if err == nil {
		t.Fail()
	}
}

// TestNegativeWeight checks to make sure that a CDF
// can't be built with any negative-weight items.
func TestNegativeWeight(t *testing.T) {
	var w WeightedItems

	w = append(w, WeightedItem{5, 0})
	w = append(w, WeightedItem{-64, 1})
	w = append(w, WeightedItem{3, 2})

	_, err := w.BuildCDF()

	if err == nil {
		t.Fail()
	}
}

// TestNegativeWeightFloat checks to make sure that a CDF
// can't be built with any negative-weight items.
func TestNegativeWeightFloat(t *testing.T) {
	var w WeightedItemsFloat

	w = append(w, WeightedItemFloat{3.7473, 0})
	w = append(w, WeightedItemFloat{-100.937, 1})
	w = append(w, WeightedItemFloat{1.373, 2})

	_, err := w.BuildCDF()

	if err == nil {
		t.Fail()
	}
}

// // TestDuplicateIndices checks that the builder
// // rejects a weighted array with two items that point
// // to the same index.
// func TestDuplicateIndices

func buildWeightedFloatArray() WeightedItemsFloat {
	a := make(testFloatWeights, 0)

	a = append(a, testFloatWeighted{"str", 1.5})
	a = append(a, testFloatWeighted{"str2", 2.33})
	a = append(a, testFloatWeighted{"str3", 5.8999})

	var w WeightedItemsFloat

	for i, val := range a {
		w = append(w, WeightedItemFloat{val.weight, i})
	}

	return w
}

// TestBuildFloat checks to see if a floating point
// weighted array can be correctly processed
// by the builder function.
func TestBuildFloat(t *testing.T) {
	w := buildWeightedFloatArray()

	_, err := w.BuildCDF()

	if err != nil {
		t.Fail()
	}
}

// TestSelectItemsFromFloat checks to see if the
// function returned after building correctly
// returns indices in the correct range.
func TestSelectItemsFromFloat(t *testing.T) {
	w := buildWeightedFloatArray()

	f, err := w.BuildCDF()

	if err != nil {
		t.Fail()
	}

	// Check the returned function many times
	for i := 0; i < 100; i++ {
		index := f()
		if index < 0 || index > len(w) {
			t.Fail()
		}
	}
}
