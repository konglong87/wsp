package conf4redis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/simplejia/wsp/demo/conf"

	"github.com/garyburd/redigo/redis"
	"github.com/simplejia/utils"
)

type Conf struct {
	ConnMaxLifetime string
	MaxIdleConns    int
	MaxOpenConns    int
	Addr            string
}

var (
	RDS  map[string]*redis.Pool = map[string]*redis.Pool{}
	Envs map[string]*Conf
	Env  string
	C    *Conf
)

func parseRDFile(path string) {
	fcontent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fcontent = utils.RemoveAnnotation(fcontent)
	if err := json.Unmarshal(fcontent, &Envs); err != nil {
		panic(err)
	}

	Env = conf.Env
	C = Envs[Env]
	if C == nil {
		fmt.Println("env not right:", Env)
		os.Exit(-1)
	}

	rd := &redis.Pool{
		MaxIdle:     C.MaxIdleConns,
		IdleTimeout: time.Second * 240,
		Dial: func() (c redis.Conn, err error) {
			server := C.Addr
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
