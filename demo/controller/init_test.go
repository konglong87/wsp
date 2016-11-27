package controller_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"

	_ "github.com/simplejia/wsp/demo"
)

type gpp struct {
	Path    string
	Timeout time.Duration
	Headers map[string]string
	Params  interface{}
}

func get(p *gpp) (body []byte, err error) {
	path, timeout, headers, params := p.Path, p.Timeout, p.Headers, p.Params
	uri := fmt.Sprintf("http://127.0.0.1/%s", strings.TrimPrefix(path, "/"))
	switch v := params.(type) {
	case map[string]string:
		u, _ := url.Parse(uri)
		values := u.Query()
		for key, value := range v {
			values.Set(key, value)
		}
		u.RawQuery = values.Encode()
		uri = u.String()
	}
	return sendHttpRequest("GET", uri, timeout, headers, nil)
}

func post(p *gpp) (body []byte, err error) {
	path, timeout, headers, params := p.Path, p.Timeout, p.Headers, p.Params
	uri := fmt.Sprintf("http://127.0.0.1/%s", strings.TrimPrefix(path, "/"))
	var reader io.Reader
	switch v := params.(type) {
	case map[string]string:
		values := url.Values{}
		for key, value := range v {
			values.Set(key, value)
		}
		reader = strings.NewReader(values.Encode())
	case string:
		reader = strings.NewReader(v)
	case []byte:
		reader = bytes.NewReader(v)
	}
	return sendHttpRequest("POST", uri, timeout, headers, reader)
}

func sendHttpRequest(
	method string,
	uri string,
	timeout time.Duration,
	headers map[string]string,
	bodyReader io.Reader,
) (body []byte, err error) {
	r, err := http.NewRequest(method, uri, bodyReader)
	if err != nil {
		return
	}
	if host, ok := headers["Host"]; ok {
		r.Host = host
	}
	for name, value := range headers {
		r.Header.Set(name, value)
	}
	if bodyReader != nil {
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	w := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(w, r)

	if g, e := w.Code, http.StatusOK; g != e {
		err = fmt.Errorf("http resp code: %d", g)
		return
	}

	body = w.Body.Bytes()
	return
}
