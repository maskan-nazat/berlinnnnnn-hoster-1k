package discord

func (d *DiscordSession) HeaderRegister() map[string]string {
	return map[string]string{
		`authority`:       `discord.com`,
		`accept`:          `*/*`,
		`accept-language`: `fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`,
		`content-type`:    `application/json`,
		// `cookie`: `__dcfduid=1deca2b01adf11ed9fee919fa9fad8fd; __sdcfduid=1deca2b11adf11ed9fee919fa9fad8fd1de8e69d387a6b1bee3abf47b9b6e20be8dafdb981d63a729938122da5deb417; locale=fr; __cf_bm=P_ZkFvINmqipEcMDRELnUa46KS3beBbD5h9GAHwwkVE-1660378110-0-AUVfXD1Qx8p8uOOy/Q8TieyTvP1ai+X0YWZJZci23ZTWJlQYFOfafg8Gd0ILn6yKRbwJ6JsGUSmQxBnBww3v57jcYhhMrR+v6Md+d+F2AXlOoc+0EedV+03TmEm092S2jw==`,
		`origin`:             `https://discord.com`,
		`referer`:            `https://discord.com/`,
		`sec-ch-ua`:          `"Chromium";v="104", " Not A;Brand";v="99", "Google Chrome";v="104"`,
		`sec-ch-ua-mobile`:   `?0`,
		`sec-ch-ua-platform`: `"Windows"`,
		`sec-fetch-dest`:     `empty`,
		`sec-fetch-mode`:     `cors`,
		`sec-fetch-site`:     `same-origin`,
		`user-agent`:         `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36`,
		//`x-fingerprint`: `1007923590630490142.Uc_tXtxLQZMqSpm-OkV7j4ccP2k`,
		// `x-track`: `eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6ImZyLUZSIiwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzEwNC4wLjAuMCBTYWZhcmkvNTM3LjM2IiwiYnJvd3Nlcl92ZXJzaW9uIjoiMTA0LjAuMC4wIiwib3NfdmVyc2lvbiI6IjEwIiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjk5OTksImNsaWVudF9ldmVudF9zb3VyY2UiOm51bGx9`,
	}
}
