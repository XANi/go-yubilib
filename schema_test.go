package yubilib

import (
	"testing"
	 . "github.com/smartystreets/goconvey/convey"
)

func TestSchema (t *testing.T) {
	Convey("schema table names", t ,func () {
		So(YubikeyKSM{}.TableName(),ShouldContainSubstring,"yubi")
		So(YubikeyOTP{}.TableName(),ShouldContainSubstring,"yubi")

	})
}
