package wxcrypt

import "testing"

func TestPKCS7Encoder(t *testing.T) {
	s := pkcg7Encode([]byte("123"))
	t.Log(s)
	t.Log(string(s))

	// decode
	d := pkcg7Decode(s)
	t.Log(d)
	t.Log(string(d))
}

func TestChr(t *testing.T) {
	b := 23
	t.Log(string(b))
	b1 := b & 0xFF
	t.Log(string(b1))
	if string(b) == string(b1) {
		t.Log("is the same")

	}

}
