package realsmile

import "testing"

func TestCombinationUtils_Error(t *testing.T) {
	combination, _ := Combination.GetCombination(60, 3)
	for index, item := range combination.Values() {
		Log.Infof(" item[%v] = %v", index, item)
	}

}
