package service

import "github.com/simplejia/wsp/demo/model"

func (demo *Demo) Get(key string) (value string) {
	demoModel := model.NewDemo()
	value = demoModel.Get(key)
	return
}
