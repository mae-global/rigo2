package ri

import (
	"testing"

	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/smartystreets/goconvey/convey"
)


type testRiContext struct {
	last string
}

func (t *testRiContext) HandleError(err *RtError) *RtError { return err }
func (t *testRiContext) GenHandle(name,typeof string) (string,error) {	return name,nil }
func (t *testRiContext) Set(name,declaration string) RtToken {	return RtToken(declaration + " " + name) }
func (t *testRiContext) GetProgress() RtInt { return RtInt(100) }
func (t *testRiContext) GetLastRIB() string {	return t.last }
func (t *testRiContext) HandleV(name RtString,args,tokens,values []RtPointer) {
	t.Handle(List(name,args,Mix(tokens,values)))
}

func (t *testRiContext) Handle(list []RtPointer) {

	/* Do the minimal amount of work to generate a RIB output */
	name := list[0].(RtString)
	args := make([]RtPointer, 0)
	params := make([]RtPointer, 0)

	inannotations := false
	inargs := false
	inparams := false
	trigger := -1

	for i, param := range list {
		/* RtToken ---- */
		if p, ok := param.(RtToken); ok {

			switch string(p) {
			case string(ARGUMENTS):
				inparams = false
				inargs = true
				trigger = i
				break
			case string(PARAMETERLIST):
				inparams = true
				inargs = false
				trigger = i
				break
			case string(ANNOTATIONS):
				inannotations = true
			break
			}
		}

		if inannotations {
			break /* we don't need the annotations for testing */
		}

		if trigger == i {
			continue
		}

		if inargs {
			args = append(args, param)
		}

		if inparams {
			params = append(params, param)
		}
	}

	tokens, values := Unmix(params)

	RIBf := func(name RtString, args, tokens, values []RtPointer) string {

		switch string(name) {
		case "##":
			out := "##"
			for i, arg := range args {
				if str, ok := arg.(RtString); ok {
					if i > 0 {
						out += " "
					}
					out += string(str)
				}
			}
			return out
			break
		}	

		out := string(name)
		if len(args) > 0 {
			out += " " + Serialise(args, false)
		}

		if len(tokens) > 0 && len(values) > 0 {
			params := Mix(tokens, values)
			out += " " + Serialise(params, true)
		}
		return out
	}

	switch string(name) {
		case "Begin":
			t.last = ""
			/* FIXME */
		break
		case "End":
			t.last = ""
			/* FIXME */
		break
		default:
			t.last = RIBf(name,args,tokens,values)
		break
	}
}

func zero(n int) RtFloatArray {
	
	f := make([]RtFloat,n)
	return RtFloatArray(f)
}	


func Test_AllofRi(t *testing.T) {

	Convey("All of RenderMan Interface",t,func() {
	
			ri := Wrap(new(testRiContext))
			last := ri.Utils.GetLastRIB

			ri.Begin(NULL)
			So(last(),ShouldBeEmpty)

			ctx := ri.GetContext()
			So(ctx,ShouldNotBeNil)
			So(last(),ShouldBeEmpty)			
			
			ri.Context(ctx)
			So(last(),ShouldBeEmpty)



			ri.Attribute("displacementbound",RtToken("sphere"),RtFloat(2))
			So(last(),ShouldEqual,`Attribute "displacementbound" "sphere" [2]`)

			ri.Atmosphere("fog")
			So(last(),ShouldEqual,`Atmosphere "fog"`)
		
			
			light := ri.AreaLightSource("finite","test",RtToken("decayexponent"),RtFloat(.5),
									RtToken("lightcolor"),RtColor{.5,0,0},RtToken("intensity"),RtFloat(.6))
			So(string(light),ShouldEqual,"test")
			So(last(),ShouldEqual,`AreaLightSource "finite" "test" "decayexponent" [.5] "lightcolor" [.5 0 0] "intensity" [.6]`)

			ri.AttributeBegin()
			So(last(),ShouldEqual,`AttributeBegin`)

			ri.AttributeEnd()
			So(last(),ShouldEqual,`AttributeEnd`)

			ri.Basis(BSplineBasis,BSplineStep,RtBasis{-1,3,-3,1,3,-6,3,0,-3,3,0,0,1,0,0,0},1)
			So(last(),ShouldEqual,`Basis "b-spline" 1 "bezier" 1`)

			ri.Bound(RtBound{0,0.5,0,0.5,0.9,1})
			So(last(),ShouldEqual,`Bound [0 .5 0 .5 .9 1]`)

			ri.ConcatTransform(RtMatrix{2,0,0,0,  0,2,0,0,  0,0,2,0,  0,0,0,1})
			So(last(),ShouldEqual,`ConcatTransform [2 0 0 0 0 2 0 0 0 0 2 0 0 0 0 1]`)

			ri.CoordinateSystem("lamptop")
			So(last(),ShouldEqual,`CoordinateSystem "lamptop"`)

			ri.CoordSysTransform("lamptop")
			So(last(),ShouldEqual,`CoordSysTransform "lamptop"`)

			ri.Clipping(0.1,10000)
			So(last(),ShouldEqual,`Clipping .1 10000`)

			ri.ClippingPlane(3,0,0,0,0,-1)
			So(last(),ShouldEqual,`ClippingPlane 3 0 0 0 0 -1`)

			ri.CropWindow(0.0,0.3,0.0,0.5)
			So(last(),ShouldEqual,`CropWindow 0 .3 0 .5`)

			ri.Color(RtColor{0.2, 0.3, 0.9})
			So(last(),ShouldEqual,`Color [.2 .3 .9]`)

			ri.ColorSamples(1,RtFloatArray{.3,.3,.4},RtFloatArray{1,1,1})
			So(last(),ShouldEqual,`ColorSamples [.3 .3 .4] [1 1 1]`)

			So(ri.Declare("Np","uniform point"),ShouldEqual,RtToken("uniform point Np"))
			So(last(),ShouldEqual,`Declare "Np" "uniform point"`)

			ri.Detail(RtBound{10,20,42,69,0,1})
			So(last(),ShouldEqual,`Detail [10 20 42 69 0 1]`)

			ri.DetailRange(0,0,10,20)
			So(last(),ShouldEqual,`DetailRange 0 0 10 20`)

			ri.DepthOfField(22,45,1200)
			So(last(),ShouldEqual,`DepthOfField 22 45 1200`)

			ri.Display("pixar0","framebuffer","rgba",RtToken("int[2] origin"),RtIntArray{10,10})
			So(last(),ShouldEqual,`Display "pixar0" "framebuffer" "rgba" "int[2] origin" [10 10]`)

			ri.Displacement("displaceit")
			So(last(),ShouldEqual,`Displacement "displaceit"`)
		
			ri.End()
			So(last(),ShouldBeEmpty)

			ri.Exposure(1.5,2.3)
			So(last(),ShouldEqual,`Exposure 1.5 2.3`)

			ri.Exterior("fog")
			So(last(),ShouldEqual,`Exterior "fog"`)

			ri.FrameBegin(14)
			So(last(),ShouldEqual,`FrameBegin 14`)
			
			ri.FrameEnd()
			So(last(),ShouldEqual,`FrameEnd`)
	
			ri.Format(512,512,1)
			So(last(),ShouldEqual,`Format 512 512 1`)

			ri.FrameAspectRatio(4.0/3.0)
			So(last(),ShouldEqual,`FrameAspectRatio 1.333333`)

			ri.GeometricApproximation("flatness",2.5)
			So(last(),ShouldEqual,`GeometricApproximation "flatness" 2.5`)

			ri.GeneralPolygon(2,RtIntArray{4,3},P,RtFloatArray{0,0,0,0,1,0,0,1,1,0,0,1,
																											 0,.25,.5,0,.75,.75,0,.75,.25})

			So(last(),ShouldEqual,`GeneralPolygon [4 3] "vertex point P" [0 0 0 0 1 0 0 1 1 0 0 1 0 .25 .5 0 .75 .75 0 .75 .25]`)
		
			ri.Hider("paint")
			So(last(),ShouldEqual,`Hider "paint"`)

			ri.Identity()
			So(last(),ShouldEqual,`Identity`)

			ri.Illuminate(light,true)
			So(last(),ShouldEqual,`Illuminate "test" 1`)

			ri.Interior("water")
			So(last(),ShouldEqual,`Interior "water"`)

			ri.Imager("cmyk")
			So(last(),ShouldEqual,`Imager "cmyk"`)

			light = ri.LightSource("spotlight",RtLightHandle("test"),RtToken("coneangle"),RtInt(5))
			So(string(light),ShouldEqual,"test")
			So(last(),ShouldEqual,`LightSource "spotlight" "test" "coneangle" [5]`)

			ri.Matte(true)
			So(last(),ShouldEqual,`Matte 1`)

			ri.NuPatch(9,3,RtFloatArray{0,0,0,1,1,2,2,3,3,4,4,4},0,4,2,2,RtFloatArray{0,0,1,1},0,1)
			So(last(),ShouldEqual,`NuPatch 9 3 [0 0 0 1 1 2 2 3 3 4 4 4] 0 4 2 2 [0 0 1 1] 0 1`)


			ri.Option("limits",RtToken("gridsize"),RtInt(32),RtToken("bucketsize"),RtIntArray{12,12})
			So(last(),ShouldEqual,`Option "limits" "gridsize" [32] "bucketsize" [12 12]`)

			ri.Opacity(RtColor{0.5,1,1})
			So(last(),ShouldEqual,`Opacity [.5 1 1]`)

			ri.Orientation("lh")
			So(last(),ShouldEqual,`Orientation "lh"`)

			ri.Patch("bilinear",RtToken("P"),RtFloatArray{-.08,.04,.05, 0,.04,.05, -.08,.03,.05, 0,.03,.05})
			So(last(),ShouldEqual,`Patch "bilinear" "P" [-.08 .04 .05 0 .04 .05 -.08 .03 .05 0 .03 .05]`)

			ri.PatchMesh("bicubic",7,"nonperiodic",4,"nonperiodic",RtToken("P"),zero(28))
			So(last(),ShouldEqual,`PatchMesh "bicubic" 7 "nonperiodic" 4 "nonperiodic" "P" [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]`)
			


			ri.PointsPolygons(3,RtIntArray{3,3,3},RtIntArray{0,3,2,0,1,3,1,4,3},RtToken("P"),
												RtFloatArray{0,1,1,0,3,1,0,0,0,0,2,0,0,4,0},RtToken("Cs"),
												RtFloatArray{0,.3,.4,0,.3,.9,.2,.2,.5,.2,0,.9,.8})
			So(last(),ShouldEqual,`PointsPolygons [3 3 3] [0 3 2 0 1 3 1 4 3] "P" [0 1 1 0 3 1 0 0 0 0 2 0 0 4 0] "Cs" [0 .3 .4 0 .3 .9 .2 .2 .5 .2 0 .9 .8]`)

			ri.PointsGeneralPolygons(2,RtIntArray{2,2},RtIntArray{4,3,4,3},RtIntArray{0,1,4,3,6,7,8,1,2,5,4,9,10,11},
															 RtToken("P"),RtFloatArray{0,0,1,0,1,1,0,2,1,0,0,0,0,1,0,0,2,0,0,0.25,0.5,0.75,0.75,0.75,0.25,0,
																												 1.25,0.5,0,1.75,0.75,0,1.75,0.25})
			So(last(),ShouldEqual,`PointsGeneralPolygons [2 2] [4 3 4 3] [0 1 4 3 6 7 8 1 2 5 4 9 10 11] "P" [0 0 1 0 1 1 0 2 1 0 0 0 0 1 0 0 2 0 0 .25 .5 .75 .75 .75 .25 0 1.25 .5 0 1.75 .75 0 1.75 .25]`)

			ri.Projection(PERSPECTIVE,RtToken("fov"),RtFloat(45))
			So(last(),ShouldEqual,`Projection "perspective" "fov" [45]`)

			ri.PixelVariance(0.01)
			So(last(),ShouldEqual,`PixelVariance .01`)

			ri.PixelSamples(2,2)
			So(last(),ShouldEqual,`PixelSamples 2 2`)

			ri.PixelFilter(GaussianFilter,2.0,1.0)
			So(last(),ShouldEqual,`PixelFilter "gaussian" 2 1`)

			ri.Perspective(90.0)
			So(last(),ShouldEqual,`Perspective 90`)

			ri.Polygon(4,P,RtFloatArray{0,1,0,0,1,1,0,0,1,0,0,0})
			So(last(),ShouldEqual,`Polygon "vertex point P" [0 1 0 0 1 1 0 0 1 0 0 0]`)


			ri.Quantize(RGBA,2048,-1024,3071,1.0)
			So(last(),ShouldEqual,`Quantize "rgba" 2048 -1024 3071 1`)

			ri.RelativeDetail(0.6)
			So(last(),ShouldEqual,`RelativeDetail .6`)

			ri.ReverseOrientation()
			So(last(),ShouldEqual,`ReverseOrientation`)

			ri.Rotate(90.0,0.0,1.0,0.0)
			So(last(),ShouldEqual,`Rotate 90 0 1 0`)

			ri.Scale(0.5,1,1)
			So(last(),ShouldEqual,`Scale .5 1 1`)

			ri.Skew(45,0,1,0,1,0,0)
			So(last(),ShouldEqual,`Skew 45 0 1 0 1 0 0`)
				
			ri.Surface("wood",RtToken("roughness"),RtFloat(0.3),RtToken("Kd"),RtFloat(1.0),
												 RtToken("float ringwidth"),RtFloat(0.25))
			So(last(),ShouldEqual,`Surface "wood" "roughness" [.3] "Kd" [1] "float ringwidth" [.25]`)

			ri.ScreenWindow(-1,1,-1,1)
			So(last(),ShouldEqual,`ScreenWindow -1 1 -1 1`)

			ri.Shutter(0.1,0.9)
			So(last(),ShouldEqual,`Shutter .1 .9`)

			ri.ShadingRate(1.0)
			So(last(),ShouldEqual,`ShadingRate 1`)

			ri.ShadingInterpolation("smooth")
			So(last(),ShouldEqual,`ShadingInterpolation "smooth"`)

			ri.Sides(1)
			So(last(),ShouldEqual,`Sides 1`)

			/*
			ri.TransformPoints("current","lamptop",4,RtPointArray{RtPoint{0,0,0},RtPoint{1,1,1},
																													  RtPoint{2,2,2},RtPoint{3,3,3}})
			// FIXME check return points 
			*/

			ri.TrimCurve(1,RtIntArray{1},RtIntArray{3},RtFloatArray{0,0,0,1,1,2,2,3,3,4,4,4},RtFloatArray{0},RtFloatArray{4},RtIntArray{9},
									 RtFloatArray{1,1,1,0,0,0,1,1,1},RtFloatArray{.5,1,2,1,.5,0,0,0,.5},RtFloatArray{1,1,2,1,1,1,2,1,1})

			So(last(),ShouldEqual,`TrimCurve 1 [1] [3] [0 0 0 1 1 2 2 3 3 4 4 4] [0] [4] [9] [1 1 1 0 0 0 1 1 1] [.5 1 2 1 .5 0 0 0 .5] [1 1 2 1 1 1 2 1 1]`)
			
			ri.TransformBegin()
			So(last(),ShouldEqual,`TransformBegin`)
		
			ri.TransformEnd()
			So(last(),ShouldEqual,`TransformEnd`)

			ri.Translate(0.0,1.0,0.0)
			So(last(),ShouldEqual,`Translate 0 1 0`)

			ri.Transform(RtMatrix{0.5,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1})
			So(last(),ShouldEqual,`Transform [.5 0 0 0 0 1 0 0 0 0 1 0 0 0 0 1]`)

			ri.TextureCoordinates(0.0,0.0,2.0,-0.5,-0.5,1.75,3.0,3.0)
			So(last(),ShouldEqual,`TextureCoordinates 0 0 2 -.5 -.5 1.75 3 3`)
	


			ri.WorldBegin()
			So(last(),ShouldEqual,`WorldBegin`)

			ri.WorldEnd()
			So(last(),ShouldEqual,`WorldEnd`)

			
			



	})
}




