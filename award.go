package main

import (
	"sort"
)

func sortbyPoint(cs [5]Card) [5]Card {

	var newCs []Card
	for i := 0; i < 5; i++ {
		newCs = append(newCs, cs[i])
	}

	// 依年齡由小而大排序
	sort.SliceStable(newCs, func(i, j int) bool {
		return newCs[i].P < newCs[j].P
	})

	var newCs2 [5]Card
	for index, v := range newCs {
		newCs2[index] = v
	}

	return newCs2
}

func check(csOri [5]Card) int {

	cs := sortbyPoint(csOri)

	// 先看是否順子
	if isStraight(cs) {
		// 再看是否同花
		if isFlush(cs) {
			// 最多同花大順
			if isRoyalFlush(cs) {
				return 1001 // 同花大順 1001
			}

			return 1003 // 同花順 1003
		}
		return 1007 // 順子 1007

	} else {
		if isThreeOfAKind(cs) {

			if isFiveOfAKind(cs) {
				return 1002 // 5支
			}

			if isFourOfAKind(cs) {
				return 1004 // 鐵支 1004
			}

			if isFullHouse(cs) {
				return 1005 // 葫蘆 1005
			}

			if isFlush(cs) {
				return 1006 // 同花 1006
			}
			return 1008 // 三條 1008

		} else if isFlush(cs) {
			return 1006 // 同花 1006

		} else if isTwoPair(cs) {
			// two pair 是任何牌湊兩對
			return 1009 // 兩對 1009

		} else if isOnePair(cs) {
			// one pair 是僅限JQKA湊對
			return 1010 // 一對 1010
		}
	}
	loserCards(cs)

	return 1000 // 散牌
}

// isRoyalFlush 同花大順 1001
// 已經先檢查過順子+同花
func isRoyalFlush(cs [5]Card) bool {
	// 兩張鬼牌，鬼牌一定在最後兩張
	if cs[3].N == 53 && cs[4].N == 54 {

		// 有替代A-> 前3張大於等於10
		if cs[0].P >= 10 && cs[1].P >= 10 && cs[2].P >= 10 {
			return true
		}
		// 不是替代A -> 前1張必有A (先排序過)
		// 不能有2 3 4
		if cs[0].P == 1 && cs[1].P != 2 && cs[1].P != 3 && cs[1].P != 4 {
			return true
		}

		// 一張鬼牌，鬼牌一定在最後一張
	} else if cs[4].N == 53 || cs[4].N == 54 {

		// 鬼牌是替代A
		if cs[0].P == 10 && cs[1].P == 11 && cs[2].P == 12 && cs[3].P == 13 {
			return true
			// 鬼牌不是替代A,A一定會再第一張,第二張是10或11即可
		} else if (cs[0].P == 1 && cs[1].P == 10) || (cs[0].P == 1 && cs[1].P == 11) {
			return true
		}
	} else {
		// 沒有鬼牌
		if cs[0].P == 1 && cs[1].P == 10 && cs[2].P == 11 && cs[3].P == 12 && cs[4].P == 13 {
			return true
		}
	}

	return false
}

// isFiveOfAKind 5支 1002
func isFiveOfAKind(cs [5]Card) bool {
	// 兩張鬼牌
	if cs[3].N == 53 && cs[4].N == 54 {

		// 前三張一樣
		if cs[0].P == cs[1].P && cs[0].P == cs[2].P {
			return true
		}

		// 一張鬼牌
	} else if cs[4].N == 53 || cs[4].N == 54 {

		// 前四張一樣
		if cs[0].P == cs[1].P && cs[0].P == cs[2].P && cs[0].P == cs[3].P {
			return true
		}

	}
	// 沒有鬼牌
	// 就不可能5支

	return false
}

// isFourOfAKind 鐵支 1004
func isFourOfAKind(cs [5]Card) bool {
	// 兩張鬼牌
	if cs[3].N == 53 && cs[4].N == 54 {

		// 前三張，任兩張一樣
		if cs[0].P == cs[1].P || cs[0].P == cs[2].P || cs[1].P == cs[2].P {
			return true
		}

		// 一張鬼牌
	} else if cs[4].N == 53 || cs[4].N == 54 {

		// 不是替代主要，前4張一樣->5支

		// 替代主要，會有3張一樣
		// 前面4個位置取3位

		// AAA B J

		if (cs[1].P == cs[2].P && cs[1].P == cs[3].P) ||
			(cs[0].P == cs[2].P && cs[0].P == cs[3].P) ||
			(cs[0].P == cs[1].P && cs[0].P == cs[3].P) ||
			(cs[0].P == cs[1].P && cs[0].P == cs[2].P) {
			return true
		}

		// 沒有鬼牌
		// 4張一樣
	} else {
		var markMap = make(map[int]int) // 出現的點數和次數
		// 跑一圈處理完
		for i := 0; i < 5; i++ {
			markMap[cs[i].P]++
		}
		// 兩對應該只有2種點數，並且其中一種有4個
		if len(markMap) == 2 {
			for _, v := range markMap {
				if v == 4 || v == 1 {
					return true
				}
			}
		}
	}

	return false
}

// isFullHouse 葫蘆 1005
func isFullHouse(cs [5]Card) bool {
	// 兩張鬼牌
	if cs[3].N == 53 && cs[4].N == 54 {
		// 如果其它牌有兩張以上，一定會被湊鐵支/5支
		// 如果其它牌都只有一張，一定只能湊三條
		return false

		// 一張鬼牌
	} else if cs[4].N == 53 || cs[4].N == 54 {

		// 其它4張一定是two pair的形式
		var markMap = make(map[int]int) // 出現的點數和次數
		// 跑一圈處理完
		for i := 0; i < 4; i++ {
			markMap[cs[i].P]++
		}
		// 前面4張是兩對，應該只有2種點數
		if len(markMap) == 2 {
			return true
		}

		// 沒有鬼牌
	} else {
		// AAA BB
		if cs[0].P == cs[1].P && cs[1].P == cs[2].P && cs[3].P == cs[4].P {
			return true
			// AA BBB
		} else if cs[0].P == cs[1].P && cs[2].P == cs[3].P && cs[3].P == cs[4].P {
			return true
		}
	}

	return false
}

// isFlush 同花 1006
func isFlush(cs [5]Card) bool {
	// 兩張鬼牌
	if cs[3].N == 53 && cs[4].N == 54 {

		// 鬼牌一定在最後2張，前3張同色
		if cs[1].C == cs[0].C && cs[2].C == cs[1].C {
			return true
		}

		// 一張鬼牌
	} else if cs[4].N == 53 || cs[4].N == 54 {

		// 鬼牌一定在最後一張，前4張同色
		if cs[1].C == cs[0].C && cs[2].C == cs[1].C && cs[3].C == cs[2].C {
			return true
		}

		// 沒有鬼牌
	} else {
		// 5張同色
		if cs[1].C == cs[0].C && cs[2].C == cs[1].C && cs[3].C == cs[2].C && cs[4].C == cs[3].C {
			return true
		}
	}

	return false
}

// isStraight 順子 1007
func isStraight(cs [5]Card) bool {

	// A 10 J Q K

	// 兩張鬼牌
	if cs[3].N == 53 && cs[4].N == 54 {

		// 鬼牌不是替代A
		if cs[0].P == 1 {
			// A 10 J Q K 牌形
			//  A  10 J
			//  A  10 Q
			//  A  10 K
			//  A  J  Q
			//  A  J  K
			//  A  Q  K
			if (cs[0].P == 1 && cs[1].P == 10 && cs[2].P == 11) ||
				(cs[0].P == 1 && cs[1].P == 10 && cs[2].P == 12) ||
				(cs[0].P == 1 && cs[1].P == 10 && cs[2].P == 13) ||
				(cs[0].P == 1 && cs[1].P == 11 && cs[2].P == 12) ||
				(cs[0].P == 1 && cs[1].P == 11 && cs[2].P == 13) ||
				(cs[0].P == 1 && cs[1].P == 12 && cs[2].P == 13) {
				return true
				// A 2 3 4 5 牌形
			} else if (cs[0].P == 1 && cs[1].P == 2 && cs[2].P == 3) ||
				(cs[0].P == 1 && cs[1].P == 2 && cs[2].P == 4) ||
				(cs[0].P == 1 && cs[1].P == 2 && cs[2].P == 5) ||
				(cs[0].P == 1 && cs[1].P == 3 && cs[2].P == 4) ||
				(cs[0].P == 1 && cs[1].P == 3 && cs[2].P == 5) ||
				(cs[0].P == 1 && cs[1].P == 4 && cs[2].P == 5) {
				return true
			}
		} else {
			// p1 p2 O O X X X
			// p4 p5 X X X O O
			// p1 p5 O X X X O
			if cs[1].P == cs[0].P+1 && cs[2].P == cs[1].P+1 {
				return true
				// p1 p3 O X O X X
				// p2 p5 X O X X O
			} else if cs[1].P == cs[0].P+2 && cs[2].P == cs[1].P+1 {
				return true
				// p1 p4 O X X O X
				// p3 p5  X X O X O
			} else if cs[1].P == cs[0].P+1 && cs[2].P == cs[1].P+2 {
				return true
				// p2 p3 X O O X X
			} else if cs[1].P == cs[0].P+3 && cs[2].P == cs[1].P+1 {
				return true
				// p2 p4 X O X O X
			} else if cs[1].P == cs[0].P+2 && cs[2].P == cs[1].P+2 {
				return true
				// p3 p4 X X O O X
			} else if cs[1].P == cs[0].P+1 && cs[2].P == cs[1].P+3 {
				return true
			}

		}

		// 一張鬼牌
	} else if cs[4].N == 53 || cs[4].N == 54 {

		// A 10 J Q K  要另外處理
		// A 10 J Q
		// A 10 J K
		// A 10 Q K
		// A J Q K
		if cs[0].P == 1 && ((cs[1].P == 10 && cs[2].P == 11 && cs[3].P == 12) ||
			(cs[1].P == 10 && cs[2].P == 11 && cs[3].P == 13) ||
			(cs[1].P == 10 && cs[2].P == 12 && cs[3].P == 13) ||
			(cs[1].P == 11 && cs[2].P == 12 && cs[3].P == 13)) {
			return true

		} else {
			// p1 O X X X X
			// p5 X X X X O
			if cs[1].P == cs[0].P+1 && cs[2].P == cs[1].P+1 && cs[3].P == cs[2].P+1 {
				return true
				// p2 X O X X X
			} else if cs[1].P == cs[0].P+2 && cs[2].P == cs[1].P+1 && cs[3].P == cs[2].P+1 {
				return true
				// p3 X X O X X
			} else if cs[1].P == cs[0].P+1 && cs[2].P == cs[1].P+2 && cs[3].P == cs[2].P+1 {
				return true
				// p4 X X X O X
			} else if cs[1].P == cs[0].P+1 && cs[2].P == cs[1].P+1 && cs[3].P == cs[2].P+2 {
				return true
			}
		}

		// 沒有鬼牌
	} else {
		// 連續5張
		if cs[1].P == cs[0].P+1 && cs[2].P == cs[1].P+1 && cs[3].P == cs[2].P+1 && cs[4].P == cs[3].P+1 {
			return true
			// A 10 J Q K
		} else if cs[0].P == 1 && cs[1].P == 10 && cs[2].P == 11 && cs[3].P == 12 && cs[4].P == 13 {
			return true
		}
	}

	return false
}

// isThreeOfAKind 三條 1008
func isThreeOfAKind(cs [5]Card) bool {

	// 兩張鬼牌,一定最少三條
	if cs[3].N == 53 && cs[4].N == 54 {
		return true
		// 一張鬼牌
	} else if cs[4].N == 53 || cs[4].N == 54 {

		// 不是替代主要
		// AAA B C 替B或C
		// A BBB C 替A
		if cs[0].P == cs[1].P && cs[1].P == cs[2].P {
			return true
			// A BBB C 替C
			// A B CCC 替A或B
		} else if cs[1].P == cs[2].P && cs[2].P == cs[3].P {
			return true
		}

		// 替代主要
		// 其中有兩張一樣 C4取2
		if cs[0].P == cs[1].P || cs[0].P == cs[2].P || cs[0].P == cs[3].P || cs[1].P == cs[2].P || cs[1].P == cs[3].P || cs[2].P == cs[3].P {
			return true
		}

		// 沒有鬼牌
		// A BBB C
		// AAA B C
		// A B CCCC
	} else {
		if cs[1].P == cs[2].P && cs[2].P == cs[3].P {
			return true
		} else if cs[0].P == cs[1].P && cs[1].P == cs[2].P {
			return true
		} else if cs[2].P == cs[3].P && cs[3].P == cs[4].P {
			return true
		}
	}

	return false
}

// isTwoPair 兩對 1009
func isTwoPair(cs [5]Card) bool {

	// 兩張鬼牌
	if cs[3].N == 53 && cs[4].N == 54 {
		// 兩張鬼牌,一定最少三條
		return false
		// 一張鬼牌
	} else if cs[4].N == 53 || cs[4].N == 54 {
		// 除了鬼牌之外，4張內一定有一對
		if (cs[0].P == cs[1].P || cs[0].P == cs[2].P || cs[0].P == cs[3].P) ||
			(cs[1].P == cs[2].P || cs[1].P == cs[3].P) ||
			cs[2].P == cs[3].P {
			return true
		}

		// 沒有鬼牌
	} else {
		var markMap = make(map[int]int) // 出現的點數和次數
		// 跑一圈處理完
		for i := 0; i < 5; i++ {
			markMap[cs[i].P]++
		}
		// 兩對應該只有3種點數
		if len(markMap) == 3 {
			return true
		}
	}
	return false
}

// isOnePair 一對 1010
func isOnePair(cs [5]Card) bool {

	// 兩張鬼牌
	if cs[3].N == 53 && cs[4].N == 54 {
		// 兩張鬼牌,一定最少三條
		return false
		// 一張鬼牌
	} else if cs[4].N == 53 || cs[4].N == 54 {
		// 一定4張都不一樣
		// 然後有一張是JQKA
		if cs[0].P != cs[1].P && cs[0].P != cs[2].P && cs[0].P != cs[3].P &&
			cs[1].P != cs[2].P && cs[1].P != cs[3].P &&
			cs[2].P != cs[3].P {

			for i := 0; i < 4; i++ {
				if cs[i].P == 1 || cs[i].P == 11 || cs[i].P == 12 || cs[i].P == 13 {
					return true
				}
			}
		}

		// 沒有鬼牌
	} else {

		var t1, t11, t12, t13 int // 出現JQKA的次數
		// 跑一圈處理完
		for i := 0; i < 5; i++ {
			if cs[i].P == 1 || cs[i].P >= 11 {
				switch cs[i].P {
				case 1:
					t1++
				case 11:
					t11++
				case 12:
					t12++
				case 13:
					t13++
				}
			}
		}
		if t1 > 1 || t11 > 1 || t12 > 1 || t13 > 1 {
			return true
		}
	}

	return false
}

// 5 取 3
//
// O O X X X
// O X O X X
// O X X O X
// O X X X O
// X O O X X
//
// X O X O X
// X O X X O
// X X O O X
// X X O X O
// X X X O O

// 5 取 4
//
// O X X X X
// X O X X X
// X X O X X
// X X X O X
// X X X X O
