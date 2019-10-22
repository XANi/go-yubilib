package yubilib

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

// decrypt using AES-128-ECB, used in yubikey OTP password
// expects hex-encoded strings as that is what YK tools store passwords as

func DecryptAES128ECB(hexKey string, hexCiphertext string) (plain []byte, err error) {
	key, err := hex.DecodeString(hexKey)
	if err != nil {return []byte{}, fmt.Errorf("error decoding key: %s", err)}
	data, err := hex.DecodeString(hexCiphertext)
	if err != nil {return []byte{}, fmt.Errorf("error decoding data: %s", err)}
	cipher, err := aes.NewCipher(key)
	if err != nil {return []byte{}, fmt.Errorf("error initializing AES: %s", err)}

    decrypted := make([]byte, len(data))
    size := 16

    for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
        cipher.Decrypt(decrypted[bs:be], data[bs:be])
    }
    return decrypted, nil
}

func EncryptAES128ECB(hexKey string, plain []byte) (encryptedHex string, err error) {
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		return "", fmt.Errorf("error decoding key: %s", err)
	}
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("error initializing AES: %s", err)
	}

	encrypted := make([]byte, len(plain))
	size := 16
	if len(plain) % size != 0 {
		return "", fmt.Errorf("plain data must be a multiple of %d", size)
	}

	for bs, be := 0, size; bs < len(plain); bs, be = bs+size, be+size {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return hex.EncodeToString(encrypted), nil
}