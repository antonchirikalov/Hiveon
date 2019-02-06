package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func TestBlockClientConnection(t *testing.T) {
	Convey("Test IBlockRepository connection", t, func() {
		Convey("Trying to establish connection", func() {
			So(newBlockRepo, ShouldNotPanic)
		})
	})
}

func newBlockRepo(){
	NewBlockRepository()
}
