package main

import (
	"log"

	"github.com/simplejia/clog"
	"github.com/simplejia/lc"

	"net/http"

	_ "github.com/simplejia/wsp/demo/clog"
	_ "github.com/simplejia/wsp/demo/conf"
	_ "github.com/simplejia/wsp/demo/mysql"
	_ "github.com/simplejia/wsp/demo/redis"
)

func init() {
	lc.Init(1e5)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
}

func main() {
	clog.Info("main()")

	log.Panic(http.ListenAndServe(":8080", nil))
}
