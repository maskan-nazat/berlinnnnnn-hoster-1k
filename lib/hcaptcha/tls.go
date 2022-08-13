package hcaptcha

import (
	//"fmt"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/its-vichy/GoCycle"
)

var (
	j_cycle = GoCycle.New(&[]string{"771,4865-4866-4867-49196-49195-52393-49200-49199-52392-49162-49161-49172-49171-157-156-53-47-49160-49170-10,0-23-65281-10-11-16-5-13-18-51-45-43-27,29-23-24-25,0", "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0"})
)

func (h *HcaptchaSession) CycleOptions(Body string) cycletls.Options {
	/*h.Headers["content-length"] = fmt.Sprintf("%d", len(Body))
	defer func() {
		delete(h.Headers, "content-length")
	}()*/

	return cycletls.Options{
		Body:      Body,
		Headers:   h.Headers,
		UserAgent: h.Headers["user-agent"],

		// TLSVersion,Ciphers,Extensions,EllipticCurves,EllipticCurvePointFormats

		// Chrome: 771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0
		// Phone: 771,4865-4866-4867-49196-49195-52393-49200-49199-52392-49162-49161-49172-49171-157-156-53-47-49160-49170-10,0-23-65281-10-11-16-5-13-18-51-45-43-27,29-23-24-25,0
		Ja3:   "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0", //j_cycle.Next(),
		Proxy: "http://" + h.Proxy,
	}
}
