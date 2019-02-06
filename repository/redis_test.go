package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func newRedisRepo() {
	NewRedisRepository()
}

func TestRedisClientConnection(t *testing.T) {
	Convey("Test Redis client connection", t, func() {
		Convey("Trying to establish connection", func() {
			So(newRedisRepo, ShouldNotPanic)
		})
	})
}

func TestGetLatestWorker(t *testing.T){
    NewRedisRepository().GetLatestWorker("fbb80630a30bb959671f1dfa4546dc9b5fddbc6f")
}
