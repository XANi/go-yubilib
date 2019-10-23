package yubilib

import (
	"encoding/hex"
	"fmt"
	"testing"
	 . "github.com/smartystreets/goconvey/convey"
)

func TestDecodeOTP (t *testing.T) {
	for idx, key := range goodKeys {
		Convey(fmt.Sprintf("Decode OTP %d",idx), t ,func () {
			plainBytes, _ := hex.DecodeString(key.Plain)
			out, err := DecodeOTP(plainBytes)
			So(err,ShouldBeNil)
			So(out.SessionCounter,ShouldEqual,key.SessionCounter)
			So(out.YKTSLow,ShouldEqual,key.Low)
			So(out.YKTSHigh,ShouldEqual,key.High)
			So(out.UseCounter,ShouldEqual,key.Use)

		})
	}
}

func TestOTPDecoder(t *testing.T) {
	for idx, key := range goodKeys {
		Convey(fmt.Sprintf("OTP decoder good key %d",idx), t ,func () {
			d := NewDecoder(
				func(v string)(a string, b string) {
					if v == key.Publicname {
						return key.Key,key.Internalname
					}
					return
				})

			otp, err := d.Decode(key.Publicname + key.Modhex)
			So(err,ShouldBeNil)
			So(otp.SessionCounter,ShouldEqual,key.SessionCounter)
			So(otp.YKTSLow,ShouldEqual,key.Low)
			So(otp.YKTSHigh,ShouldEqual,key.High)
			So(otp.UseCounter,ShouldEqual,key.Use)
			So(otp.PublicName,ShouldEqual,key.Publicname)
		})
		Convey(fmt.Sprintf("OTP decoder absent key %d",idx), t ,func () {
			d := NewDecoder(
				func(v string)(a string, b string) {
					return
				})

			otp, err := d.Decode(key.Publicname + key.Modhex)
			So(err,ShouldNotBeNil)
			So(err,ShouldHaveSameTypeAs,&ErrorNotFound{})
			So(otp,ShouldBeNil)
		})
		Convey(fmt.Sprintf("OTP decoder bad modhex %d",idx), t ,func () {
			d := NewDecoder(
				func(v string)(a string,b string) {
					if v == key.Publicname {
						return key.Key, key.Internalname
					}
					return
				})
			otp, err := d.Decode(key.Publicname + key.Modhex + "z")
			So(err,ShouldNotBeNil)
			So(otp,ShouldBeNil)
		})
		Convey(fmt.Sprintf("OTP decoder bad aeskey %d",idx), t ,func () {
			d := NewDecoder(
				func(v string)(a string,b string) {
					if v == key.Publicname {
						return key.Key + "aa", key.Internalname
					}
					return
				})
			otp, err := d.Decode(key.Publicname + key.Modhex)
			So(err,ShouldNotBeNil)
			So(otp,ShouldBeNil)
		})

		Convey(fmt.Sprintf("OTP decoder bad publicname %d",idx), t ,func () {
			d := NewDecoder(
				func(v string)(a string,b string) {
					if v == key.Publicname {
						return key.Key, key.Internalname[1:len(key.Internalname)] + "a"
					}
					return
				})
			otp, err := d.Decode(key.Publicname + key.Modhex)
			So(err,ShouldNotBeNil)
			So(otp,ShouldBeNil)
		})
		Convey(fmt.Sprintf("OTP decoder bad checksum %d",idx), t ,func () {
			d := NewDecoder(
				func(v string)(a string,b string) {
					if v == key.Publicname {
						return key.Key, key.Internalname
					}
					return
				})
			otp, err := d.Decode(key.Publicname + key.Modhex[:len(key.Modhex)-1]+"v")
			So(err,ShouldNotBeNil)
			So(err.Error(),ShouldContainSubstring,"crc")
			So(otp,ShouldBeNil)
		})
	}
}