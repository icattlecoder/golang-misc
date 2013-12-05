package main

import (
	"progressbar"
	"time"
)

func main() {
	pbar := progressbar.New("上传中")
	for i := 1; i <= 100; i++ {
		pbar.SetVal(i)
		time.Sleep(1e9 * 0.03)
	}
}
