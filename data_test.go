package yubilib


// all hex-encoded
type testKey struct {
	Modhex string
	Ciphertext string
	Key string
	Plain string
	Publicname string
	Internalname string
	SessionCounter uint16
	Low uint16
	High uint8
	Use uint8
}

var goodKeys =  []testKey {
	{
		Modhex:     "hrjtdcrcrijvhfggcvdninchhujvhdcn",
		Ciphertext: "6c8d20c0c78f64550f2b7b066e8f620b",
		Key:        "486510b839d4613dbedddeb397b092f6",
		Plain:      "e7f7f99481c0020016882400fe07745c",
		Internalname: "e7f7f99481c0",
		Publicname: "ggfgfvkunief",
		SessionCounter: 0x02,
		Low: 0x8816,
		High:0x24,
		Use: 0x00,
	},
	{
		Modhex:     "jinnkirrlffcllbnvbgbdeefugekvrgt",
		Ciphertext: "87bb97cca440aa1bf1512334e539fc5d",
		Key:        "17cb015b54a22b297e4ea6beb65b8dc3",
		Plain:      "c0236b85f56c4a01fe3da3008bc6d83e",
		Internalname: "c0236b85f56c",
		Publicname: "ggfhjgvuiflg",
		SessionCounter: 0x014a,
		Low: 0x3dfe,
		High:0xa3,
		Use: 0x00,
	},

}
