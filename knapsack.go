package util

type Item interface {
	Weight() int32
	Value() int32
}

func Max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

// Pick items to maximize value, within constrain of weight capacity
func PickItems(items []Item, capacity int32) (value int32, picked []Item) {
	// int32 to save memory usage...
	dp := make([][]int32, len(items)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int32, capacity+1)
	}
	for i := 0; i <= len(items); i++ {
		var w int32
		for w = 0; w <= capacity; w++ {
			if i == 0 || w == 0 {
				dp[i][w] = 0
			} else if items[i-1].Weight() <= w {
				dp[i][w] = Max(dp[i-1][w],
					items[i-1].Value()+dp[i-1][w-items[i-1].Weight()])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	value = dp[len(items)][capacity]
	i := len(items)
	w := capacity
	v := value
	for i > 0 && w > 0 && v > 0 {
		if dp[i][w] == dp[i-1][w] {
			i = i - 1
			continue
		}
		picked = append(picked, items[i-1])
		v = v - items[i-1].Value()
		w = w - items[i-1].Weight()
		i = i - 1
	}
	return value, picked
}
