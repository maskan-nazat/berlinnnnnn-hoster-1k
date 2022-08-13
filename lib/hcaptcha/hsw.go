// credit: my friend dropout, btw full lock so use undetected-browser

package hcaptcha

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/its-vichy/generator/lib/utils"
	playwright "github.com/playwright-community/playwright-go"
)

var page playwright.Page

func init() {
	pw, err := playwright.Run()
	if err != nil {
		err := playwright.Install()
		utils.HandleError(err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		Args: []string{
			`--lang=fr`,
			`--no-sandbox`,
			`--disable-setuid-sandbox`,
			`--disable-infobars`,
			`--single-process`,
			`--no-zygote`,
			`--no-first-run`,
			`--window-position=0,0`,
			`--ignore-certificate-errors`,
			`--ignore-certificate-errors-skip-list`,
			`--disable-dev-shm-usage`,
			`--disable-accelerated-2d-canvas`,
			`--disable-gpu`,
			`--hide-scrollbars`,
			`--disable-notifications`,
			`--disable-background-timer-throttling`,
			`--disable-backgrounding-occluded-windows`,
			`--disable-breakpad`,
			`--disable-component-extensions-with-background-pages`,
			`--disable-extensions`,
			`--disable-features=TranslateUI,BlinkGenPropertyTrees`,
			`--disable-ipc-flooding-protection`,
			`--disable-renderer-backgrounding`,
			`--enable-features=NetworkService,NetworkServiceInProcess`,
			`--force-color-profile=srgb`,
			`--metrics-recording-only`,
			`--mute-audio`,
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
		},
	})
	utils.HandleError(err)

	page, err = browser.NewPage()
	utils.HandleError(err)

	_, err = page.AddScriptTag(playwright.PageAddScriptTagOptions{Content: playwright.String(GetHSW())})
	utils.HandleError(err)

}

func GetHSW() string {
	resp, err := http.Get("https://newassets.hcaptcha.com/c/c916818a/hsw.js")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func HSWHashProof(request string) (string, error) {
	resp, err := page.Evaluate(fmt.Sprintf(`%s("%v")`, utils.Params.Solver.HcaptchaType, request))
	if err != nil {
		return "", err
	}

	return resp.(string), nil
}
