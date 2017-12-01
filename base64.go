package gohash

import (
	"encoding/base64"
)

func hashByBase64(encodeStr string) string {
	return base64.StdEncoding.EncodeToString([]byte(encodeStr))
}
