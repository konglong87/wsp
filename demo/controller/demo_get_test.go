package controller_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/simplejia/wsp/demo/controller"
)

func DemoGet(key string) (value string, err error) {
	params := map[string]string{
		"key": key,
	}
	c := &controller.Demo{}
	respStr, err := h(c.Get, params)
	if err != nil {
		return
	}

	var respStru *struct {
		Code int
		Data string
	}
	json.Unmarshal(respStr, &respStru)
	if respStru == nil || respStru.Code != 0 {
		err = fmt.Errorf("respStru: %v, respStr: %s", respStru, respStr)
		return
	}

	value = respStru.Data
	return
}

func TestDemoGet(t *testing.T) {
	key, value := "key", "value"
	DemoSet(key, value)
	_value, err := DemoGet(key)
	if err != nil {
		t.Fatal(err)
	}
	if value != _value {
		t.Fatal("")
	}
}
