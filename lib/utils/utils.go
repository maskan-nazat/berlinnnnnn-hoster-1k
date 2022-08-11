package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func HandleError(err error) {
	if err != nil {
		Log(err.Error(), 4)
		return
	}
}

func URLEncode(sitekey string) string {
	var b strings.Builder
	b.Grow(len(sitekey) * 3)

	for _, c := range sitekey {
		if rand.Intn(3) == 1 {
			b.WriteString(fmt.Sprintf("%%%02x", c))
		} else {
			b.WriteRune(c)
		}
	}

	return b.String()
}

func Elapsed(TaskId string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("[Task.%s] took %v\n", TaskId, time.Since(start))
	}
}
