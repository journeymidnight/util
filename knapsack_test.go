package util

import (
	"testing"
)

type item struct {
	weight int
	value  int
}

func (i item) Weight() int {
	return i.weight
}

func (i item) Value() int {
	return i.value
}

func correctWeights(picked []Item, weights []int) bool {
	if len(picked) != len(weights) {
		return false
	}

	memo := make(map[int]int)
	for _, weights := range weights {
		memo[weights] += 1
	}
	for _, item := range picked {
		memo[item.Weight()] -= 1
		if memo[item.Weight()] < 0 {
			return false
		}
	}
	return true
}

func TestPickItems(t *testing.T) {
	case1 := []Item{
		item{10, 60},
		item{20, 100},
		item{30, 120},
	}
	v, picked := PickItems(case1, 50)
	if v != 220 || !correctWeights(picked, []int{30, 20}) {
		t.Error("back knapsack:", v, picked)
	}

	case2 := []Item{
		item{20, 40},
		item{10, 100},
		item{40, 50},
		item{30, 60},
	}
	v, picked = PickItems(case2, 60)
	if v != 200 || !correctWeights(picked, []int{30, 20, 10}) {
		t.Error("back knapsack:", v, picked)
	}

	case3 := []Item{
		item{1, 10},
		item{2, 15},
		item{3, 40},
	}
	v, picked = PickItems(case3, 6)
	if v != 65 || !correctWeights(picked, []int{1, 2, 3}) {
		t.Error("back knapsack:", v, picked)
	}

	case4 := []Item{
		item{1, 10},
		item{1, 20},
		item{1, 30},
	}
	v, picked = PickItems(case4, 2)
	if v != 50 || !correctWeights(picked, []int{1, 1}) {
		t.Error("back knapsack:", v, picked)
	}

	case5 := []Item{
		item{3, 3},
		item{10, 10},
		item{10, 10},
		item{2, 2},
		item{2, 2},
	}
	v, picked = PickItems(case5, 16)
	if v != 15 || !correctWeights(picked, []int{10, 2, 3}) {
		t.Error("back knapsack:", v, picked)
	}
}
