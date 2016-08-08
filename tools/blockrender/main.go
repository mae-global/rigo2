package main

import (
	"fmt"
	"os"
	"log"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	"github.com/mae-global/rigo2/drivers"
	"github.com/mae-global/rigo2"
)

/* TODO: at the moment this is just development code, in the future this will
 * be a useful tool for render block diagrams of RIB files 
 */

func main() {

	rigo.AddDriver("block",builder)

	ri := rigo.New(&rigo.Configuration{Debug:true}) /* FIXME, custom driver only gets loaded when debug is set?! */

	/* Nice Simple example, should contain a single block, with two assets */
	
	/* use the acsii renderer */
	ri.Option("driver",RtToken("render"),RtString("acsii")) 	

	ri.Begin("block")
	
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
	
	/* More involved example */	

	/* use the svg renderer */
	ri.Option("driver",RtToken("render"),RtString("svg"))

	ri.Begin("block")
		ri.Display("redsphere.tiff", "file", "rgba")
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
			fmt.Fprintf(os.Stderr,"progress %03d%%\n", progress)
			if int(progress) < 100 {
				return false
			}
			return true
		})

	ri.End()

	rigo.RemoveDriver("block")
}


type custom struct {
	
	last string
}

func (c *custom) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) { }
func (c *custom) GetProgress() RtInt { return RtInt(100) }
func (c *custom) Handle(name RtString,args,tokens,values []RtPointer) *RtError {
	

	/* TODO: add all processing code here */


	return nil
} 

func (c *custom) Close() *RtError { return nil }
func (c *custom) GetLastRIB() string { return c.last }


/* custom driver builder */
func builder(logger *log.Logger, options []RtPointer, args ...string) (drivers.Driver,error) {
	
	logger.Printf("custom driver loaded\n")


	c := &custom{}
	
	return c,nil
}
