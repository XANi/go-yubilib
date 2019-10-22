package yubilib

import (
	"encoding/hex"
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)


func TestDecryptAES128ECB (t *testing.T) {
	for idx, key := range goodKeys {
		Convey(fmt.Sprintf("decrypt correct key %d",idx), t ,func () {
			out, err :=  DecryptAES128ECB(key.Key,key.Ciphertext)
			So(err,ShouldBeNil)
			So(hex.EncodeToString(out),ShouldEqual,key.Plain)

		})
		Convey(fmt.Sprintf("decrypt incorrect key lenght %d",idx), t ,func () {
			out, err :=  DecryptAES128ECB(key.Key + "af",key.Ciphertext)
			So(err,ShouldNotBeNil)
			So(hex.EncodeToString(out),ShouldEqual,"")

		})
		Convey(fmt.Sprintf("decrypt incorrect hex key %d",idx), t ,func () {
			out, err :=  DecryptAES128ECB(key.Key + "zz11111111111111",key.Ciphertext)
			So(err,ShouldNotBeNil)
			So(hex.EncodeToString(out),ShouldEqual,"")

		})
		Convey(fmt.Sprintf("decrypt incorrect ciphertext %d",idx), t ,func () {
			out, err :=  DecryptAES128ECB(key.Key,key.Ciphertext+"zz")
			So(err,ShouldNotBeNil)
			So(hex.EncodeToString(out),ShouldEqual,"")

		})
	}

}
func TestEncryptAES128ECB (t *testing.T) {
	for idx, key := range goodKeys {
		Convey(fmt.Sprintf("encrypt correct %d", idx), t, func() {
			data, _ := hex.DecodeString(key.Plain)
			out, err := EncryptAES128ECB(key.Key, data)
			So(err, ShouldBeNil)
			So(out, ShouldEqual, key.Ciphertext)

		})
		Convey(fmt.Sprintf("encrypt bad data length %d", idx), t, func() {
			data, _ := hex.DecodeString(key.Plain+"11")
			out, err := EncryptAES128ECB(key.Key, data)
			So(err, ShouldNotBeNil)
			So(out, ShouldEqual, "")
		})
		Convey(fmt.Sprintf("encrypt bad aes length %d", idx), t, func() {
			data, _ := hex.DecodeString(key.Plain)
			out, err := EncryptAES128ECB(key.Key+"11", data)
			So(err, ShouldNotBeNil)
			So(out, ShouldEqual, "")
		})
		Convey(fmt.Sprintf("encrypt bad aes hex %d", idx), t, func() {
			data, _ := hex.DecodeString(key.Plain)
			out, err := EncryptAES128ECB("112233445566778899aabbccddeeffgg", data)
			So(err, ShouldNotBeNil)
			So(out, ShouldEqual, "")
		})
	}

}
func BenchmarkDecryptAES128ECB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_,_ = DecryptAES128ECB("00112233445566778899aabbccddeeff","6c8d20c0c78f64550f2b7b066e8f620b")
	}
}
func BenchmarkEncryptAES128ECB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_,_ = EncryptAES128ECB("00112233445566778899aabbccddeeff", []byte("eerow6looz0Cei3heeNgiSeicu3elae8"))
	}
}