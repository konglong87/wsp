package controller

import (
	"encoding/json"
	"net/http"

	"github.com/simplejia/wsp/demo/service"
)

// @postfilter("Boss")
func (demo *Demo) Get(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	demoService := service.NewDemo()
	value := demoService.Get(key)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 0,
		"data": value,
	})
}
