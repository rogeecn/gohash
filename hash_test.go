package gohash

import (
	"testing"
)

func Test_Hash(t *testing.T) {
	encodeStr := "admin"
	t.Logf("encodestr: %s", encodeStr)

	t.Logf("use md5: %s", Hash(MD5, encodeStr))
	t.Logf("use base64: %s", Hash(BASE64, encodeStr))

	t.Log("DONE!")
}
