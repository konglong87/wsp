package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func DemoGet(key string) (value string, err error) {
	respStr := h("Demo/Get",
		map[string]string{
			"key": key,
		})

	respStru := &struct {
		Code int
		Data string
	}{}
	json.Unmarshal([]byte(respStr), &respStru)
	if respStru.Code != 0 {
		err = fmt.Errorf("respStru: %v, respStr: %v", respStru, respStr)
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
