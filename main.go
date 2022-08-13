package main

import (
	"fmt"
	"time"

	"github.com/its-vichy/generator/lib/discord"
	"github.com/its-vichy/generator/lib/hcaptcha"
	"github.com/its-vichy/generator/lib/utils"
	"github.com/zenthangplus/goccm"
)

func Title() {
	for {
		utils.SetTitle(fmt.Sprintf("Generated: %d Unlocked: %d Locked: %d FailedSubmit: %d", utils.Generated, utils.Unlocked, utils.Locked, utils.FailedSubmit))
		time.Sleep(50 * time.Millisecond)
	}
}

func Worker() {
	proxy := utils.ProxyList.Next()

	captchaKey, err := hcaptcha.SolveHcaptcha(proxy)

	if err != nil {
		utils.HandleError(err)
		return
	}

	session := discord.NewDiscordSession(proxy)
	session.Register(captchaKey)
}

func main() {
	go Title()
	utils.PrintLogo()
	utils.InitFiles()

	fmt.Printf("+ | Running GoToken using %d threads\n", utils.Params.Advanced.Threads)

	c := goccm.New(utils.Params.Advanced.Threads)

	for {
		c.Wait()

		go func() {
			Worker()
			c.Done()
		}()
	}
}
