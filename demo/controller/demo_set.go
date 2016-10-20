package controller

import (
	"encoding/json"
	"net/http"

	"github.com/simplejia/wsp/demo/service"
)

// @prefilter("Login", {"Method":{"type":"get"}})
// @postfilter("Boss")
func (demo *Demo) Set(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	value := r.FormValue("value")
	demoService := service.NewDemo()
	demoService.Set(key, value)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 0,
	})
}
