package controller_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/simplejia/wsp/demo/controller"
)

func DemoSet(key, value string) (err error) {
	params := map[string]string{
		"key":   key,
		"value": value,
	}
	c := &controller.Demo{}
	respStr, err := h(c.Set, params)
	if err != nil {
		return
	}

	var respStru *struct {
		Code int
	}
	json.Unmarshal(respStr, &respStru)
	if respStru == nil || respStru.Code != 0 {
		err = fmt.Errorf("respStru: %v, respStr: %s", respStru, respStr)
		return
	}

	// sleep a while, because lc cache may have a delayed effect
	time.Sleep(time.Millisecond)

	return
}

func TestDemoSet(t *testing.T) {
	key, value := "key", "value"
	err := DemoSet(key, value)
	if err != nil {
		t.Fatal(err)
	}
	_value, _ := DemoGet(key)
	if value != _value {
		t.Fatal("")
	}
}
