package table

import (
	"fmt"
	"sort"
	"strconv"
)

func RenderSum(one float64, num string) (float64, error) {
	n, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}

	sum := float64(n) * one

	fmt.Println(sum)
	return sum, err
}

func RenderList(rendr map[string]float64) string {
	list := rendr
	keys := make([]string, 0, len(list))
	for k := range list {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, list[k])
	}
	return ""
}
