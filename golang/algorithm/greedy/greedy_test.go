package greedy_test

import (
	"learn/golang/algorithm/greedy"
	"testing"
)

func TestStationCover(t *testing.T) {
	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}
	states := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	ans := []string{"kone", "ktwo", "kfive", "kthree"}

	res := greedy.StationCover(states, stations)
	if !compareSlice(ans, res) {
		t.Errorf("\nexcept:%v\nactual:%v", ans, res)
	}

}

// 比较切片是否一致
func compareSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
