package hcaptcha

import "github.com/Danny-Dasilva/CycleTLS/cycletls"

func (h *HcaptchaSession) CycleOptions(Body string) cycletls.Options {
	return cycletls.Options{
		Body:      Body,
		Headers:   h.Headers,
		UserAgent: h.Headers["user-agent"],

		// TLSVersion,Ciphers,Extensions,EllipticCurves,EllipticCurvePointFormats
		Ja3:   "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0",
		Proxy: "http://" + h.Proxy,
	}
}
