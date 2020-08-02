package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Bar struct {
	percent int    // 百分比
	cur     int64  // 当前进度 实际值
	total   int64  // 总进度 实际值
	rate    string // 进度条字符串
	graph   string // 进度条的符号
}

func (bar *Bar) NewOption(cur, total int64) {
	bar.cur = cur
	bar.total = total
	bar.percent = bar.getPercent()
	if bar.graph == "" {
		bar.graph = "#"
	}
	for i := 0; i <= int(bar.cur); i++ {
		bar.rate += bar.graph
	}
}

func (bar *Bar) NewOptionWithGraph(graph string, cur, total int64) {
	bar.graph = "#"
	bar.NewOption(cur, total)
}

func (bar *Bar) getPercent() int {
	return int(float64(bar.cur) / float64(bar.total) * 100)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Println()
}

func main() {
	bar := &Bar{}
	bar.NewOption(0, 10000)
	for i := 0; i <= 10000; i += 50 {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}
