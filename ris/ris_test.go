package ris

import (
	"os"

	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var localtest bool /* if !localTest then on CI */

func init() {
	debug := os.Getenv("RIGO2_DEBUG")
	if debug == "testing" {
		localtest = false
	} else {
		localtest = true
	}
}


func Test_RIS(t *testing.T) {

	Convey("RIS test", t, func() {
		
		if !localtest {
			t.Skip()
		}

		Convey("Load Bxdf", func() {

			bxdf, err := Bxdf("PxrConstant")
			So(err, ShouldBeNil)
			So(bxdf, ShouldNotBeNil)

		})

		Convey("Load Integrator", func() {

			integrator, err := Integrator("PxrDefault")
			So(err, ShouldBeNil)
			So(integrator, ShouldNotBeNil)

		})

		Convey("Load LightFilter", func() {

			light, err := LightFilter("PxrGobo")
			So(err, ShouldBeNil)
			So(light, ShouldNotBeNil)

		})

		Convey("Load Projection", func() {

			projection, err := Projection("PxrCamera")
			So(err, ShouldBeNil)
			So(projection, ShouldNotBeNil)

		})

		Convey("Load Pattern", func() {

			pattern, err := Pattern("PxrBlend")
			So(err, ShouldBeNil)
			So(pattern, ShouldNotBeNil)
		})

		PrintStats()

	})
}
