package main

type Card struct {
	N int    // 流水號,1~54，最後兩張是joker
	P int    // 點數，1~13
	C string // 花色，A、B、C、D、J(joker)
}

func makeCards() map[int]Card {

	var cMap = make(map[int]Card)
	for i := 1; i <= 54; i++ {

		var c Card
		if i >= 1 && i <= 13 {
			c.N = i
			c.P = i
			c.C = "A"
		} else if i >= 14 && i <= 26 {
			c.N = i
			c.P = i - 13
			c.C = "B"
		} else if i >= 27 && i <= 39 {
			c.N = i
			c.P = i - 26
			c.C = "C"
		} else if i >= 40 && i <= 52 {
			c.N = i
			c.P = i - 39
			c.C = "D"
		} else {
			if i == 53 {
				c.N = 53
				c.P = 53
				c.C = "J"
			}
			if i == 54 {
				c.N = 54
				c.P = 54
				c.C = "J"
			}

		}
		cMap[c.N] = c

	}
	return cMap
}

func makeLoserCards(ex [5]int) map[int]Card {
	var (
		cMap     = make(map[int]Card)
		j        = 0
		remainIx = 1
	)

	for i := 1; i <= 54; i++ {

		var (
			c Card
		)
		if i >= 1 && i <= 13 {
			// 如果點數等於排除者，則跳過
			if j < 5 && i == ex[j] {
				j++
				continue
			}

			c.N = i
			c.P = i
			c.C = "A"
		} else if i >= 14 && i <= 26 {
			// 如果點數等於排除者，則跳過
			if j < 5 && i-13 == ex[j] {
				j++
				continue
			}

			c.N = i
			c.P = i - 13
			c.C = "B"
		} else if i >= 27 && i <= 39 {
			// 如果點數等於排除者，則跳過
			if j < 5 && i-26 == ex[j] {
				j++
				continue
			}

			c.N = i
			c.P = i - 26
			c.C = "C"
		} else if i >= 40 && i <= 52 {
			// 如果點數等於排除者，則跳過
			if j < 5 && i-39 == ex[j] {
				j++
				continue
			}

			c.N = i
			c.P = i - 39
			c.C = "D"
		} else {
			if i == 53 {
				// 如果點數等於排除者，則跳過
				if ex[4] == 53 {
					continue
				}
				c.N = 53
				c.P = 53
				c.C = "J"
			}
			if i == 54 {
				if ex[4] == 54 {
					continue
				}
				c.N = 54
				c.P = 54
				c.C = "J"
			}

		}
		cMap[remainIx] = c

		remainIx++
	}
	return cMap

}
