package rigo

import (
	"testing"
	"log"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	"github.com/mae-global/rigo2/drivers"
	"github.com/mae-global/rigo2/drivers/RIB"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Driver(t *testing.T) {

	/* construct a custom driver (we cheat by using an existing driver) */
	custom := func(logger *log.Logger, options []RtPointer, args ...string) (drivers.Driver,error) {
		return ribdriver.BuildStdoutDriver(logger,options,args...)			
	}

	Convey("Custom Driver test",t,func() {

		/* add a custom driver as "custom" */
		So(AddDriver("custom",custom),ShouldBeNil)	
	
		ri := New(nil)		

		ri.Begin("custom")

		ri.Display("simplesphere.tiff", "file", "rgba")
		ri.Format(320, 240, 1)
		ri.Projection(PERSPECTIVE, RtToken("float fov"), RtFloat(30))
		ri.Translate(0, 0, 6)
		ri.WorldBegin()
			ri.Light("PxrEnvDayLight", "-", RtToken("float intensity"), RtFloat(0.5))
			ri.Color(RtColor{1, 0, 0})
			ri.Sphere(1, -1, 1, 360)
		ri.WorldEnd()

		ri.End()	

		/* remove custom driver */
		So(RemoveDriver("custom"),ShouldBeNil)
	})
}

func Test_RigoDriver(t *testing.T) {

	/* This driver should write to tmp/out.go a formalised version as a program */

	Convey("Rigo Driver test",t,func() {

		ri := New(nil)
		ri.Begin("rigo tmp/out.go")

		ri.Display("simplesphere.exr","openexr","rgba")
		ri.Format(320,240,1)
		ri.Projection(PERSPECTIVE,RtToken("float fov"),RtFloat(30))
		ri.Translate(0,0,6)
		ri.WorldBegin()
			ri.Light("PxrEnvDayLight","-",RtToken("float intensity"),RtFloat(0.5))
			ri.Color(RtColor{1,0,0})
			ri.Sphere(1,-1,1,360)
		ri.WorldEnd()

		ri.End()

		/* TODO: check for outputted file */
	})
}








