package filter

import (
	"net/http"
	"time"

	"github.com/simplejia/clog"
)

func Boss(w http.ResponseWriter, r *http.Request, m map[string]interface{}) bool {
	defer func() {
		recover()
	}()

	err := m["__E__"]
	elapse := m["__T__"].(time.Time)

	// ...

	clog.Info("Boss() %v, %v", err, elapse)
	return true
}
