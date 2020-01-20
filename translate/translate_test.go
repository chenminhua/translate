package translate

import "testing"

func TestTranslateEnToZh(t *testing.T)  {
	if "你好" != Translate("en", "zh", "hello") {
		t.Fatal("hello wrong")
	}

	if "你好，世界" != Translate("en", "zh", "hello world") {
		t.Fatal("hello wrong")
	}
}