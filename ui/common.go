package ui

import (
	URL "net/url"
)

type JsonData = map[string]any

func JSONHandler(status int, message string, data any) JsonData {
	return JsonData{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

func addAuthtoUrl(url, username, password string) string {
	if username == "" && password == "" {
		return url
	}

	urlParse, err := URL.Parse(url)
	if err != nil {
		return ""
	}

	urlParse.User = URL.UserPassword(username, password)
	return urlParse.String()
}
