package rigo

import (
	"fmt"
	"os"
	"testing"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/mae-global/rigo2/rie"
)

var localtest bool /* if !localTest then on CI */
var shorttest bool 

func init() {
	debug := os.Getenv("RIGO2_DEBUG")
	if debug == "testing" {
		localtest = false
	} else {
		localtest = true
	}
	short := os.Getenv("RIGO2_DEBUGSHORT")
	if short != "" {
		shorttest = true
	} 
}

func Test_RIBOptions(t *testing.T) {

	Convey("RIB Options",t,func() {

		ri := New(nil)
		
		ri.Option("rib",RtToken("format"),RtString("ascii"))
		ri.Option("rib",RtToken("asciistyle"),RtString("indented,wide"))	/* can be set with env RIASCIISTYLE also */
		ri.Option("rib",RtToken("compression"),RtString("gzip"))    /* can be set with env RICOMPRESSION also */
		ri.Option("rib",RtToken("percision"),RtInt(6))

		/* turn debugging on */
		ri.Option("driver",RtToken("debug"),RtBoolean(true))

		ri.Begin("out/rib_options.rib.gz")

		ri.WorldBegin()
			ri.Sphere(1,-1,1,360)
		ri.WorldEnd()

		ri.End()
	})
}

func Test_Callback(t *testing.T) {

	Convey("Context Callback",t,func() {
	
		count := 0

		/* define a custom callback handler, this gets called via a goroutine */
		customFrameBeginCallback := func(name RtString,args,tokens,values []RtPointer) bool {
			fmt.Fprintf(os.Stderr,"CustomCallbackHandler called -- %s\n",name)
			count++		

			return false /* tell the handler to ignore this particular call */
		}


		config := NewConfiguration()
		So(config,ShouldNotBeNil)
		So(config.Callbacks,ShouldNotBeNil)

		config.Callbacks["FrameBegin"] = customFrameBeginCallback

		ctx := NewContext(config)
		So(ctx,ShouldNotBeNil)
	
		ri := Wrap(ctx)
	
		ri.Option("driver",RtToken("int fragment"),RtInt(1))

		ri.Begin("tmp/out.rib")

		ri.FrameBegin(1)
		
		/* Because our custom handler ignores all FrameBegin calls we should expect a RenderMan RIB structure
		 * statement instead */
		So(ri.Utils.GetLastRIB(),ShouldEqual,`##RenderMan RIB`) 		
		ri.End()

		So(count,ShouldEqual,1)
		
	})
}
		



func Test_Context(t *testing.T) {

	Convey("Context", t, func() {

		ctx := NewContext(&Configuration{})
		So(ctx, ShouldNotBeNil)

		ri := Wrap(ctx)
		So(ri, ShouldNotBeNil)

		ri.Option("rib",RtToken("asciistyle"),RtString("wide")) /* default is "indent,wide", so for ease error checking we drop the indent */

		ri.Begin("tmp/context.rib")

		ri.FrameBegin(1)

		ri.Declare("name", "uniform string")
		ri.Declare("angle", "constant float")
		ri.Declare("base", "constant color")

		ri.Attribute("identifier", RtToken("name"), RtString("hero"))
		So(ri.Utils.GetLastRIB(), ShouldEqual, `Attribute "identifier" "name" ["hero"]`)

		ri.Attribute("identifier", RtToken("float[2] angle"), RtFloatArray{1.23, 32.1}, RtToken("float tint"), RtFloat(1.0))
		So(ri.Utils.GetLastRIB(), ShouldEqual, `Attribute "identifier" "float[2] angle" [1.23 32.1] "float tint" [1]`)

		ri.Attribute("identifier", RtToken("constant float[2] st"), RtFloatArray{1, 0.2})
		So(ri.Utils.GetLastRIB(), ShouldEqual, `Attribute "identifier" "constant float[2] st" [1 .2]`)

		ri.Attribute("identifier", RtToken("int size"), RtInt(3))
		So(ri.Utils.GetLastRIB(), ShouldEqual, `Attribute "identifier" "int size" [3]`)

		ri.Attribute("identifier", RtToken("int[3] numbers"), RtIntArray{1, 2, 3})
		So(ri.Utils.GetLastRIB(), ShouldEqual, `Attribute "identifier" "int[3] numbers" [1 2 3]`)

		ri.Attribute("identifier", RtToken("base"), RtColor{0, 0, 1})
		So(ri.Utils.GetLastRIB(), ShouldEqual, `Attribute "identifier" "base" [0 0 1]`)

		ri.FrameEnd()
		ri.End()

	})
	

	Convey("NotStarted Error", t, func() {

		/* setup an error handler and hook the error -- without this
		 * the default is to abort and throw a panic
		 */
		called := false

		handler := func(code, severity int, msg string) error {
			So(code, ShouldEqual, rie.NotStarted)
			So(severity, ShouldEqual, rie.Error)
			So(msg, ShouldNotBeEmpty)
			called = true
			return nil
		}

		/* throw a Not Started error by using RiFrameBegin
		 * before calling RiBegin
		 */
		ctx := NewContext(&Configuration{Errorf: handler})
		So(ctx, ShouldNotBeNil)

		ri := Wrap(ctx)
		So(ri, ShouldNotBeNil)
	

		ri.FrameBegin(1)
		So(called, ShouldBeTrue)

	})

	Convey("Calling End before Begin", t, func() {

		called := false

		handler := func(code, severity int, msg string) error {
			So(code, ShouldEqual, rie.NotStarted)
			So(severity, ShouldEqual, rie.Error)
			So(msg, ShouldNotBeEmpty)
			called = true
			return nil
		}

		ctx := NewContext(&Configuration{Errorf: handler})
		So(ctx, ShouldNotBeNil)

		ri := Wrap(ctx)
		So(ri, ShouldNotBeNil)

		ri.End()
		So(called, ShouldBeTrue)
	})

	Convey("NULL and '-' defaults to out.rib", t, func() {

		ctx := NewContext(nil)
		So(ctx, ShouldNotBeNil)

		ri := Wrap(ctx)
		So(ri, ShouldNotBeNil)

		/* TODO: clean up here */

		ri.Begin(NULL)
		ri.End()

		/* TODO: check file has been created, then clean up */

		ri.Begin("-")
		ri.End()

		/* TODO: check file has been created, then clean up */

	})

	Convey("stdout", t, func() {

		ctx := NewContext(nil)
		So(ctx, ShouldNotBeNil)

		ri := Wrap(ctx)
		So(ri, ShouldNotBeNil)

		ri.Begin("stdout")
		ri.End()
	})

	Convey("catrib", t, func() {
		if !localtest {
			t.Skip()
		}
		ri := Wrap(NewContext(nil))

		last := ri.Utils.GetLastRIB

		ri.Option("rib",RtToken("asciistyle"),RtString("wide")) /* default is "indent,wide", so for ease error checking we drop the indent */

		ri.Begin("catrib")
		ri.Display("render_sphere.tiff", "multires", "rgba")
		So(last(), ShouldEqual, `Display "render_sphere.tiff" "multires" "rgba"`)
		ri.Format(320, 240, 1)
		So(last(), ShouldEqual, `Format 320 240 1`)
		ri.Projection("perspective", RtToken("float fov"), RtFloat(30))
		So(last(), ShouldEqual, `Projection "perspective" "float fov" [30]`)
		ri.Translate(0, 0, 6)
		So(last(), ShouldEqual, `Translate 0 0 6`)
		ri.WorldBegin()
		So(last(), ShouldEqual, `WorldBegin`)
		ri.LightSource("ambientlight", "ambient", RtToken("float intensity"), RtFloat(0.5))
		So(last(), ShouldEqual, `LightSource "ambientlight" "ambient" "float intensity" [.5]`)
		ri.Color(RtColor{1, 0, 0})
		So(last(), ShouldEqual, `Color [1 0 0]`)
		ri.Sphere(1, -1, 1, 360)
		So(last(), ShouldEqual, `Sphere 1 -1 1 360`)
		ri.WorldEnd()
		So(last(), ShouldEqual, `WorldEnd`)
		ri.End()
	})

	/* Useful example of parsing through a strict RIB parser */
	Convey("", t, func() {
		if !localtest {
			t.Skip()
		}
		ri := Wrap(NewContext(nil))
		
		ri.Option("rib",RtToken("asciistyle"),RtString("wide")) /* default is "indent,wide", so for ease error checking we drop the indent */

		ri.Begin("catrib -o tmp/catribtofile.rib")
		ri.Display("render_sphere.tiff", "multires", "rgba")
		ri.Format(320, 240, 1)
		ri.Projection("perspective", RtToken("float fov"), RtFloat(30))
		ri.Translate(0, 0, 6)
		ri.WorldBegin()
		ri.LightSource("ambientlight", "-", RtToken("float intensity"), RtFloat(0.5))
		ri.Color(RtColor{1, 0, 0})
		ri.Sphere(1, -1, 1, 360)
		ri.WorldEnd()
		ri.End()
	})

	Convey("render", t, func() {
		if !localtest {
			t.Skip()
		}
		ri := Wrap(NewContext(&Configuration{}))
		ri.Begin("render -progress")
		ri.Display("tmp/redsphere.tiff", "file", "rgba")
		ri.Format(320, 240, 1)
		ri.Projection("perspective", RtToken("float fov"), RtFloat(30))
		ri.Translate(0, 0, 6)
		ri.WorldBegin()
		ri.LightSource("ambientlight", "-", RtToken("float intensity"), RtFloat(0.5))
		ri.Color(RtColor{1, 0, 0})
		ri.Sphere(1, -1, 1, 360)
		ri.WorldEnd()

		/* wait on the render progress to end */
		ri.Utils.Wait(func(progress RtInt) bool {
			fmt.Printf("progress %03d%%\n", progress)
			if int(progress) < 100 {
				return false
			}
			return true
		})

		ri.Display("tmp/redsphere2.tiff", "file", "rgba")
		ri.Format(320, 240, 1)
		ri.Projection("perspective", RtToken("float fov"), RtFloat(30))
		ri.Translate(0, 0, 6)
		ri.WorldBegin()
		ri.LightSource("ambientlight", "-", RtToken("float intensity"), RtFloat(1.0))
		ri.Color(RtColor{1, .2, .2})
		ri.Sphere(1, -1, 1, 360)
		ri.WorldEnd()

		ri.Utils.Wait(func(progress RtInt) bool {
			fmt.Printf("progress %03d%%\n", progress)
			if int(progress) < 100 {
				return false
			}
			return true
		})

		ri.End()
	})

	Convey("render to file", t, func() {

		/* requires renderer (usually prman) to be installed, skip if not localtest */
		if !localtest || shorttest {
			t.Skip()
		}

		ri := Wrap(NewContext(&Configuration{}))

		ri.Begin("render -capture out/rendertofile.rib")
		ri.Display("tmp/redsphere.tiff", "file", "rgba")
		ri.Format(320, 240, 1)
		ri.Projection("perspective", RtToken("float fov"), RtFloat(30))
		ri.Translate(0, 0, 6)
		ri.WorldBegin()
		ri.LightSource("ambientlight", "-", RtToken("float intensity"), RtFloat(0.5))
		ri.Color(RtColor{1, 0, 0})
		ri.Sphere(1, -1, 1, 360)
		ri.WorldEnd()
		ri.End()
	})

	Convey("long render", t, func() {
		if !localtest {
			t.Skip()
		}

		ctx := NewContext(nil)
		ri := Wrap(ctx)

		ri.Begin("render -progress")
		ri.Display("tmp/lredsphere.tiff", "file", "rgba")
		ri.Format(320, 240, 1)
		ri.ShadingRate(100)
		ri.ShadingInterpolation("smooth")
		ri.PixelVariance(0.001)
		ri.PixelFilter(GaussianFilter, RtFloat(3), RtFloat(3))
		ri.Projection("perspective", RtToken("float fov"), RtFloat(30))

		ri.Imager("background", RtToken("color color"), RtColor{1, 1, 1}, RtToken("float aplha"), RtFloat(1))

		ri.Hider("raytrace", RtToken("int maxsamples"), RtInt(64), RtToken("int minsamples"), RtInt(16),
			RtToken("int incremental"), RtInt(12), RtToken("float[4] aperture"), RtFloatArray{0, 0, 0, 0},
			RtToken("string integrationmode"), RtString("path"))

		ri.Integrator("PxrPathTracer", "redbox", RtToken("int maxPathLength"), RtInt(20),
			RtToken("string sampleMOD"), RtString("bxdf"), RtToken("int rouletteDepth"), RtInt(12),
			RtToken("float rouletteThreshold"), RtFloat(0.2), RtToken("int clampDepth"), RtInt(2),
			RtToken("int clampLuminance"), RtInt(10))

		ri.FrameBegin(1)
			ri.WorldBegin()
				ri.Translate(0, 0, 15)

				ri.AttributeBegin()
					ri.Translate(5, 5, 5)

						light := ri.AreaLightSource("PxrStdEnvDayLight", RtLightHandle("-"), RtToken("float importance"), RtFloat(2),
							RtToken("float exposure"), RtFloat(1), RtToken("vector directionVector"),
							RtVector{1, 1, 1}, RtToken("color specAmount"), RtColor{.5, .5, .5},
							RtToken("float haziness"), RtFloat(1.7), RtToken("float enableShadows"),
							RtFloat(1))

					ri.Bxdf("PxrLightEmission", RtToken(light))
					ri.Geometry("envsphere", RtToken("constant float[2] resolution"), RtFloatArray{1024, 1024})
				ri.AttributeEnd()

				ri.Illuminate(light, true)

				ri.AttributeBegin()
					ri.Bxdf("PxrDisney", "-", RtToken("color baseColor"), RtColor{1, 0, 0})
					ri.Sphere(1, -1, 1, 360)
				ri.AttributeEnd()
			ri.WorldEnd()
		ri.FrameEnd()

		/* wait on the render progress to end */
		ri.Utils.Wait(func(progress RtInt) bool {
			fmt.Printf("progress %03d%%\n", progress)
			if int(progress) < 100 {
				return false
			}
			return true
		})

		ri.End()

		/* do some statistics */
		stats := ctx.PeelStatistics()

		fmt.Fprintf(os.Stderr,stats.PrettyPrint())

	})

}
