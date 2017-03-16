package controller_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/simplejia/lc"
	_ "github.com/simplejia/wsp/demo/clog"
	_ "github.com/simplejia/wsp/demo/conf"
	_ "github.com/simplejia/wsp/demo/mysql"
	_ "github.com/simplejia/wsp/demo/redis"
)

func init() {
	lc.Init(1e5)
}

type ft func(http.ResponseWriter, *http.Request)

func h(f ft, params interface{}) (body []byte, err error) {
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

	r, err := http.NewRequest("POST", "", reader)
	if err != nil {
		return
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	f(w, r)
	body = w.Body.Bytes()
	if g, e := w.Code, http.StatusOK; g != e {
		err = fmt.Errorf("http resp status not ok: %s", http.StatusText(g))
		return
	}
	return
}
