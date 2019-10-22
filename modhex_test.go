package yubilib

import (
	hex2 "encoding/hex"
	"fmt"
	"testing"
	 . "github.com/smartystreets/goconvey/convey"
)


func TestModHex2Hex (t *testing.T) {
	for idx, key := range goodKeys {
		Convey(fmt.Sprintf("Modhex test %d",idx), t ,func () {
			hex := ModHex2Hex(key.Modhex)
			So(hex,ShouldEqual,key.Ciphertext)

		})
	}
}


func TestByte2ModHex (t *testing.T) {
	for idx, key := range goodKeys {
		Convey(fmt.Sprintf("Modhex test %d",idx), t ,func () {
			by,_ := hex2.DecodeString(key.Ciphertext)
			hex := Byte2ModHex(by)
			So(hex,ShouldEqual,key.Modhex)

		})
	}
}

func BenchmarkModhex2Hex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_= ModHex2Hex("hrjtdcrcrijvhfggcvdninchhujvhdcn")
	}
}
func BenchmarkByte2Modhex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_= Byte2ModHex([]byte("hrj25tdcr4crijvh"))
	}
}