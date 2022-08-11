package discord

import "github.com/Danny-Dasilva/CycleTLS/cycletls"

type DiscordSession struct {
	Client  cycletls.CycleTLS
	Headers map[string]string
	Proxy   string
}

type XContext struct {
	Location            string  `json:"location"`
	LocationGuildID     string  `json:"location_guild_id"`
	LocationChannelID   string  `json:"location_channel_id"`
	LocationChannelType float64 `json:"location_channel_type"`
}

type XProperties struct {
	Os                     string      `json:"os"`
	Browser                string      `json:"browser"`
	Device                 string      `json:"device"`
	SystemLocale           string      `json:"system_locale"`
	BrowserUserAgent       string      `json:"browser_user_agent"`
	BrowserVersion         string      `json:"browser_version"`
	OsVersion              string      `json:"os_version"`
	Referrer               string      `json:"referrer"`
	ReferringDomain        string      `json:"referring_domain"`
	ReferrerCurrent        string      `json:"referrer_current"`
	ReferringDomainCurrent string      `json:"referring_domain_current"`
	ReleaseChannel         string      `json:"release_channel"`
	ClientBuildNumber      int         `json:"client_build_number"`
	ClientEventSource      interface{} `json:"client_event_source"`

	// client: x64
	OsArch        string `json:"os_arch"`
	ClientVersion string `json:"client_version"`
}

type Experiments struct {
	Fingerprint string  `json:"fingerprint"`
	Assignments [][]int `json:"assignments"`
}

type PayloadRegister struct {
	Consent     bool   `json:"consent"`
	Fingerprint string `json:"fingerprint"`
	Username    string `json:"username"`
	CaptchaKey  string `json:"captcha_key"`
	Invite      string `json:"invite"`
}

type ResponseRegister struct {
	Token string `json:"token"`
}
