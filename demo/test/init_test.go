package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	_ "github.com/simplejia/wsp/demo"
)

func h(path string, m map[string]string) string {
	var (
		w *httptest.ResponseRecorder
		r *http.Request
	)

	params := ""
	for k, v := range m {
		params += fmt.Sprintf("%s=%s&", k, v)
	}

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1/%s?%s", path, params), nil)
	http.DefaultServeMux.ServeHTTP(w, r)

	return w.Body.String()
}
