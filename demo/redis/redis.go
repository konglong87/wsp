package conf4redis

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/simplejia/wsp/demo/conf"

	"github.com/garyburd/redigo/redis"
	"github.com/simplejia/utils"
)

var RDS map[string]*redis.Pool = map[string]*redis.Pool{}

func parseRDFile(path string) {
	type c struct {
		ConnMaxLifetime string
		MaxIdleConns    int
		MaxOpenConns    int
		Addr            string
	}
	var cs struct {
		Env  string
		Envs map[string]*c
	}
	fcontent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fcontent = utils.RemoveAnnotation(fcontent)
	if err := json.Unmarshal(fcontent, &cs); err != nil {
		panic(err)
	}
	env := conf.Env
	if cs.Env != "" {
		env = cs.Env
	}
	cc := cs.Envs[env]
	if cc == nil {
		panic("env not right")
	}

	rd := &redis.Pool{
		MaxIdle:     cc.MaxIdleConns,
		IdleTimeout: time.Second * 240,
		Dial: func() (c redis.Conn, err error) {
			server := cc.Addr
			c, err = redis.Dial("tcp", server,
				redis.DialReadTimeout(time.Second*5),
				redis.DialConnectTimeout(time.Second),
			)
			if err != nil {
				return
			}
			return
		},
	}

	key := strings.TrimSuffix(filepath.Base(path), ".json")
	RDS[key] = rd
}

func init() {
	dir, _ := os.Getwd()
	err := filepath.Walk(
		filepath.Join(dir, "redis"),
		func(path string, info os.FileInfo, err error) (reterr error) {
			if err != nil {
				reterr = err
				return
			}
			if info.IsDir() {
				return
			}
			if filepath.Ext(path) != ".json" {
				return
			}
			parseRDFile(path)
			return
		},
	)
	if err != nil {
		panic(err)
	}
}
