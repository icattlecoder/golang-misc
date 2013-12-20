package main

import (
	"progressbar"
	"time"
)

func main() {

	pbar := progressbar.New("上传中")
	for i := 1; i <= 100; i++ {
		pbar.SetVal(i, "32Kb/s,32min")
		time.Sleep(1e9 * 0.09)
	}
}
