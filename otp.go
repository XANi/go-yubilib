package yubilib

import (
	"encoding/binary"
	"fmt"
	"strings"
)

func DecodeOTP(decryptedOTPData []byte) (*YubikeyOTP,error) {
	var y YubikeyOTP
	y.SessionCounter =  binary.LittleEndian.Uint16(decryptedOTPData[6:8])
	y.YKTSLow =  binary.LittleEndian.Uint16(decryptedOTPData[8:10])
	y.YKTSHigh =  uint8(decryptedOTPData[10])
	y.UseCounter =  uint8(decryptedOTPData[11])
	fmt.Printf("%x", decryptedOTPData)
	return &y,nil

}

type OTPDecoder struct {
	keystore func(publicName string) (hexAesKey string)
}

type Decoder interface {
	Decode(otp string) (*YubikeyOTP, error)
}

// NewDecoder takes function that resolves PublicName of Yubikey and returns hex encoded AES key of it and returns decoder that returns YubikeyOTP structures when presented with valid OTP
func NewDecoder(keystore func(publicName string) (hexAesKey string)) Decoder {
	var d OTPDecoder
	d.keystore = keystore
	return &d
}


type ErrorNotFound struct {
}

func (e *ErrorNotFound) Error() string {
    return "key not found"
}

func (d *OTPDecoder)Decode(otp string) (*YubikeyOTP, error) {
	aeskey := d.keystore(otp[:12])
	if len(aeskey) == 0 {
		return nil, &ErrorNotFound{}
	}

	modciphertext := ModHex2Hex(otp[12:])
	if strings.Contains(modciphertext, "X") {
		return nil, fmt.Errorf("error when converting from modhex, X marks the wrong characters: %s", modciphertext)
	}
	plain, err := DecryptAES128ECB(aeskey,modciphertext)
	if err != nil {
		return nil, fmt.Errorf("decryption error: %s", err)
	}
	return DecodeOTP(plain)

}