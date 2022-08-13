package utils

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/its-vichy/GoCycle"
)

var (
	FailedSubmit int = 0
	Generated    int = 0
	Unlocked     int = 0
	Locked       int = 0
	CPM          int = 0

	ProxyList    *GoCycle.Cycle
	UsernameList *GoCycle.Cycle

	Params Config
)

/*
def gen_calculator_thread(self):
        start_time = time.time()

        while True:
            try:
                self._gen_min   = round((self._locked + self._unlocked) / ((time.time() - start_time) / 60))
                self._valid_min = round(self._unlocked / ((time.time() - start_time) / 60))
                self._lock_rate = (100 * self._locked) / (self._locked + self._unlocked)


                if self._solved_captcha_times:
                    self._sucess_rate = (100 * self._solved_succes) / (self._solved_succes + self._solved_fail)
                else:
                    self._sucess_rate = 0
            except:
                pass

            time.sleep(1)
*/

func CounterThread() {
	lasted := 0

	for {
		time.Sleep(1 * time.Second)

		// calculate how many accounts generated in 1 second
		CPM = Generated - lasted
		lasted = Generated
	}
}

type Config struct {
	Solver struct {
		HcaptchaEndpoint string `toml:"hcaptchaEndpoint"`
		HcaptchaVersion  string `toml:"hcaptchaVersion"`
		HcaptchaType     string `toml:"hcaptchaType"`
	} `toml:"solver"`
	Advanced struct {
		Debug    bool `toml:"debug"`
		Threads  int  `toml:"threads"`
		OneClick bool `toml:"oneClick"`
	} `toml:"advanced"`
	Discord struct {
		InviteCode string `toml:"inviteCode"`
	}
}

func InitFiles() {
	ProxyList, _ = GoCycle.NewFromFile("data/input/proxies.txt")
	UsernameList, _ = GoCycle.NewFromFile("data/input/usernames.txt")

	toml.DecodeFile("data/input/config.toml", &Params)
}
