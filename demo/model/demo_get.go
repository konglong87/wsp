package model

import (
	"time"

	"github.com/simplejia/clog"
	"github.com/simplejia/lm"
)

func (demo *Demo) Get(key string) (value string) {
	demoModel := NewDemo()
	lmStru := &lm.LmStru{
		Input:  key,
		Output: demoModel,
		Proc: func(p, r interface{}) error {
			return nil
		},
		Key: func(p interface{}) string {
			return p.(string)
		},
		Lc: &lm.LcStru{
			Expire: time.Second,
			Safety: true,
		},
	}
	err := lm.GlueLc(lmStru)
	if err != nil {
		clog.Error("Demo:Get() %v", err)
		return
	}

	value = demoModel.Value
	return
}
