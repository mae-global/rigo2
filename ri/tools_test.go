package ri

import (
	"testing"

	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ToolsSpecification(t *testing.T) {

	Convey("Specification", t, func() {

		Convey("float base", func() {
			info := Specification("float base")
			So(info, ShouldNotBeNil)

			So(info.Class, ShouldEqual, "")
			So(info.Type, ShouldEqual, "float")
			So(info.Name, ShouldEqual, "base")
			So(info.Size, ShouldEqual, 1)
		})

		Convey("constant string path", func() {
			info := Specification("constant string path")
			So(info, ShouldNotBeNil)

			So(info.Class, ShouldEqual, "constant")
			So(info.Type, ShouldEqual, "string")
			So(info.Name, ShouldEqual, "path")
			So(info.Size, ShouldEqual, 1)
		})

		Convey("uniform float[3] rgb", func() {
			info := Specification("uniform float[3] rgb")
			So(info, ShouldNotBeNil)

			So(info.Class, ShouldEqual, "uniform")
			So(info.Type, ShouldEqual, "float")
			So(info.Name, ShouldEqual, "rgb")
			So(info.Size, ShouldEqual, 3)
		})

		Convey("reference float base", func() {
			info := Specification("reference float base")
			So(info, ShouldNotBeNil)

			So(info.Class, ShouldEqual, "reference")
			So(info.Type, ShouldEqual, "float")
			So(info.Name, ShouldEqual, "base")
			So(info.Size, ShouldEqual, 1)
		})
	})
}

func concat(in []RtPointer) string {
	out := ""
	for i, p := range in {
		if i > 0 {
			out += " "
		}
		out += p.String()
	}
	return out
}

func Test_ToolsParseBegin(t *testing.T) {

	Convey("ParseBegin", t, func() {

		Convey("normal .rib file", func() {
			out, err := ParseBegin("normal.rib")
			So(err, ShouldBeNil)
			So(concat(out), ShouldEqual, `"normal.rib"`)
		})

		Convey("null and -", func() {
			out, err := ParseBegin(NULL)
			So(err, ShouldBeNil)
			So(concat(out), ShouldEqual, `"out.rib"`)
		})

		Convey("stdout", func() {
			out, err := ParseBegin("stdout")
			So(err, ShouldBeNil)
			So(concat(out), ShouldEqual, `"|" "stdout"`)
		})

		Convey("render", func() {
			out, err := ParseBegin("render")
			So(err, ShouldBeNil)
			So(concat(out), ShouldEqual, `"|" "render"`)
		})

		Convey("render to file", func() {
			out, err := ParseBegin("render -capture test.rib")
			So(err, ShouldBeNil)
			So(concat(out), ShouldEqual, `"|" "render" "-capture" "test.rib"`)
		})

		Convey("catrib to stdout", func() {
			out, err := ParseBegin("catrib")
			So(err, ShouldBeNil)
			So(concat(out), ShouldEqual, `"|" "catrib"`)
		})

		Convey("catrib to file", func() {
			out, err := ParseBegin("catrib -o test.rib")
			So(err, ShouldBeNil)
			So(concat(out), ShouldEqual, `"|" "catrib" "-o" "test.rib"`)
		})

		/* TODO: future work
		Convey("stream over network",func() {
			out,err := ParseBegin("net tcp://192.168.1.1")
			So(err,ShouldBeNil)
			So(concat(out),ShouldEqual,`"|" "net" "-host" "tcp://192.168.1.1"`)
		})
		*/

	})
}
