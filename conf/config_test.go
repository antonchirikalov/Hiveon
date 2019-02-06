package conf__test

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"testing"
	"os"
)

func TestMySqlConfig(t *testing.T) {
	Convey("Given yaml config ", t, func() {
		//Check env
		config := "config.yaml"
		env := os.Getenv("hiveonEnv")
		if (len(env)) != 0 && ((env == "dev") || (env == "stage" )) {
			config = "config." + env + ".yaml"
		}else if (len(env) != 0) && (env != "stage") {
			config = "config.dev.yaml"
		}

		//fmt.Println(config)
		v := viper.New()
		v.SetConfigType("yaml")
		v.SetConfigFile(config)
		v.AddConfigPath("*")

		Convey("When config is read", func() {
			v.ReadInConfig()
			Convey("Values must be extracted", func() {
				So(v.GetString("sequelize2.database"), ShouldEqual, "hiveos_eth")
				So(v.GetString("influx.database"), ShouldEqual, "minerdash")
			})
		})
	})
}
