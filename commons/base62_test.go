package commons

import "testing"

func TestEncode(t *testing.T) {
	ret := Encode(56900235584)
	t.Log(ret)
	t.Log(Decode(ret))
}

func TestDecode(t *testing.T) {
	t.Log()
}
