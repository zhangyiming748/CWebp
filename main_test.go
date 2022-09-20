package CWebp

import "testing"

func TestCWebp(t *testing.T) {
	src := "/Users/zen/Github/CWebp/mac"
	dst := "/Users/zen/Github/CWebp/mac"
	pattern := "png"
	CWebp(src, pattern, dst)
}
func TestGetFile(t *testing.T) {
	ret := getFiles("/Users/zen/Github/CWebp/mac", "png")
	t.Log(ret)
}
