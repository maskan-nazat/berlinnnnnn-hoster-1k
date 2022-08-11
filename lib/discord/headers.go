package discord

func (d *DiscordSession) HeaderRegister() map[string]string {
	return map[string]string{
		`authority`: `canary.discord.com`,
		//`x-super-properties`: `eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJjYW5hcnkiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC40OCIsIm9zX3ZlcnNpb24iOiIxMC4wLjIyMDAwIiwib3NfYXJjaCI6Ing2NCIsInN5c3RlbV9sb2NhbGUiOiJmciIsImNsaWVudF9idWlsZF9udW1iZXIiOjEzODkyNCwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`,
		//`x-fingerprint`: `1002379148485988426.90hnnPzrrAPe5kqQE9NXvpO2daI`,
		//`x-discord-locale`: `fr`,
		//`x-debug-options`: `bugReporterEnabled`,
		`accept-language`: `fr,fr-FR;q=0.9`,
		`user-agent`:      `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.48 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36`,
		// Already added when you pass json=
		`content-type`:   `application/json`,
		`accept`:         `*/*`,
		`origin`:         `https://canary.discord.com`,
		`sec-fetch-site`: `same-origin`,
		`sec-fetch-mode`: `cors`,
		`sec-fetch-dest`: `empty`,
		`referer`:        `https://canary.discord.com/register`,
		// Requests sorts cookies= alphabetically
		// `cookie`: `__dcfduid=54ec07400ed911ed89c74d9dda7364b9; __sdcfduid=54ec07410ed911ed89c74d9dda7364b95ced8233c2fce1baab00bb89c08a6c2445f343f0bef174e0e1b74a698bf8c992; __cf_bm=8C_LtYHpleV6wrRxIFPxqonlu3dhBNPNGyKYMkkzeCE-1659056212-0-AW7LPmF2YRHdTqOgLWbcUAIwD0V4pzNw90sTN9VkBT9dSfx/aBIPecn+kx+gk+RMM7EJ7j0C6f9gD2jbN3l7RUrW1nW7bmg8BhJOfN2SZSKdV+u8r2zWsLeuUrQLuZJtvw==`,
	}
}
