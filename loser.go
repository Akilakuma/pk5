package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var loserTypeMap = make(map[string]int)

func loserCards(cs [5]Card) {

	var tempStrKey string

	for _, value := range cs {
		if tempStrKey == "" {
			tempStrKey += strconv.Itoa(value.P)
		} else {
			tempStrKey += "_" + strconv.Itoa(value.P)
		}

	}
	loserTypeMap[tempStrKey]++
}

// 應該可以改成goroutine版本，計算速度更快
func loserTryNext() {

	// 獲獎數量
	var (
		i                = 0
		s                = time.Now()
		AwardMap         = make(map[int]int)
		total            = len(loserTypeMap)
		num49    float64 = 1906884
	)

	// 排列組合的可能性
	// 剔除5張，剩下49個位置的排列組合
	combinationList := Combination(49)

	for key := range loserTypeMap {
		pointList := strings.Split(key, "_")

		var lcs [5]int
		lcs[0], _ = strconv.Atoi(pointList[0])
		lcs[1], _ = strconv.Atoi(pointList[1])
		lcs[2], _ = strconv.Atoi(pointList[2])
		lcs[3], _ = strconv.Atoi(pointList[3])
		lcs[4], _ = strconv.Atoi(pointList[4])

		// fmt.Println(lcs)

		// 製作專屬這些key的卡池
		lCards := makeLoserCards(lcs)

		for _, seriesNum := range combinationList {
			var cc [5]Card

			// 根據卡排的流水號做排序
			sort.Ints(seriesNum)
			for index, num := range seriesNum {
				// 0 1 2 3 4 位置，塞進卡牌留水號對應的內容
				cc[index] = lCards[num]
			}

			r := check(cc)

			AwardMap[r]++

		}
		fmt.Println("i:", i, ", total:", total)
		i++
	}

	fmt.Println()
	for aw, num := range AwardMap {
		numAvg := num / total
		numPercent := float64(numAvg) / num49

		fmt.Println(aw, numAvg, numPercent)
	}

	fmt.Println()
	fmt.Println("loser return 花費時間:", time.Since(s).Seconds(), " 秒")
}
