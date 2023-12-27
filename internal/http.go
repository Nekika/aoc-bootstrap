package internal

import (
	"io"
	"net/http"
)

func HttpRequestWithSessionCookie(method, url, cookie string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})

	return req, nil
}

func HttpGetWithSessionCookie(url, cookie string) (*http.Request, error) {
	return HttpRequestWithSessionCookie(http.MethodGet, url, cookie, nil)
}
