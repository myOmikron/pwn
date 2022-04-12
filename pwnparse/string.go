package pwnparse

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func Pointer(in string) []string {
	ret := make([]string, 0)
	s := strings.Split(in, "0x")
	for _, elem := range s {
		decode, _ := hex.DecodeString(elem)
		ret = append(ret, fmt.Sprint(string(decode)))
	}
	return ret
}

func DecodeLittleEndian(in []string) []string {
	ret := make([]string, len(in))

	for _, elem := range in {
		var result string
		for _, v := range elem {
			result = string(v) + result
		}
		ret = append(ret, result)
	}

	return ret
}

func From(in []string) []string {
	r := make([]string, len(in))
	for _, elem := range in {
		var s string
		for i := 0; i < len(elem); i++ {
			if i%2 == 0 {
			}
		}
		r = append(r, s)
	}
	return r
}
