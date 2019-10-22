package yubilib


// all hex-encoded
type testKey struct {
	Modhex string
	Ciphertext string
	Key string
	Plain string
	Publicname string
	SessionCounter uint8
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
		Publicname: "e7f7f99481c0",
		SessionCounter: 0x02,
		Low: 0x8816,
		High:0x24,
		Use: 0x00,
	},

}
