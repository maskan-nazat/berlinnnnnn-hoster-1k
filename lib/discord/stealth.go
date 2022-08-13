package discord

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

func (d *DiscordSession) CycleOptions(Body string) cycletls.Options {
	return cycletls.Options{
		Body:    Body,
		Headers: d.Headers,
		Ja3:     "771,4865-4866-4867-49196-49195-52393-49200-49199-52392-49162-49161-49172-49171-157-156-53-47-49160-49170-10,0-23-65281-10-11-16-5-13-18-51-45-43-27,29-23-24-25,0",
		Proxy:   "http://" + d.Proxy,
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

		d.Headers["cookie"] = fmt.Sprintf("__dcfduid=%s; __sdcfduid=%s; locale=fr", strings.Split(strings.Split(cookies, "__dcfduid=")[1], ";")[0], strings.Split(strings.Split(cookies, "__sdcfduid=")[1], ";")[0])
		d.Headers["x-fingerprint"] = experiments.Fingerprint
		break
	}
}

func (d *DiscordSession) GrabProperties(IsXTrack bool) {
	BuildNumber := 141471
	HeaderName := "x-super-properties"

	if IsXTrack {
		BuildNumber = 9999
		HeaderName = "x-track"
	} else {
		d.Headers["x-discord-locale"] = "fr-FR"
		d.Headers["x-debug-options"] = "bugReporterEnabled"
	}

	payload, _ := json.Marshal(XProperties{
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
	})

	d.Headers[HeaderName] = base64.StdEncoding.EncodeToString(payload)
}

func (d *DiscordSession) GetCypher() *tls.Config {
	return &tls.Config{
		MaxVersion: tls.VersionTLS13,
		CipherSuites: []uint16{
			tls.TLS_AES_128_GCM_SHA256, tls.TLS_AES_256_GCM_SHA384, tls.TLS_CHACHA20_POLY1305_SHA256, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256, tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, tls.TLS_RSA_WITH_AES_256_GCM_SHA384, tls.TLS_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_RSA_WITH_AES_256_CBC_SHA, tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA, tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		},
		InsecureSkipVerify: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveID(0x0a0a),
			tls.X25519,
			tls.CurveP256,
			tls.CurveP384,
		},
	}
}
