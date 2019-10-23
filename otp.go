package yubilib

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/efigence/go-yubilib/crc16"
	"strings"
)

var yubiCRC = crc16.MakeTable(
		crc16.Params{
			0x1021,
			0xffff,
			true,
			true,
			0x0000,
			0xf0b8, "YK/ISO13239",
		})



//  | field  | offset | size |
//  | uid    |      0 |    6 |
//  | useCtr |      6 |    2 |
//  | ts     |      8 |    3 |
//  | sesCtr |     11 |    1 |
//  | rnd    |     12 |    2 |
//  | crc    |     14 |    2 |

func DecodeOTP(decryptedOTPData []byte) (*YubikeyOTP,error) {
	var y YubikeyOTP
	y.SessionCounter =  binary.LittleEndian.Uint16(decryptedOTPData[6:8])
	y.YKTSLow =  binary.LittleEndian.Uint16(decryptedOTPData[8:10])
	y.YKTSHigh =  uint8(decryptedOTPData[10])
	y.UseCounter =  uint8(decryptedOTPData[11])
	y.internalName = hex.EncodeToString(decryptedOTPData[:6])
	if crc16.Checksum(decryptedOTPData,yubiCRC) != 0xf0b8 {
		return nil, fmt.Errorf("bad crc checksum")
	}
	return &y,nil

}

type OTPDecoder struct {
	keystore func(publicName string) (hexAesKey string, hexInternalName string)
}

type Decoder interface {
	Decode(otp string) (*YubikeyOTP, error)
}

// NewDecoder takes function that resolves PublicName of Yubikey to hex encoded AES key and internalname of it, and returns decoded that decrypts OTP content
func NewDecoder(keystore func(publicName string) (hexAesKey string,hexInternalName string)) Decoder {
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
	publicname := otp[:12]
	aeskey,internalname := d.keystore(publicname)
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
	out, err  :=  DecodeOTP(plain)
	if err != nil { return out, err }
	if out.internalName != internalname {
		return nil, fmt.Errorf("internalname does not match the key %s %s ", out.internalName, internalname)
	}
	out.PublicName = publicname
	return out, err



}