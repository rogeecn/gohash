package gohash

import (
	"crypto/md5"
	"io"
	"fmt"
)

func hashByMd5(encodeStr string) string {
	instance := md5.New()
	io.WriteString(instance, encodeStr)
	return fmt.Sprintf("%x", instance.Sum(nil))
}
