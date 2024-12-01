package main

import (
	. "fmt"
	. "math"
	. "sort"
)

func main() {
	var a, b, count1, count2 int
	var as, bs []int
	counts := map[int]int{}
	var err error

	for err == nil {
		_, err = Scanf("%d %d", &a, &b)
		as, bs = append(as, a), append(bs, b)
		counts[b]++
	}

	Ints(as)
	Ints(bs)
	
	for i, a := range as {
		count1 += int(Abs(float64(a - bs[i])))
		count2 += a * counts[a]
	}
	println(count1, count2)
}
