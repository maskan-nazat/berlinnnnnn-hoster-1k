package utils

import (
	"github.com/BurntSushi/toml"
	"github.com/its-vichy/GoCycle"
)

var (
	FailedSubmit int = 0
	Generated    int = 0
	Unlocked     int = 0
	Locked       int = 0

	ProxyList    *GoCycle.Cycle
	UsernameList *GoCycle.Cycle

	Params Config
)

type Config struct {
	Solver struct {
		HcaptchaEndpoint string `toml:"hcaptchaEndpoint"`
		HcaptchaVersion  string `toml:"hcaptchaVersion"`
		HcaptchaType     string `toml:"hcaptchaType"`
	} `toml:"solver"`
	Advanced struct {
		Debug bool `toml:"debug"`
	} `toml:"advanced"`
}

func InitFiles() {
	ProxyList, _ = GoCycle.NewFromFile("data/input/proxies.txt")
	UsernameList, _ = GoCycle.NewFromFile("data/input/usernames.txt")

	toml.DecodeFile("data/input/config.toml", &Params)
}
