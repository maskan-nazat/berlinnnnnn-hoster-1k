package utils

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/gookit/color"
	"github.com/inancgumus/screen"
)

var (
	debug      = false
	PromptType = map[int]string{
		1: "<fg=02fa44>SUCCESS</>",
		2: "<fg=fc2323>ERROR</>",
		3: " <fg=f0900a>INFO</> ",
		4: "<fg=0af0ab>DEBUG</>",
	}
)

func BlockConsoleStd() {
	Sc := make(chan os.Signal, 1)
	signal.Notify(Sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-Sc
}

func PrintLogo() {
	screen.Clear()
	screen.MoveTopLeft()
}

func Log(Content string, Type int) {
	if Type == 4 && !debug {
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
