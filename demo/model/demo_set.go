package model

import (
	"time"

	"github.com/simplejia/lc"
)

func (demo *Demo) Set(key, value string) (err error) {
	demoModel := NewDemo()
	demoModel.Key = key
	demoModel.Value = value
	lc.Set(key, demoModel, time.Minute)
	return
}
