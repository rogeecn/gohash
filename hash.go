package gohash

const (
	MD5 = iota
	BASE64
)

func Hash(t int, encodeStr string) string {
	switch t {
	case MD5:
		return hashByMd5(encodeStr)
	case BASE64:
		return hashByBase64(encodeStr)
	default:
		return hashByMd5(encodeStr)
	}
	return ""
}
