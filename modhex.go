package yubilib

import (
	"regexp"
	"strings"
)
var RegexpModhex = regexp.MustCompile("^[cbdefghijklnrtuv]+$")
func modhexMap(r rune) (o rune) {
	switch r {
	case 'c':
		return '0'
	case 'b':
		return '1'
	case 'd':
		return '2'
	case 'e':
		return '3'
	case 'f':
		return '4'
	case 'g':
		return '5'
	case 'h':
		return '6'
	case 'i':
		return '7'
	case 'j':
		return '8'
	case 'k':
		return '9'
	case 'l':
		return 'a'
	case 'n':
		return 'b'
	case 'r':
		return 'c'
	case 't':
		return 'd'
	case 'u':
		return 'e'
	case 'v':
		return 'f'
	default: return 'X'
	}
}

var modhexHexMap [256]byte
var byteModhexMap [256]string
var nibbleHexMap =  []byte("cbdefghijklnrtuv")

func init() {
	for i, _ := range modhexHexMap {
		modhexHexMap[i] = byte(modhexMap(rune(i)))
		byteModhexMap[i] = string(nibbleHexMap[i>>4&0x0f]) + string(nibbleHexMap[i&0x0f])
	}
}

// ModHex2Hex returns hex encoded string from modhex. unconvertable characters are replaced with X
func ModHex2Hex (modhex string) (hex string) {
	out := make([]byte,len(modhex))
	for i, v := range([]byte(modhex)) {
		out[i] = modhexHexMap[int(v)]
	}
	return string(out)
}

func Byte2ModHex(b []byte) (string) {
	var strb strings.Builder
	for _, by := range b {
		strb.WriteString(byteModhexMap[by])
	}
	return strb.String()
}