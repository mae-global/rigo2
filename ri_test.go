package rigo

import (
	"testing"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Examples(t *testing.T) {

	Convey("Unit Cube Example", t, func() {

		ri := New(nil)

		ri.Begin("out/examples/unitcube.rib")
		ri.AttributeBegin()
		ri.Attribute("identifier", RtToken("name"), RtToken("unitcube"))
		ri.Bound(RtBound{-.5, .5, -.5, .5, -.5, .5})
		ri.TransformBegin()

		points := RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5}

		ri.ArchiveRecord(COMMENT, "far face")
		ri.Polygon(4, P, points)
		ri.Rotate(90, 0, 1, 0)

		ri.ArchiveRecord(COMMENT, "right face")
		ri.Polygon(4, P, points)
		ri.Rotate(90, 0, 1, 0)

		ri.ArchiveRecord(COMMENT, "near face")
		ri.Polygon(4, P, points)
		ri.Rotate(90, 0, 1, 0)

		ri.ArchiveRecord(COMMENT, "left face")
		ri.Polygon(4, P, points)

		ri.TransformEnd()
		ri.TransformBegin()

		ri.ArchiveRecord(COMMENT, "bottom face")
		ri.Rotate(90, 1, 0, 0)
		ri.Polygon(4, P, points)

		ri.TransformEnd()
		ri.TransformBegin()

		ri.ArchiveRecord(COMMENT, "top face")
		ri.Rotate(-90, 1, 0, 0)
		ri.Polygon(4, P, points)

		ri.TransformEnd()
		ri.AttributeEnd()
		ri.End()
	})

	Convey("Simple Sphere Example", t, func() {

		ri := New(nil)		

		ri.Begin("out/examples/simplesphere.rib")
		ri.Display("simplesphere.tiff", "file", "rgba")
		ri.Format(320, 240, 1)
		ri.Projection(PERSPECTIVE, RtToken("float fov"), RtFloat(30))
		ri.Translate(0, 0, 6)
		ri.WorldBegin()
		ri.LightSource("ambientlight", "-", RtToken("float intensity"), RtFloat(0.5))
		ri.Color(RtColor{1, 0, 0})
		ri.Sphere(1, -1, 1, 360)
		ri.WorldEnd()
		ri.End()
	})

}
