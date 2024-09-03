package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// https://github.com/logpai/logparser/blob/master/logparser/Drain/Drain.py

// cluster ...
type clustersss struct {
	event []string // log template
	logs  int64
}

func adj(as []int, idx, leng int) {
	if idx >= leng {
		return
	}
	maxi := idx
	l := (idx*2 + 1)
	r := (idx + 1) * 2
	if l < leng && as[maxi] < as[l] {
		maxi = l
	}
	if r < leng && as[maxi] < as[r] {
		maxi = r
	}
	if maxi == idx {
		return
	}

	as[idx],as[maxi] = as[maxi],as[idx]
	adj(as, maxi, leng)
	return
}

func findKthLargest(nums []int, k int) int {
	for i := len(nums) / 2; i >= 0; i-- {
		adj(nums, i, len(nums))
	}

	leng := len(nums)
	for i := 1; i < k; i++ {
		nums[0], nums[leng-1] = nums[leng-1], nums[0]
		leng--
		adj(nums, 0, leng)
	}
	return nums[0]
}

func main() {

	return
	// csv.NewWriter
	dr, err := os.OpenFile("/tmp/env_dr.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	live, err := os.OpenFile("/tmp/env_live.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// abc def ghi jkl mno pqrs tuv wxyz
	drs, err := csv.NewReader(dr).ReadAll()
	if err != nil {
		panic(err)
	}
	lives, err := csv.NewReader(live).ReadAll()
	if err != nil {
		panic(err)
	}
	type app struct {
		name string
		id   string
	}
	set := map[string]app{}
	for _, dri := range drs {
		m := strings.Split(dri[0], "-dr-")
		set[m[0]] = app{
			name: dri[1],
			id:   dri[2],
		}
	}
	for _, li := range lives {
		m := strings.Split(li[0], "-live-")
		appi, ok := set[m[0]]
		if !ok {
			continue
		}
		li[1] = appi.name
		li[2] = appi.id
	}
	out, err := os.OpenFile("/tmp/out.csv", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(csv.NewWriter(out).WriteAll(lives))
}

type LogParserTree struct {
	Depth               int
	MaxChild            int
	SimilarityThreshold float64
	MaxTemplateSize     int
}

// DefaultFilterRemoveDigitBegin ...
func DefaultFilterRemoveDigitBegin(tokens []string) []string {
	isDigital := func(r rune) bool {
		return r >= '0' && r <= '9'
	}
	var digitalBegin bool
	for i, token := range tokens {
		for _, ri := range token {
			digitalBegin = isDigital(ri)
			break
		}
		if digitalBegin {
			tokens[i] = "*"
		}
	}
	return tokens
}
