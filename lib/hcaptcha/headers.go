package hcaptcha

import "fmt"

func (h *HcaptchaSession) HeaderCheckSiteConfig() map[string]string {
	//`content-length`: `0`,

	return map[string]string{
		`authority`:       h.Domain,
		`accept`:          `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`,
		`accept-language`: `fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`,
		`cache-control`:   `max-age=0`,
		//# Requests sorts cookies= alphabetically
		//# `cookie`: `__cflb=0H28vk2VKwPbLoawEveeCyPgtaDgEZWxdS8fUuutvqz; session=439a5fa3-9cb6-44a5-a3cd-344453b28cce; INGRESSCOOKIE=1660207380.413.40.700977|8ad9e52e7227a1781f7d2cc2db535011; hmt_id=1d9ef33d-8e2f-4121-bd3f-9e6d905f1183`,
		`sec-ch-ua`:                 `"Chromium";v="104", " Not A;Brand";v="99", "Google Chrome";v="104"`,
		`sec-ch-ua-mobile`:          `?0`,
		`sec-ch-ua-platform`:        `"Windows"`,
		`sec-fetch-dest`:            `document`,
		`sec-fetch-mode`:            `navigate`,
		`sec-fetch-site`:            `none`,
		`sec-fetch-user`:            `?1`,
		`upgrade-insecure-requests`: `1`,
		`user-agent`:                `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36`,
	}
}

func (h *HcaptchaSession) HeaderGetCaptcha() map[string]string {
	return map[string]string{
		//`authority`:          `hcaptcha.com`,
		`accept`:             `application/json`,
		`accept-language`:    `fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`,
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
		`authority`:       `hcaptcha.com`,
		`accept-encoding`: `gzip, deflate, br`,
		`accept-language`: `fr-FR,fr;q=0.9`,
		`content-type`:    `application/json;charset=UTF-8`,

		`origin`:     `https://newassets.hcaptcha.com`,
		`referer`:    fmt.Sprintf(`https://newassets.hcaptcha.com/captcha/v1/%s/static/hcaptcha.html`, h.version),
		`user-agent`: `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36`,
	}
}
