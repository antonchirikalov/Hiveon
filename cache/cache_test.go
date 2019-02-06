package cache

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCacheLoad(t *testing.T) {
	Convey("Test cache load", t, func() {
		c := NewCache(5, 1)
		Convey("Put elements", func() {
			c.Store("1", "one")
			c.Store("2", "two")
			Convey("Then expect two values", func() {
				So(c.Count(), ShouldEqual, 2)
			})
		})
	})
}

func TestCacheExpiration(t *testing.T){
	Convey("Test elements to be expired", t, func() {
		c := NewCache(1, 1)
		Convey("Put elements", func() {
			c.Store("1", "one")
			c.Store("2", "two")
			Convey("Elements must be cleaned", func() {
				<-time.After(2 * time.Second)
				So(c.Count(), ShouldEqual, 0)
			})
		})
	})
}

func TestCacheHasElement(t *testing.T){
	Convey("Test elements presented", t, func() {
		c := NewCache(1, 1)
		Convey("Put element", func() {
			c.Store("1", "one")
			Convey("Expected to load one", func() {
				value, _ := c.Load("1")
				So(value.data,ShouldEqual,"one")
				value, ok := c.Load("Fake")
				So(ok, ShouldEqual, false)
			})
		})
	})
}


