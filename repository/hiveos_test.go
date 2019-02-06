package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHiveosClientConnection(t *testing.T) {
	Convey("Test IHiveosRepository connection", t, func() {
		Convey("Trying to establish connection", func() {
			So(newHiveosRepo, ShouldNotPanic)
		})

	})
}

func newHiveosRepo() {
	NewHiveosRepository()
}
