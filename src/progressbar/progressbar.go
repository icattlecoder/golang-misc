package progressbar

import (
	"fmt"
)

type ProgressBar struct {
	label string
}

func New(label string) (pbar *ProgressBar) {
	label = "]" + label
	return &ProgressBar{label: label}
}

func (pbar *ProgressBar) SetVal(p int) {
	fmt.Printf("]\n\033[F\033[J%3d%%[", p)
	for i := 0; i < p; i++ {
		fmt.Print("#")
	}
	for j := p; j < 100; j++ {
		fmt.Print(" ")
	}
	fmt.Print(pbar.label)
	if p == 100 {
		fmt.Print("\n")
	}
}
