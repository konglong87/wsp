package mysql

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/simplejia/wsp/demo/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/simplejia/utils"
)

var DBS map[string]*sql.DB = map[string]*sql.DB{}

func parseDBFile(path string) {
	type c struct {
		ConnMaxLifetime string
		MaxIdleConns    int
		MaxOpenConns    int
		Dsn             string
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

	db, err := sql.Open("mysql", cc.Dsn)
	if err != nil {
		panic(err)
	}
	dur, _ := time.ParseDuration(cc.ConnMaxLifetime)
	db.SetConnMaxLifetime(dur)
	db.SetMaxIdleConns(cc.MaxIdleConns)
	db.SetMaxOpenConns(cc.MaxOpenConns)

	key := strings.TrimSuffix(filepath.Base(path), "_db.json")
	DBS[key] = db
}

func init() {
	dir, _ := os.Getwd()
	err := filepath.Walk(
		filepath.Join(dir, "mysql"),
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

			parseDBFile(path)
			return
		},
	)
	if err != nil {
		panic(err)
	}
}
