# using solving node

```go
func (c *Captcha) GetAnwser() map[string]string {
	item := strings.ReplaceAll(strings.Split(strings.ReplaceAll(c.RequesterQuestion.En, "an", "a"), "Please click each image containing a ")[1], ".", "")
	response := map[string]string{}

	/*client := http.Client{
		Timeout: 500 * time.Second,
	}*/

	for _, task := range c.Tasklist {
		/*resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:1337/check/%s/%s", base64.StdEncoding.EncodeToString([]byte(item)), base64.StdEncoding.EncodeToString([]byte(task.DatapointURI))))

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
		*/

		response[task.TaskKey] = strconv.FormatBool(strings.Contains(task.DatapointURI, item))
	}

	return response
}
```

# get hsw with browser
```go
/*client := http.Client{
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
	}*/
```

# xx
```go

	func (c *Captcha) GetAnswer() map[string]string {
		item := strings.ReplaceAll(strings.Split(strings.ReplaceAll(c.RequesterQuestion.En, "an", "a"), "Please click each image containing a ")[1], ".", "")
		response := map[string]string{}

		for _, task := range c.Tasklist {
			response[task.TaskKey] = strconv.FormatBool(strings.Contains(task.DatapointURI, item))
		}

		return response
	}

```