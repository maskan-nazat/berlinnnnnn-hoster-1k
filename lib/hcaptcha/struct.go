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
		"GetCaptcha":   `{"st":1660097996719,"mm":[[293,24,1660097996885],[294,27,1660097996904],[295,32,1660097996921],[294,40,1660097996938],[286,49,1660097996954],[273,58,1660097996970],[252,64,1660097996986],[224,67,1660097997002],[186,67,1660097997018],[139,63,1660097997034],[86,55,1660097997050],[43,47,1660097997066],[19,44,1660097997082]],"mm-mp":1.3164556962025322,"md":[[17,43,1660097997105]],"md-mp":0,"mu":[[17,43,1660097997169]],"mu-mp":0,"v":1,"topLevel":{"st":1660097996558,"sc":{"availWidth":1920,"availHeight":1040,"width":1920,"height":1080,"colorDepth":24,"pixelDepth":24,"availLeft":0,"availTop":0,"onchange":null,"isExtended":false},"nv":{"vendorSub":"","productSub":"20030107","vendor":"Google Inc.","maxTouchPoints":0,"scheduling":{},"userActivation":{},"doNotTrack":null,"geolocation":{},"connection":{},"pdfViewerEnabled":true,"webkitTemporaryStorage":{},"webkitPersistentStorage":{},"hardwareConcurrency":16,"cookieEnabled":true,"appCodeName":"Mozilla","appName":"Netscape","appVersion":"5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36","platform":"Win32","product":"Gecko","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36","language":"fr-FR","languages":["fr-FR"],"onLine":true,"webdriver":false,"bluetooth":{},"clipboard":{},"credentials":{},"keyboard":{},"managed":{},"mediaDevices":{},"storage":{},"serviceWorker":{},"wakeLock":{},"deviceMemory":8,"ink":{},"hid":{},"locks":{},"mediaCapabilities":{},"mediaSession":{},"permissions":{},"presentation":{},"serial":{},"virtualKeyboard":{},"usb":{},"xr":{},"userAgentData":{"brands":[{"brand":"Chromium","version":"104"},{"brand":" Not A;Brand","version":"99"},{"brand":"Google Chrome","version":"104"}],"mobile":false,"platform":"Windows"},"plugins":["internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer"]},"dr":"","inv":false,"exec":false,"wn":[[460,969,1,1660097996559]],"wn-mp":0,"xy":[[0,0,1,1660097996559]],"xy-mp":0,"mm":[[372,369,1660097996885]],"mm-mp":0},"session":[],"widgetList":["0lhpa73g7qbm"],"widgetId":"0lhpa73g7qbm","href":"https://discord.com/","prev":{"escaped":false,"passed":false,"expiredChallenge":false,"expiredResponse":false}}`,
		"CheckCaptcha": `{"st":1660097997582,"dct":1660097997582,"mm":[[26,142,1660097998169],[37,142,1660097998185],[62,144,1660097998202],[105,147,1660097998218],[155,151,1660097998234],[213,155,1660097998250],[270,159,1660097998266],[319,200,1660097998601],[304,207,1660097998617],[290,213,1660097998633],[276,219,1660097998650],[265,224,1660097998666],[259,226,1660097998684],[257,227,1660097998701],[257,228,1660097998914],[256,231,1660097998930],[253,237,1660097998950],[246,244,1660097998967],[236,254,1660097998983],[222,263,1660097998999],[209,271,1660097999015],[196,277,1660097999031],[187,281,1660097999048],[182,283,1660097999067],[180,284,1660097999083],[178,284,1660097999110],[174,284,1660097999127],[169,281,1660097999143],[164,277,1660097999161],[160,274,1660097999183],[158,273,1660097999363],[152,277,1660097999380],[143,279,1660097999396],[136,282,1660097999412],[130,283,1660097999430],[126,284,1660097999446],[122,285,1660097999463],[118,285,1660097999483],[114,285,1660097999502],[110,285,1660097999521],[107,284,1660097999545],[106,283,1660097999561],[106,282,1660097999577],[105,280,1660097999593],[106,276,1660097999609],[112,269,1660097999625],[123,259,1660097999641],[137,243,1660097999657],[155,223,1660097999673],[172,203,1660097999689],[183,191,1660097999705],[188,185,1660097999723],[191,183,1660097999740],[191,182,1660097999784],[191,181,1660097999957],[188,178,1660097999975],[184,175,1660097999991],[181,174,1660098000010],[180,173,1660098000111],[175,174,1660098000130],[165,177,1660098000147],[153,182,1660098000163],[138,190,1660098000179],[119,200,1660098000195],[96,213,1660098000211],[73,228,1660098000227],[52,242,1660098000243],[32,254,1660098000259],[17,266,1660098000275],[6,278,1660098000291],[1,285,1660098000307],[0,290,1660098000325],[0,293,1660098000346],[6,296,1660098000363],[16,299,1660098000379],[34,303,1660098000395],[58,307,1660098000411],[89,310,1660098000427],[124,313,1660098000443],[160,315,1660098000459],[198,316,1660098000475],[228,316,1660098000491],[250,316,1660098000507],[262,315,1660098000523],[270,314,1660098000544],[271,315,1660098000764],[271,321,1660098000780],[270,331,1660098000796],[264,341,1660098000812],[256,348,1660098000829],[246,356,1660098000845],[233,362,1660098000862],[220,367,1660098000878],[207,369,1660098000894],[194,370,1660098000910],[179,371,1660098000927],[165,371,1660098000943],[149,368,1660098000959],[129,365,1660098000975],[108,362,1660098000991],[84,359,1660098001007],[56,356,1660098001023],[32,355,1660098001039],[14,355,1660098001056],[9,355,1660098001173],[13,355,1660098001192],[19,355,1660098001208],[25,356,1660098001225],[31,356,1660098001242],[38,357,1660098001260],[48,358,1660098001276],[58,359,1660098001292],[71,360,1660098001308],[91,361,1660098001324],[116,361,1660098001340],[143,362,1660098001356],[167,362,1660098001372],[189,362,1660098001389],[207,362,1660098001406],[219,360,1660098001423],[228,360,1660098001439],[235,360,1660098001456],[241,359,1660098001473],[246,358,1660098001492],[248,358,1660098001510],[247,358,1660098002047],[243,359,1660098002064],[236,361,1660098002080],[226,362,1660098002097],[214,363,1660098002113],[203,364,1660098002130],[195,365,1660098002146],[193,365,1660098002322],[198,365,1660098002338],[207,364,1660098002354],[220,362,1660098002370],[233,359,1660098002386],[247,355,1660098002402],[258,352,1660098002419],[267,351,1660098002437],[273,350,1660098002454],[276,350,1660098002477],[276,349,1660098002628],[271,347,1660098002644],[260,340,1660098002660],[242,333,1660098002676],[221,326,1660098002692],[196,318,1660098002708],[168,310,1660098002724],[138,300,1660098002740],[112,291,1660098002756],[90,283,1660098002772],[75,278,1660098002788],[66,274,1660098002805],[61,272,1660098002823],[59,271,1660098002848],[59,270,1660098003560],[59,269,1660098003578],[57,266,1660098003596],[57,262,1660098003613],[57,258,1660098003630],[57,254,1660098003647],[57,250,1660098003668],[57,247,1660098003687],[56,245,1660098003703],[56,242,1660098003726],[55,241,1660098003745],[56,241,1660098004044],[65,237,1660098004061],[78,232,1660098004077],[97,223,1660098004093],[122,211,1660098004109],[144,198,1660098004125],[165,186,1660098004141],[181,174,1660098004157],[191,166,1660098004173],[195,160,1660098004195],[196,159,1660098004284],[197,167,1660098004300],[199,185,1660098004316],[202,215,1660098004332],[205,253,1660098004348],[209,291,1660098004364],[213,330,1660098004380],[219,364,1660098004396],[224,391,1660098004412],[229,408,1660098004428],[233,419,1660098004445],[236,424,1660098004466],[238,426,1660098004486],[238,427,1660098004526],[239,427,1660098004550],[240,427,1660098004589],[241,427,1660098004629],[242,429,1660098004646],[244,430,1660098004663],[246,432,1660098004679],[250,435,1660098004696],[254,438,1660098004714],[258,442,1660098004730],[261,446,1660098004747],[263,449,1660098004763],[266,451,1660098004780],[269,454,1660098004798],[273,456,1660098004814],[275,459,1660098004833],[277,460,1660098004849],[278,461,1660098004865]],"mm-mp":3.4992167101827647,"md":[[159,273,1660097999284],[181,173,1660098000039],[192,365,1660098002200],[276,350,1660098002503],[59,271,1660098003258],[55,241,1660098003899],[279,461,1660098004874]],"md-mp":931.6666666666666,"mu":[[159,273,1660097999349],[181,173,1660098000101],[192,365,1660098002284],[276,350,1660098002561],[59,271,1660098003388],[55,241,1660098004011],[279,461,1660098004957]],"mu-mp":934.6666666666666,"topLevel":{"st":1660097996558,"sc":{"availWidth":1920,"availHeight":1040,"width":1920,"height":1080,"colorDepth":24,"pixelDepth":24,"availLeft":0,"availTop":0,"onchange":null,"isExtended":false},"nv":{"vendorSub":"","productSub":"20030107","vendor":"Google Inc.","maxTouchPoints":0,"scheduling":{},"userActivation":{},"doNotTrack":null,"geolocation":{},"connection":{},"pdfViewerEnabled":true,"webkitTemporaryStorage":{},"webkitPersistentStorage":{},"hardwareConcurrency":16,"cookieEnabled":true,"appCodeName":"Mozilla","appName":"Netscape","appVersion":"5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36","platform":"Win32","product":"Gecko","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36","language":"fr-FR","languages":["fr-FR"],"onLine":true,"webdriver":false,"bluetooth":{},"clipboard":{},"credentials":{},"keyboard":{},"managed":{},"mediaDevices":{},"storage":{},"serviceWorker":{},"wakeLock":{},"deviceMemory":8,"ink":{},"hid":{},"locks":{},"mediaCapabilities":{},"mediaSession":{},"permissions":{},"presentation":{},"serial":{},"virtualKeyboard":{},"usb":{},"xr":{},"userAgentData":{"brands":[{"brand":"Chromium","version":"104"},{"brand":" Not A;Brand","version":"99"},{"brand":"Google Chrome","version":"104"}],"mobile":false,"platform":"Windows"},"plugins":["internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer","internal-pdf-viewer"]},"dr":"","inv":false,"exec":false,"wn":[[460,969,1,1660097996559]],"wn-mp":0,"xy":[[0,0,1,1660097996559]],"xy-mp":0,"mm":[[372,369,1660097996885],[97,388,1660097998169],[391,408,1660097998283],[425,412,1660097998299],[441,414,1660097998315],[444,416,1660097998506],[442,418,1660097998522],[437,422,1660097998539],[427,427,1660097998555],[415,432,1660097998571],[401,439,1660097998587]],"mm-mp":18.06315789473684},"v":1}`,
	}

	MotionTime = map[string]string{
		"GetCaptcha":   "1660097",
		"CheckCaptcha": "1660097",
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

/*
type Captcha struct {
	Success bool `json:"success"`
	ErrorCode []string `json:"error-codes"`
	C struct {
		Type string `json:"type"`
		Req  string `json:"req"`
	} `json:"c"`
	ChallengeURI  string `json:"challenge_uri"`
	Key           string `json:"key"`
	RequestConfig struct {
		Version                      int         `json:"version"`
		ShapeType                    interface{} `json:"shape_type"`
		MinPoints                    interface{} `json:"min_points"`
		MaxPoints                    interface{} `json:"max_points"`
		MinShapesPerImage            interface{} `json:"min_shapes_per_image"`
		MaxShapesPerImage            interface{} `json:"max_shapes_per_image"`
		RestrictToCoords             interface{} `json:"restrict_to_coords"`
		MinimumSelectionAreaPerShape interface{} `json:"minimum_selection_area_per_shape"`
		MultipleChoiceMaxChoices     int         `json:"multiple_choice_max_choices"`
		MultipleChoiceMinChoices     int         `json:"multiple_choice_min_choices"`
		OverlapThreshold             interface{} `json:"overlap_threshold"`
	} `json:"request_config"`
	RequestType       string `json:"request_type"`
	RequesterQuestion struct {
		Fr string `json:"fr"`
		En string `json:"en"`
	} `json:"requester_question"`
	RequesterQuestionExample []string `json:"requester_question_example"`
	Tasklist                 []struct {
		DatapointURI string `json:"datapoint_uri"`
		TaskKey      string `json:"task_key"`
	} `json:"tasklist"`
	BypassMessage string `json:"bypass-message"`
	GeneratedPassUUID string `json:"generated_pass_UUID"`
}
*/
