package model

import (
	"time"

	"github.com/simplejia/clog"
	"github.com/simplejia/lm"
)

func (demo *Demo) Get(key string) (value string) {
	demoModel := NewDemo()
	err := lm.GlueLc(
		key,
		demoModel,
		func(p, r interface{}) error {
			return nil
		},
		func(p interface{}) string {
			return p.(string)
		},
		&lm.LcStru{
			Expire: time.Second,
			Safety: true,
		},
	)
	if err != nil {
		clog.Error("Demo:Get() %v", err)
		return
	}

	value = demoModel.Value
	return
}
