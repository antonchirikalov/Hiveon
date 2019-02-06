package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func  newMinerdashRepo(){
	m := NewMinerdashRepository()
	m.GetETHHashrate()
}

func  TestMinerdashClientConnection(t *testing.T) {
	Convey("Test IMinerdashRepository connection", t, func() {
		Convey("We need to send real querySingle as Ping doesn't guarantee connection", func() {
			So(newMinerdashRepo, ShouldNotPanic)
		})

	})
}
