package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	// testCase()
	// testCase2()
	// cList := makeCards()
	// fmt.Println(cList)

	work()

}
func work() {
	// 排列組合的可能性
	combinationList := Combination(54)

	// 卡池
	cards := makeCards()

	var (
		AwardMap         = make(map[int]int)
		s                = time.Now()
	)

	for _, seriesNum := range combinationList {
		var cc [5]Card

		// 根據卡排的流水號做排序
		sort.Ints(seriesNum)
		for index, num := range seriesNum {
			// 0 1 2 3 4 位置，塞進卡牌留水號對應的內容
			cc[index] = cards[num]
		}
		// fmt.Println(cc)
		r := check(cc)
		// s := codeMean(r)

		AwardMap[r]++

	}

	// fmt.Println(bb.result)
	fmt.Println()
	fmt.Println("花費時間:", time.Since(s).Seconds(), " 秒")
	fmt.Println()
	// fmt.Println("AwardMap:", AwardMap)
	fmt.Println("loser Map 數量:", len(loserTypeMap))
	loserTryNext()

}

func testCase() {
	var cc [5]Card

	cc[0].N = 2
	cc[0].P = 2
	cc[0].C = "A"

	cc[1].N = 12
	cc[1].P = 12
	cc[1].C = "A"

	cc[2].N = 13
	cc[2].P = 13
	cc[2].C = "B"

	cc[3].N = 16
	cc[3].P = 3
	cc[3].C = "C"

	cc[4].N = 54
	cc[4].P = 0
	cc[4].C = "J"

	ccc := sortbyPoint(cc)

	fmt.Println(ccc)

	r := check(cc)
	s := codeMean(r)

	fmt.Println(s)
}

func testCase2() {
	r := makeLoserCards([5]int{2, 4, 5, 6, 53})
	fmt.Println(r)
}

func codeMean(code int) string {
	switch code {
	case 1001:
		return "同花大順 1001"
	case 1002:
		return "5支 1002"
	case 1003:
		return "同花順 1003"
	case 1004:
		return "鐵支 1004"
	case 1005:
		return "葫蘆 1005"
	case 1006:
		return "同花 1006"
	case 1007:
		return "順子 1007"
	case 1008:
		return "三條 1008"
	case 1009:
		return "兩對 1009"
	case 1010:
		return "一對 1010"
	default:
		return "散牌 1000"
	}
}

func sortCard(cs [5]Card) {
	// 依流水號由小而大排序
	sort.SliceStable(cs, func(i, j int) bool {
		return cs[i].N < cs[j].N
	})
}
