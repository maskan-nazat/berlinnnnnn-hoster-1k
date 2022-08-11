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
	})
	utils.HandleError(err)

	page, err = browser.NewPage()
	utils.HandleError(err)

	_, err = page.AddScriptTag(playwright.PageAddScriptTagOptions{Content: playwright.String(GetHSW())})
	utils.HandleError(err)

}

func GetHSW() string {
	resp, err := http.Get("https://newassets.hcaptcha.com/c/f543c527/hsw.js")
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
	resp, err := page.Evaluate(fmt.Sprintf(`hsw("%v")`, request))
	if err != nil {
		return "", err
	}

	return resp.(string), nil
}