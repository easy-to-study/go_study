package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	name := "Go"
	text := "Hello from go"
	channel := "times_t-sataga"

	jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + text + `"}`

	req, err := http.NewRequest(
		"POST",
		"https://hooks.slack.com/services/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",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(resp)
	defer resp.Body.Close()
}
