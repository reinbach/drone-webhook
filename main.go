package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	body := strings.NewReader(
		os.Getenv("PLUGIN_BODY"),
	)
	client := &http.Client{}

	req, err := http.NewRequest(
		os.Getenv("PLUGIN_METHOD"),
		os.Getenv("PLUGIN_URL"),
		body,
	)
	if err != nil {
		os.Exit(1)
	}

	content_type := os.Getenv("PLUGIN_CONTENT_TYPE")
	if content_type != "" {
		req.Header.Add("Content-Type", content_type)
	}
	resp, err := client.Do(req)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("%s", resp.Status)
}
