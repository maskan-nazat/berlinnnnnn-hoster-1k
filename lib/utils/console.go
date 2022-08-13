package utils

import (
	"fmt"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/gookit/color"
	"github.com/inancgumus/screen"
)

var (
	PromptType = map[int]string{
		1: "<fg=02fa44>+</>",
		2: "<fg=fc2323>-</>",
		3: "<fg=f0900a>*</>",
		4: "<fg=0af0ab>%</>",
	}
)

func PrintLogo() {
	screen.Clear()
	screen.MoveTopLeft()
}

func Log(Content string, Type int) {
	if Type == 4 && !Params.Advanced.Debug {
		return
	}

	date := strings.ReplaceAll(time.Now().Format("15:04:05"), ":", "<fg=353a3b>:</>")
	content := fmt.Sprintf("[%s] [%s] %s.", date, PromptType[Type], Content)

	content = strings.ReplaceAll(content, "[", "<fg=ffffff>[</>")
	content = strings.ReplaceAll(content, "]", "<fg=ffffff>]</>")

	content = strings.ReplaceAll(content, "(", "<fg=6c7273;op=bold>(")
	content = strings.ReplaceAll(content, ")", ")</>")

	color.Println(content)
}

func SetTitle(title string) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return
	}
	defer syscall.FreeLibrary(handle)

	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return
	}

	syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
}
