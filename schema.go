package yubilib

import "time"

// YubikeyKSM is GORM schema with info used by KSM, mostly similar to what Yubico tooling uses.
type YubikeyKSM struct {
	SerialNr     int       `gorm:"UNIQUE;column:serialnr;unique_index:serialnr_publicname" json:"serialnr"`
	PublicName   string    `gorm:"PRIMARY_KEY;column:publicname;unique_index:serialnr_publicname" json:"publicname"`
	CreatedAt    time.Time `gorm:"column:created"`
	UpdatedAt    time.Time `gorm:"column:modified"`
	InternalName string    `gorm:"column:internalname" json:"internalname"`
	AESKey       string    `gorm:"column:aeskey;size:32" json:"aeskey"`
	LockCode     string    `gorm:"column:lockcode" json:"lockcode"`
	Creator      string    `gorm:"column:creator" json:"creator"`
	Active       bool      `gorm:"column:active"`
	// whether it is hardware key
	Hardware bool `gorm:"column:hardware"`
	// static key if present/used
	Static string `gorm:"column:static" json:"static,omitempty"`
}
func (YubikeyKSM) TableName() string {
    return "yubiksm"
}


//YubikeyOTP is

type YubikeyOTP struct {
	PublicName   string `gorm:"PRIMARY_KEY;column:ykpublicname"`
	internalName string
	Active       bool `gorm:"column:active"`
	// names same as in gorm.Model
	CreatedAt time.Time `gorm:"column:created"`
	UpdatedAt time.Time `gorm:"column:modified"`
	// increments after first generation after power on
	// increments if use counter overflows, persistent
	SessionCounter uint16 `gorm:"column:yk_counter"`
	// increments by 1 each token generation after first (0x00-0xff), overflows into session counter
	UseCounter uint8 `gorm:"column:yk_use"`
	// 24 bit timestamp
	YKTSLow  uint16 `gorm:"column:yk_low"`
	YKTSHigh uint8  `gorm:"column:yk_high"`
	Notes    string `gorm:"column:notes"`
}

func (YubikeyOTP) TableName() string {
    return "yubiotp"
}
