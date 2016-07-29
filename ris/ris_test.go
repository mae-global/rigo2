package ris

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_RIS(t *testing.T) {

	Convey("RIS test", t, func() {
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
