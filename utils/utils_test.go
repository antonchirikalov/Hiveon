package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestViperConfig(t *testing.T) {
	Convey("Given yaml config ", t, func() {
		Convey("When config is read", func() {
			v := GetConfig()
			Convey("Values must be extracted", func() {
				So(v.GetString("sequelize2.database"), ShouldEqual, "hiveos_eth")
				So(v.GetString("influx.database"), ShouldEqual, "minerdash")
			})
		})
	})
}

func TestOSEnvVariables(t *testing.T) {
	Convey("Given OS variable", t, func() {
		os.Setenv("HIVEON-API_REDIS_PORT", "666")
		os.Setenv("HIVEON-API_OAUTH_INTROSPECT_URL", "test")
		Convey("When config is read", func() {
			Convey("Values must be extracted", func() {
				So(GetConfig().GetString("redis.port"), ShouldEqual, "666")
				So(GetConfig().GetString("oauth_introspect_url"), ShouldEqual, "test")
				os.Clearenv()
				So(GetConfig().GetString("hydraAdmin"), ShouldNotEqual, "test")
				So(GetConfig().GetString("postgres.port"), ShouldNotEqual, "666")
			})
		})
	})
}