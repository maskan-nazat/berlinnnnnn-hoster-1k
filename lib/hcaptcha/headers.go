package hcaptcha

import "fmt"

func (h *HcaptchaSession) HeaderCheckSiteConfig() map[string]string {
	//`content-length`: `0`,

	return map[string]string{
		`authority`:          `hcaptcha.com`,
		`accept`:             `application/json`,
		`accept-language`:    `fr-FR,fr;q=0.9`,
		`content-type`:       `text/plain`,
		`origin`:             `https://newassets.hcaptcha.com`,
		`referer`:            `https://newassets.hcaptcha.com/`,
		`sec-ch-ua`:          `"Chromium";v="104", " Not A;Brand";v="99", "Google Chrome";v="104"`,
		`sec-ch-ua-mobile`:   `?0`,
		`sec-ch-ua-platform`: `"Windows"`,
		`sec-fetch-dest`:     `empty`,
		`sec-fetch-mode`:     `cors`,
		`sec-fetch-site`:     `same-site`,
		`user-agent`:         `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36`,
	}
}

func (h *HcaptchaSession) HeaderGetCaptcha() map[string]string {
	return map[string]string{
		`authority`:          `hcaptcha.com`,
		`accept`:             `application/json`,
		`accept-language`:    `fr-FR,fr;q=0.9`,
		`content-type`:       `application/x-www-form-urlencoded`,
		`origin`:             `https://newassets.hcaptcha.com`,
		`referer`:            `https://newassets.hcaptcha.com/`,
		`sec-ch-ua`:          `"Chromium";v="104", " Not A;Brand";v="99", "Google Chrome";v="104"`,
		`sec-ch-ua-mobile`:   `?0`,
		`sec-ch-ua-platform`: `"Windows"`,
		`sec-fetch-dest`:     `empty`,
		`sec-fetch-mode`:     `cors`,
		`sec-fetch-site`:     `same-site`,
		`user-agent`:         `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36`,
	}
}

func (h *HcaptchaSession) HeaderCheckCaptcha() map[string]string {
	return map[string]string{
		`accept-encoding`: `gzip, deflate, br`,
		`accept-language`: `fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`,
		`content-type`:    `application/json;charset=UTF-8`,

		`origin`:     `https://newassets.hcaptcha.com`,
		`referer`:    fmt.Sprintf(`https://newassets.hcaptcha.com/captcha/v1/%s/static/hcaptcha.html`, h.version),
		`user-agent`: `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36`,
	}
}
