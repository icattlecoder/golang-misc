package progressbar

import (
	"bytes"
	"fmt"
	"syscall"
	"unsafe"
)

type ProgressBar struct {
	label string
}

type window struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

/*
* display demo:
* 78%[===============================================================================>                       ]30kb/s,13min
 */
func New(label string) (pbar *ProgressBar) {
	label = "]" + label
	return &ProgressBar{label: label}
}

func getConsoleWidth() (int, error) {
	w := new(window)
	tio := syscall.TIOCGWINSZ
	res, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), uintptr(tio), uintptr(unsafe.Pointer(w)))
	if int(res) == -1 {
		return 0, err
	}
	return int(w.Col), nil
}

func (pbar *ProgressBar) SetVal(p int, label string) {
	label = "]" + label
	w, err := getConsoleWidth()
	if err != nil {
		return
	}
	//only display label,no progress bar
	if w < 10 {
		return
	}
	//len(100%[)=5
	progressLen := w - len(label) - 5
	takeLen := int(float32(progressLen*p) / 100.0)
	fmt.Printf("]\n\033[F\033[J%3d%%[", p)
	if takeLen <= 1 {
		fmt.Print(">")
	} else {
		cp := bytes.Repeat([]byte{'='}, takeLen-2)
		if p == 100 {
			cp = append(cp, '=')
		} else {
			cp = append(cp, '>')
		}
		bp := bytes.Repeat([]byte{' '}, progressLen-takeLen)
		cp = append(cp, bp...)
		fmt.Print(string(cp))
	}
	fmt.Print(label)
	if p == 100 {
		fmt.Print("\n")
	}
}
