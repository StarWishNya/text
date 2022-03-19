package main

import "syscall"

var (
	kernel32    *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
	proc        *syscall.LazyProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle *syscall.LazyProc = kernel32.NewProc(`CloseHandle`)

	字体颜色 前景色 = 前景色{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

type 前景色 struct {
	黑色  int
	蓝色  int
	绿色  int
	青色  int
	红色  int
	紫色  int
	黄色  int
	淡灰色 int
	灰色  int
	亮蓝色 int
	亮绿色 int
	亮青色 int
	亮红色 int
	亮紫色 int
	亮黄色 int
	白色  int
}

func colorout(s string, i int) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	print(s)
	CloseHandle.Call(handle)
}

func colorreset() {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle.Call(handle)
}
