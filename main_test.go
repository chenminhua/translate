package main

import "testing"

func TestTranslateEnToZh(t *testing.T)  {
	if "你好" != translate("en", "zh", "hello") {
		t.Fatal("hello wrong")
	}

	if "你好，世界" != translate("en", "zh", "hello world") {
		t.Fatal("hello wrong")
	}
}