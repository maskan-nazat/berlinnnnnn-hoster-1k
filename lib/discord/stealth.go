package discord

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

func (d *DiscordSession) CycleOptions(Body string) cycletls.Options {
	return cycletls.Options{
		Body:      Body,
		Headers:   d.Headers,
		UserAgent: d.Headers["user-agent"],
		Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0",
		Proxy:     "http://" + d.Proxy,
	}
}

func (d *DiscordSession) GrabCookiesAndFingerprint() {
	for {
		resp, err := d.Client.Do("https://discord.com/api/v9/experiments", d.CycleOptions(""), "GET")

		if err != nil {
			continue
		}

		var experiments Experiments
		json.Unmarshal([]byte(resp.Body), &experiments)

		cookies := resp.Headers["Set-Cookie"]
		if !strings.Contains(cookies, "__dcfduid") {
			time.Sleep(1 * time.Second)
			continue
		}

		d.Headers["cookie"] = fmt.Sprintf("__dcfduid=%s; __sdcfduid=%s; locale=fr-FR", strings.Split(strings.Split(cookies, "__dcfduid=")[1], ";")[0], strings.Split(strings.Split(cookies, "__sdcfduid=")[1], ";")[0])
		d.Headers["x-fingerprint"] = experiments.Fingerprint
		break
	}
}

func (d *DiscordSession) GrabProperties(IsXTrack bool) {
	//BuildNumber := 140791
	HeaderName := "x-super-properties"

	if IsXTrack {
		//BuildNumber = 9999
		HeaderName = "x-track"
	} else {
		d.Headers["x-discord-locale"] = "fr-FR"
		d.Headers["x-debug-options"] = "bugReporterEnabled"
	}

	/*
			payload = {
		            "os": "Windows",
		            "browser": "Discord Client",
		            "release_channel": "canary",
		            "client_version": "1.0.48",
		            "os_version": "10.0.22000",
		            "os_arch": "x64",
		            "system_locale": "en-US",
		            "client_build_number": __config__['discord']['build_number'],
		            "client_event_source": None
		        }
	*/

	payload, _ := json.Marshal(XProperties{
		Os:                "Windows",
		Browser:           "Discord Client",
		ReleaseChannel:    "canary",
		ClientVersion:     "1.0.48",
		OsVersion:         "10.0.19044",
		OsArch:            "x64",
		SystemLocale:      "fr-FR",
		ClientBuildNumber: 140903,
		ClientEventSource: nil,
	})

	/*payload, _ := json.Marshal(XProperties{
		Os:                     "Windows",
		Browser:                "Chrome",
		Device:                 "",
		SystemLocale:           "fr-FR",
		BrowserUserAgent:       d.Headers["user-agent"],
		BrowserVersion:         strings.Split(strings.Split(d.Headers["user-agent"], "Chrome/")[1], " ")[0],
		OsVersion:              "10",
		Referrer:               "",
		ReferringDomain:        "",
		ReferrerCurrent:        "",
		ReferringDomainCurrent: "",
		ReleaseChannel:         "stable",
		ClientBuildNumber:      BuildNumber,
		ClientEventSource:      nil,
	})*/

	d.Headers[HeaderName] = base64.StdEncoding.EncodeToString(payload)
}
