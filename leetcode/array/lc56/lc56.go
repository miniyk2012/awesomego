package main

import (
	"fmt"
	"sort"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 合并区间
func merge(intervals [][]int) [][]int {
	// 所有区间按左端排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var merged [][]int
	for _, interval := range intervals {
		if len(merged) == 0 {
			merged = append(merged, interval)
			continue
		}
		// 排完序后, 只需和最后一项合并即可
		last := merged[len(merged)-1]
		if interval[0] > last[1] {
			merged = append(merged, interval)
			continue
		}
		merged[len(merged)-1][1] = max(last[1], interval[1])
	}
	return merged
}

func main() {
	merged := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Println(merged)
	merged = merge([][]int{{1, 4}, {4, 5}})
	fmt.Println(merged)
}
