package main

import (
	"fmt"
	"sync"
	"time"
)

// 組合(不分順序)
func Combination(p int) [][]int {

	var bb = &BallPick{}

	s := time.Now()
	bb.getBall5(p)

	// fmt.Println(bb.result)
	fmt.Println()
	fmt.Println("花費時間:", time.Since(s).Seconds(), " 秒")

	// 取7數量: 85900584
	// 花費時間: 39.971854635  秒
	return bb.result
}

// BallPick 取5顆球
type BallPick struct {
	wg     *sync.WaitGroup
	locker *sync.RWMutex
	result [][]int
}

func (b *BallPick) setResult(pos [][]int) {
	b.locker.Lock()
	b.result = append(b.result, pos...)
	b.locker.Unlock()
}

func (b *BallPick) getBall5(maxBall int) {

	b.result = [][]int{}
	b.wg = new(sync.WaitGroup)
	b.locker = new(sync.RWMutex)

	var n = 5
	for i := 1; i <= (maxBall - n + 1); i++ {
		b.wg.Add(1)
		go b.getBall5sub(i, maxBall)
	}
	b.wg.Wait()

	// fmt.Println(b.result)
	fmt.Println("取5數量:", len(b.result))
}

func (b *BallPick) getBall5sub(p1, maxBall int) {

	var (
		position [][]int
		m        = maxBall
	)
	// 取5
	for p2 := p1 + 1; p2 <= m; p2++ {
		for p3 := p2 + 1; p3 <= m; p3++ {
			for p4 := p3 + 1; p4 <= m; p4++ {
				for p5 := p4 + 1; p5 <= m; p5++ {
					position = append(position, []int{p5, p4, p3, p2, p1})
				}
			}
		}
	}

	b.setResult(position)
	b.wg.Done()

}
