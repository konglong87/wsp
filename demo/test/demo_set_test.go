package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func DemoSet(key, value string) (err error) {
	respStr := h("Demo/Set",
		map[string]string{
			"key":   key,
			"value": value,
		})

	respStru := &struct {
		Code int
	}{}
	json.Unmarshal([]byte(respStr), &respStru)
	if respStru.Code != 0 {
		err = fmt.Errorf("respStru: %v, respStr: %v", respStru, respStr)
		return
	}

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
