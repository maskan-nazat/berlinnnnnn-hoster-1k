package hcaptcha

import "github.com/Danny-Dasilva/CycleTLS/cycletls"

var (
	SiteKey = map[string]string{
		"EndpointRegister": "4c672d35-0701-42b2-88c3-78380b0db560",
		"EndpointVerify":   "f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34",
	}

	Domain = map[string]string{
		"DomainOfficial": "hcaptcha.com",
		"DomainByppass":  "staging.hmt.ai",
	}

	MotionData = map[string]string{
		"GetCaptcha":   `{"st":1660262869224,"mm":[[22,77,1660262869476],[17,60,1660262869492],[17,47,1660262869509],[18,40,1660262869526],[21,36,1660262869542],[23,35,1660262869562]],"mm-mp":2.15,"md":[[22,36,1660262869553]],"md-mp":0,"mu":[[23,35,1660262869621]],"mu-mp":0,"v":1,"topLevel":{"st":1660262869095,"sc":{"availWidth":1920,"availHeight":1040,"width":1920,"height":1080,"colorDepth":24,"pixelDepth":24,"availLeft":0,"availTop":0,"onchange":null,"isExtended":false},"nv":{"vendorSub":"","productSub":"20030107","vendor":"Google Inc.","maxTouchPoints":0,"scheduling":{},"userActivation":{},"doNotTrack":null,"geolocation":{},"connection":{},"pdfViewerEnabled":true,"webkitTemporaryStorage":{},"webkitPersistentStorage":{},"hardwareConcurrency":16,"cookieEnabled":true,"appCodeName":"Mozilla","appName":"Netscape","appVersion":"5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36","platform":"Win32","product":"Gecko","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36","language":"fr-FR","languages":["fr-FR","fr","en-US","en"],"onLine":true,"webdriver":false,"bluetooth":{},"clipboard":{},"credentials":{},"keyboard":{},"managed":{},"mediaDevices":{},"storage":{},"serviceWorker":{},"wakeLock":{},"deviceMemory":8,"ink":{},"hid":{},"locks":{},"mediaCapabilities":{},"mediaSession":{},"permissions":{},"presentation":{},"serial":{},"virtualKeyboard":{},"usb":{},"xr":{},"userAgentData":{"brands":[{"brand":"Chromium","version":"104"},{"brand":" Not A;Brand","version":"99"},{"brand":"Google Chrome","version":"104"}],"mobile":false,"platform":"Windows"},"plugins":["internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer"]},"dr":"","inv":false,"exec":false,"wn":[[393,969,1,1660262869096]],"wn-mp":0,"xy":[[0,0,1,1660262869096]],"xy-mp":0,"mm":[[322,541,1660262869097],[308,567,1660262869113],[289,582,1660262869129],[270,590,1660262869145],[252,593,1660262869161],[233,593,1660262869178],[216,592,1660262869194],[203,589,1660262869211],[193,585,1660262869228],[187,582,1660262869246],[184,579,1660262869264],[182,577,1660262869280],[181,574,1660262869296],[178,571,1660262869315],[174,568,1660262869331],[169,563,1660262869348],[163,558,1660262869364],[153,550,1660262869380],[138,540,1660262869396],[122,529,1660262869412],[104,517,1660262869428],[88,504,1660262869444],[76,489,1660262869460],[67,471,1660262869476]],"mm-mp":1.5220883534136547},"session":[],"widgetList":["0j2vo3r8mvve"],"widgetId":"0j2vo3r8mvve","href":"https://discord.com/","prev":{"escaped":false,"passed":false,"expiredChallenge":false,"expiredResponse":false}}`,
		"CheckCaptcha": `{"st":1660262870944,"dct":1660262870944,"mm":[[196,243,1660262871212],[199,240,1660262871228],[207,237,1660262871247],[217,231,1660262871263],[229,225,1660262871279],[243,216,1660262871295],[254,209,1660262871311],[263,203,1660262871327],[270,198,1660262871343],[273,194,1660262871359],[274,192,1660262871495],[270,190,1660262871512],[255,187,1660262871528],[223,183,1660262871544],[174,178,1660262871560],[112,171,1660262871577],[38,161,1660262871593],[0,136,1660262871932],[18,135,1660262871948],[34,134,1660262871965],[51,132,1660262871981],[64,130,1660262871997],[72,127,1660262872015],[75,125,1660262872032],[76,124,1660262872063],[76,123,1660262872101],[76,121,1660262872122],[76,119,1660262872146],[76,117,1660262872163],[76,115,1660262872181],[74,114,1660262872198],[69,112,1660262872215],[58,112,1660262872231],[42,113,1660262872247],[26,116,1660262872263],[14,122,1660262872279],[5,131,1660262872295],[1,138,1660262872312],[1,145,1660262872331],[4,153,1660262872349],[11,161,1660262872365],[25,170,1660262872382],[43,177,1660262872398],[63,180,1660262872414],[82,182,1660262872430],[101,184,1660262872447],[112,184,1660262872463],[116,182,1660262872547],[110,180,1660262872564],[98,176,1660262872580],[82,170,1660262872596],[62,163,1660262872612],[46,156,1660262872628],[37,153,1660262872644],[33,152,1660262872795],[34,157,1660262872811],[35,167,1660262872828],[36,176,1660262872847],[36,183,1660262872864],[37,192,1660262872880],[39,204,1660262872896],[42,220,1660262872912],[45,234,1660262872928],[49,243,1660262872944],[52,247,1660262872963],[53,249,1660262873061],[57,251,1660262873078],[63,254,1660262873094],[71,257,1660262873110],[86,263,1660262873126],[108,274,1660262873142],[129,285,1660262873158],[148,297,1660262873174],[161,307,1660262873191],[171,316,1660262873207],[175,322,1660262873223],[178,328,1660262873242],[178,332,1660262873260],[176,336,1660262873276],[171,341,1660262873294],[165,347,1660262873311],[160,349,1660262873327],[159,351,1660262873344],[159,353,1660262873456],[164,354,1660262873473],[173,356,1660262873489],[187,356,1660262873505],[204,356,1660262873521],[222,355,1660262873537],[238,352,1660262873553],[252,350,1660262873569],[262,350,1660262873585],[269,348,1660262873601],[274,346,1660262873617],[278,341,1660262873635],[281,334,1660262873651],[282,323,1660262873669],[277,310,1660262873685],[272,297,1660262873702],[269,287,1660262873719],[267,280,1660262873736],[265,274,1660262873752],[263,269,1660262873768],[261,264,1660262873784],[256,258,1660262873800],[249,251,1660262873817],[238,245,1660262873835],[224,241,1660262873852],[205,238,1660262873868],[183,235,1660262873884],[155,234,1660262873900],[130,232,1660262873916],[113,232,1660262873934],[103,232,1660262873953],[103,232,1660262874007],[109,232,1660262874025],[119,232,1660262874042],[132,232,1660262874059],[146,232,1660262874076],[158,233,1660262874092],[168,234,1660262874108],[173,235,1660262874130],[173,237,1660262874171],[173,239,1660262874189],[173,243,1660262874207],[173,247,1660262874224],[173,251,1660262874242],[173,253,1660262874266],[173,254,1660262874350],[173,257,1660262874366],[173,263,1660262874382],[178,270,1660262874398],[187,277,1660262874414],[202,284,1660262874430],[219,286,1660262874446],[240,287,1660262874462],[257,287,1660262874479],[268,286,1660262874495],[274,282,1660262874515],[278,279,1660262874535],[280,275,1660262874554],[280,268,1660262874571],[280,257,1660262874587],[281,241,1660262874603],[282,220,1660262874619],[284,207,1660262874635],[286,198,1660262874654],[286,197,1660262874791],[280,197,1660262874807],[271,198,1660262874823],[256,201,1660262874839],[235,208,1660262874855],[205,217,1660262874871],[169,229,1660262874887],[136,240,1660262874903],[110,249,1660262874919],[90,256,1660262874936],[82,259,1660262874956],[82,261,1660262874985],[82,262,1660262875004],[83,265,1660262875020],[87,269,1660262875036],[95,274,1660262875052],[107,280,1660262875068],[127,290,1660262875084],[155,302,1660262875101],[182,315,1660262875118],[204,327,1660262875134],[218,337,1660262875151],[225,345,1660262875171],[228,349,1660262875204],[228,350,1660262875222],[228,353,1660262875243],[228,356,1660262875262],[228,360,1660262875279],[229,366,1660262875298],[230,372,1660262875316],[233,380,1660262875334],[236,390,1660262875351],[240,401,1660262875368],[246,412,1660262875384],[251,422,1660262875400],[258,429,1660262875416],[264,435,1660262875435],[271,438,1660262875451],[277,441,1660262875468],[282,442,1660262875484],[286,443,1660262875505]],"mm-mp":2.3836757357023863,"md":[[33,151,1660262872729],[53,248,1660262873001],[159,352,1660262873354],[173,253,1660262874289],[286,443,1660262875609]],"md-mp":720,"mu":[[33,151,1660262872776],[55,250,1660262873071],[159,352,1660262873441],[173,254,1660262874350],[286,443,1660262875704]],"mu-mp":732,"topLevel":{"st":1660262869095,"sc":{"availWidth":1920,"availHeight":1040,"width":1920,"height":1080,"colorDepth":24,"pixelDepth":24,"availLeft":0,"availTop":0,"onchange":null,"isExtended":false},"nv":{"vendorSub":"","productSub":"20030107","vendor":"Google Inc.","maxTouchPoints":0,"scheduling":{},"userActivation":{},"doNotTrack":null,"geolocation":{},"connection":{},"pdfViewerEnabled":true,"webkitTemporaryStorage":{},"webkitPersistentStorage":{},"hardwareConcurrency":16,"cookieEnabled":true,"appCodeName":"Mozilla","appName":"Netscape","appVersion":"5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36","platform":"Win32","product":"Gecko","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36","language":"fr-FR","languages":["fr-FR","fr","en-US","en"],"onLine":true,"webdriver":false,"bluetooth":{},"clipboard":{},"credentials":{},"keyboard":{},"managed":{},"mediaDevices":{},"storage":{},"serviceWorker":{},"wakeLock":{},"deviceMemory":8,"ink":{},"hid":{},"locks":{},"mediaCapabilities":{},"mediaSession":{},"permissions":{},"presentation":{},"serial":{},"virtualKeyboard":{},"usb":{},"xr":{},"userAgentData":{"brands":[{"brand":"Chromium","version":"104"},{"brand":" Not A;Brand","version":"99"},{"brand":"Google Chrome","version":"104"}],"mobile":false,"platform":"Windows"},"plugins":["internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer"]},"dr":"","inv":false,"exec":false,"wn":[[393,969,1,1660262869096]],"wn-mp":0,"xy":[[0,0,1,1660262869096]],"xy-mp":0,"mm":[[322,541,1660262869097],[308,567,1660262869113],[289,582,1660262869129],[270,590,1660262869145],[252,593,1660262869161],[233,593,1660262869178],[216,592,1660262869194],[203,589,1660262869211],[193,585,1660262869228],[187,582,1660262869246],[184,579,1660262869264],[182,577,1660262869280],[181,574,1660262869296],[178,571,1660262869315],[174,568,1660262869331],[169,563,1660262869348],[163,558,1660262869364],[153,550,1660262869380],[138,540,1660262869396],[122,529,1660262869412],[104,517,1660262869428],[88,504,1660262869444],[76,489,1660262869460],[67,471,1660262869476],[54,473,1660262869769],[52,506,1660262869785],[53,544,1660262869801],[59,579,1660262869817],[65,602,1660262869833],[70,615,1660262869849],[74,622,1660262869866],[76,623,1660262870613],[82,621,1660262870631],[93,615,1660262870647],[113,603,1660262870663],[139,587,1660262870679],[166,568,1660262870695],[191,547,1660262870711],[210,527,1660262870727],[222,508,1660262870743],[229,496,1660262870759],[233,490,1660262870778],[234,489,1660262871212],[33,400,1660262871604],[0,392,1660262871620],[0,386,1660262871636],[0,384,1660262871682],[3,383,1660262871700],[4,383,1660262871871],[8,383,1660262871888],[16,383,1660262871905],[28,382,1660262871921]],"mm-mp":5.328947368421053},"v":1}`,
	}

	MotionTime = map[string]string{
		"GetCaptcha":   "1660262",
		"CheckCaptcha": "1660262",
	}
)

type HcaptchaSession struct {
	Client  cycletls.CycleTLS
	Headers map[string]string
	version string
	SiteKey string
	Domain  string
	Host    string
	Lang    string
	Proxy   string
	Type    string
}

type SiteConfig struct {
	Features struct {
	} `json:"features"`
	C struct {
		Type string `json:"type"`
		Req  string `json:"req"`
	} `json:"c"`
	Pass bool `json:"pass"`
}

type Captcha struct {
	C struct {
		Type string `json:"type"`
		Req  string `json:"req"`
	} `json:"c"`
	ChallengeURI  string `json:"challenge_uri"`
	Key           string `json:"key"`
	RequestConfig struct {
		MaxPoints                    interface{} `json:"max_points"`
		MaxShapesPerImage            interface{} `json:"max_shapes_per_image"`
		MinPoints                    interface{} `json:"min_points"`
		MinShapesPerImage            interface{} `json:"min_shapes_per_image"`
		MinimumSelectionAreaPerShape interface{} `json:"minimum_selection_area_per_shape"`
		MultipleChoiceMaxChoices     int         `json:"multiple_choice_max_choices"`
		MultipleChoiceMinChoices     int         `json:"multiple_choice_min_choices"`
		RestrictToCoords             interface{} `json:"restrict_to_coords"`
		ShapeType                    interface{} `json:"shape_type"`
		Version                      int         `json:"version"`
	} `json:"request_config"`
	RequestType       string `json:"request_type"`
	RequesterQuestion struct {
		Fr string `json:"fr"`
		En string `json:"en"`
	} `json:"requester_question"`
	RequesterQuestionExample     []string `json:"requester_question_example"`
	RequesterRestrictedAnswerSet struct {
		Num0 struct {
		} `json:"0"`
		Num1 struct {
		} `json:"1"`
	} `json:"requester_restricted_answer_set"`
	Tasklist []struct {
		DatapointHash string `json:"datapoint_hash"`
		DatapointURI  string `json:"datapoint_uri"`
		TaskKey       string `json:"task_key"`
	} `json:"tasklist"`
	BypassMessage string `json:"bypass-message"`

	GeneratedPassUUID string   `json:"generated_pass_UUID"`
	Success           bool     `json:"success"`
	ErrorCode         []string `json:"error-codes"`
}

type CaptchaResponse struct {
	C struct {
		Type string `json:"type"`
		Req  string `json:"req"`
	} `json:"c"`
	Pass              bool   `json:"pass"`
	GeneratedPassUUID string `json:"generated_pass_UUID"`
	Expiration        int    `json:"expiration"`
	Error             string `json:"error"`
}

type PayloadGetCaptcha struct {
	V            string            `json:"v"`
	JobMode      string            `json:"job_mode"`
	Answers      map[string]string `json:"answers"`
	Serverdomain string            `json:"serverdomain"`
	Sitekey      string            `json:"sitekey"`
	MotionData   string            `json:"motionData"`
	N            string            `json:"n"`
	C            string            `json:"c"`
}
