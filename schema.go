package yubilib

import "time"

// YubikeyKSM is GORM schema with info used by KSM, mostly similar to what Yubico tooling uses.
type YubikeyKSM struct {
	SerialNr int `gorm:"UNIQUE,column:serialnr"`
	PublicName string `gorm:"PRIMARY_KEY;column:publicname"`
	CreatedAt time.Time `gorm:"column:created"`
	UpdatedAt time.Time  `gorm:"column:modified"`
	InternalName string  `gorm:"column:internalname"`
	AESKey string  `gorm:"column:aeskey;size:32"`
	LockCode string `gorm:"column:lockcode"`
	Creator string `gorm:"column:creator"`
	Active bool  `gorm:"column:active"`
	// whether it is hardware key
	Hardware bool `gorm:"column:static"`
	// static key
	Static string `gorm:"column:static"`
}


//YubikeyOTP is

type YubikeyOTP struct {
  PublicName  string `gorm:"PRIMARY_KEY;column:ykpublicname"`
  Active bool `gorm:"column:active"`
  // names same as in gorm.Model
  CreatedAt time.Time `gorm:"column:created"`
  UpdatedAt time.Time `gorm:"column:modified"`
  // increments after first generation after power on
  // incerements if use counter overflows
  SessionCounter uint16 `gorm:"column:yk_counter"`
  // increments by 1 each token generation after first (0x00-0xff), overflows into session counter
  UseCounter uint8`gorm:"column:yk_use"`
  // some internal YK stuff
  YKTSLow uint16 `gorm:"column:yk_low"`
  YKTSHigh uint8 `gorm:"column:yk_high"`
  Notes string `gorm:"column:notes"`
}

