# go-yubilib

[![GoDoc](https://godoc.org/github.com/efigence/go-yubilib?status.svg)](https://godoc.org/github.com/efigence/go-yubilib)


Primitives for operating with Yubikeys (mostly around OTP parts).

Assumptions:

* most stuff like keys are encoded in hex, to be compatible with the [reference](https://github.com/Yubico/yubikey-ksm-dpkg/blob/master/ykksm-utils.php) implementation
* structs have GORM/JSON/YAML annotations where appropriate so they can be used directly in GORM and/or API
* functions operating on modhex/hex just replace invalid characters with X. Normally that is either checked or just makes `hex` module fail on conversion so it is safe and fast to use inside the module but take care, giving wrong output to `ModHex2Hex` for example will return invalid hex string 
