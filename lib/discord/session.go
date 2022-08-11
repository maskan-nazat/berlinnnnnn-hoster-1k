package discord

import (
	"encoding/json"
	"fmt"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/its-vichy/generator/lib/utils"
)

func NewDiscordSession(Proxy string) *DiscordSession {
	session := &DiscordSession{
		Client: cycletls.Init(),
		Proxy:  Proxy,
	}

	session.Headers = session.HeaderRegister()
	session.GrabCookiesAndFingerprint()
	session.GrabProperties(true)

	return session
}
func (d *DiscordSession) Register(CaptchaKey string) string {
	payload, _ := json.Marshal(PayloadRegister{
		Consent:     true,
		Fingerprint: d.Headers["x-fingerprint"],
		Username:    utils.UsernameList.Next(),
		CaptchaKey:  CaptchaKey,
		Invite:      "book-club",
	})

	response, err := d.Client.Do("https://canary.discord.com/api/v9/auth/register", d.CycleOptions(string(payload)), "POST")
	utils.HandleError(err)

	var data ResponseRegister
	json.Unmarshal([]byte(response.Body), &data)

	if data.Token != "" {
		//fmt.Println(data.Token)
		utils.AppendLine("./data/output/generated.txt", data.Token)
		go d.CheckAccount(data.Token, CaptchaKey[:2])
	}

	go utils.ProxyList.LockByTimeout(d.Proxy, 125)
	return data.Token
}

func (d *DiscordSession) CheckAccount(Token string, x string) bool {
	utils.Generated++

	delete(d.Headers, "x-super-properties")
	d.GrabProperties(false)
	d.Headers["authorization"] = Token

	response, err := d.Client.Do("https://discord.com/api/v9/users/@me/library", d.CycleOptions(""), "GET")
	utils.HandleError(err)

	if response.Status == 200 {
		utils.Log(fmt.Sprintf("UNLOCKED (%s): %s", x, Token), 1)
		utils.AppendLine("./data/output/unlocked.txt", Token)
		utils.Unlocked++
		return true
	} else {
		utils.Log(fmt.Sprintf("LOCKED (%s): %s", x, Token), 2)
		utils.AppendLine("./data/output/locked.txt", Token)
		utils.Locked++
		return false
	}
}
