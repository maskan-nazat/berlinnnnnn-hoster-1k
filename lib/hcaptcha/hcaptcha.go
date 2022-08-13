package hcaptcha

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/its-vichy/generator/lib/utils"
	"gopkg.in/square/go-jose.v2/jwt"
)

func NewHcaptchaSession(Proxy string) *HcaptchaSession {
	session := &HcaptchaSession{
		Client:  cycletls.Init(),
		version: utils.Params.Solver.HcaptchaVersion,
		SiteKey: SiteKey["EndpointRegister"],
		Domain:  Domain[utils.Params.Solver.HcaptchaEndpoint],
		Host:    "discord.com",
		Lang:    "fr",
		Proxy:   Proxy,
		Type:    utils.Params.Solver.HcaptchaType,
	}

	return session
}

func (h *HcaptchaSession) CheckSiteConfig() SiteConfig {
	h.Headers = h.HeaderCheckSiteConfig()

	response, err := h.Client.Do(fmt.Sprintf("https://%s/checksiteconfig?v=%s&host=%s&sitekey=%s&sc=1&swa=1", h.Domain, h.version, h.Host, h.SiteKey), h.CycleOptions(""), "GET")
	utils.HandleError(err)

	var siteConfig SiteConfig
	json.Unmarshal([]byte(response.Body), &siteConfig)

	return siteConfig
}

func (h *HcaptchaSession) GetProof(Config SiteConfig) string {
	if h.Type == "hsl" {
		tok, err := jwt.ParseSigned(Config.C.Req)
		utils.HandleError(err)

		var claims map[string]interface{}
		err = tok.UnsafeClaimsWithoutVerification(&claims)
		utils.HandleError(err)

		now := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
		now = now[:len(now)-5]
		now = strings.ReplaceAll(now, "-", "")
		now = strings.ReplaceAll(now, ":", "")
		now = strings.ReplaceAll(now, "T", "")

		return strings.Join([]string{"1", fmt.Sprint(int(claims["s"].(float64))), now, claims["d"].(string), "", "1"}, ":")
	}

	fast := false

	if fast {
		p, err := HSWHashProof(Config.C.Req)

		if err != nil {
			fmt.Println(err)
			return ""
		}

		return p
	} else {
		client := http.Client{
			Timeout: 5 * time.Second,
		}

		for {
			resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:3030/n?req=%s", Config.C.Req))

			if err != nil {
				continue
			}
			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}

			return string(bodyBytes)
		}
	}
}

func (h *HcaptchaSession) GetCaptcha(Config SiteConfig) (Captcha, error) {
	h.Headers = h.HeaderGetCaptcha()

	payload := url.Values{}
	for name, value := range map[string]string{
		`v`:          h.version,
		`sitekey`:    h.SiteKey,
		`host`:       h.Host,
		`hl`:         h.Lang,
		`motionData`: strings.ReplaceAll(MotionData["GetCaptcha"], MotionTime["GetCaptcha"], strconv.Itoa(int(time.Now().Unix()))[:7]),
		`n`:          h.GetProof(Config),
		`c`:          fmt.Sprintf(`{"type":"%s","req":"%s"}`, h.Type, Config.C.Req),
	} {
		payload.Set(name, value)
	}

	// utils.URLEncode(h.SiteKey)
	response, err := h.Client.Do(fmt.Sprintf("https://%s/getcaptcha/%s", h.Domain, h.SiteKey), h.CycleOptions(payload.Encode()), "POST")
	utils.HandleError(err)

	if response.Status != 200 {
		return Captcha{}, errors.New(fmt.Sprintf("(GetCaptcha.1) failed to get captcha with status: %d", response.Status))
	}

	var captcha Captcha
	json.Unmarshal([]byte(response.Body), &captcha)

	// Bypassed (F0_eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ...)
	if captcha.GeneratedPassUUID != "" {
		return captcha, nil
	}

	if _, ok := response.JSONBody()["success"]; ok {
		if !captcha.Success {
			//fmt.Println(response.JSONBody())
			return captcha, errors.New(fmt.Sprintf("(GetCaptcha.2) failed to get captcha with status: %d, error-code: %v", response.Status, captcha.ErrorCode))
		}
	}

	return captcha, nil
}

func (c *Captcha) GetAnswer() map[string]string {
	item := strings.ReplaceAll(strings.Split(strings.ReplaceAll(c.RequesterQuestion.En, "an ", "a "), "Please click each image containing a ")[1], ".", "")
	response := map[string]string{}

	client := http.Client{
		Timeout: 500 * time.Second,
	}

	for _, task := range c.Tasklist {
		resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:1337/check/%s/%s", base64.StdEncoding.EncodeToString([]byte(item)), base64.StdEncoding.EncodeToString([]byte(task.DatapointURI))))

		if err != nil {
			response[task.TaskKey] = "false"
			continue
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			response[task.TaskKey] = "false"
			continue
		}
		bodyString := string(bodyBytes)

		response[task.TaskKey] = strings.ToLower(bodyString)
	}

	return response
}

func (h *HcaptchaSession) CheckCaptcha(Captcha Captcha, Config SiteConfig) (string, error) {
	h.Headers = h.HeaderCheckCaptcha()

	payload, _ := json.Marshal(PayloadGetCaptcha{
		V:            h.version,
		JobMode:      Captcha.RequestType,
		Answers:      Captcha.GetAnswer(),
		Serverdomain: h.Host,
		Sitekey:      h.SiteKey,
		MotionData:   strings.ReplaceAll(MotionData["CheckCaptcha"], MotionTime["CheckCaptcha"], strconv.Itoa(int(time.Now().Unix()))[:7]),
		N:            h.GetProof(Config),
		C:            fmt.Sprintf(`{"type":"%s","req":"%s"}`, h.Type, Captcha.C.Req),
	})

	response, err := h.Client.Do(fmt.Sprintf("https://%s/checkcaptcha/%s/%s", h.Domain, h.SiteKey, Captcha.Key), h.CycleOptions(string(payload)), "POST")
	utils.HandleError(err)

	var captchaResponse CaptchaResponse
	json.Unmarshal([]byte(response.Body), &captchaResponse)

	if !captchaResponse.Pass {
		utils.FailedSubmit++
		fmt.Println(captchaResponse)
		return "", errors.New(fmt.Sprintf("(CheckCaptcha) submit failed: %s", captchaResponse.Error))
	}

	return captchaResponse.GeneratedPassUUID, nil
}

func SolveHcaptcha(Proxy string) (string, error) {
	start := time.Now()

	session := NewHcaptchaSession(Proxy)
	config := session.CheckSiteConfig()

	if !config.Pass {
		return "", errors.New("pass failed")
	}
	captcha, err := session.GetCaptcha(config)
	utils.HandleError(err)

	if (strings.Contains(captcha.RequesterQuestion.En, "lion with mane on its neck") || strings.Contains(captcha.RequesterQuestion.En, "dog with a collar on its neck")) && utils.Params.Solver.HcaptchaEndpoint == "DomainByppass" {
		return "", errors.New("image not supported")
	}

	if captcha.RequesterQuestion.En == "" {
		return "", errors.New("Captcha question is empty")
	}

	if captcha.GeneratedPassUUID != "" {
		fmt.Println("gotcha")
		return captcha.GeneratedPassUUID, nil
	}

	if utils.Params.Advanced.OneClick {
		utils.FailedSubmit++
		return "", errors.New("one-click activated, and captcha not passed")
	}

	key, err := session.CheckCaptcha(captcha, config)

	if err != nil {
		return "", errors.New("CheckCaptcha response rejected")
	}

	utils.Log(fmt.Sprintf("[Hcaptcha] (%s) Took %vs", key[:20], time.Since(start)), 3)

	return key, nil

}
