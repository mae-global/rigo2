/* machine generated
 * build tool version 0
 * generated on 2016-08-10 20:26:08.379494892 +0000 UTC
 * source /opt/pixar/RenderManProServer-20.8/include/ri.h
 * RiVersion 5
 */

package prototype

import (
	. "github.com/mae-global/rigo2/ri/core"
)

const (
	RIBVersion RtFloat = 3.04
	BloomFilterKeys int = 137
)


var (
	BloomFilterData = []int{ 0,1,0,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,1,1,0,0,0,0,1,0,1,1,0,0,0,0,0,0,1,0,2,0,0,1,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,0,1,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,2,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,1,1,0,0,0,0,1,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,1,0,0,0,1,0,0,0,0,0,0,0,0,1,0,0,0,0,1,1,0,0,0,1,0,0,1,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,1,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,1,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,2,0,1,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,1,1,1,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,2,0,0,0,0,1,0,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,1,0,4,1,0,1,0,0,0,0,1,2,0,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,1,1,1,0,1,2,0,0,0,1,0,0,0,0,0,0,1,2,1,0,0,1,0,0,1,0,0,0,0,0,0,0,0,1,0,0,2,0,0,0,0,1,0,1,0,0,0,0,0,0,0,2,0,1,0,0,0,1,1,0,0,1,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,2,0,0,0,0,0,0,0,1,0,1,0,0,1,1,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,2,1,0,2,1,1,1,0,0,1,0,2,0,0,0,0,0,0,0,1,1,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,1,0,1,0,0,0,1,0,1,1,0,0,0,0,0,0,0,2,0,0,0,1,1,1,3,0,0,0,0,0,0,0,0,0,0,2,0,0,1,0,0,1,2,1,1,2,0,0,0,0,0,0,1,0,0,0,0,1,2,1,0,0,1,0,2,0,0,0,0,0,0,1,1,0,0,0,0,0,1,1,0,0,0,1,1,0,1,0,1,1,2,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,0,0,1,0,0,1,1,1,0,1,0,0,1,0,0,1,0,0,0,0,0,1,0,0,1,2,0,0,0,0,0,0,0,1,0,0,1,0,0,1,0,0,2,1,0,0,0,0,0,0,0,1,0,0,0,0,0,2,0,1,0,0,1,0,0,1,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,1,0,0,0,1,1,1,0,0,0,1,0,1,0,0,2,0,0,0,0,0,0,0,0,3,1,2,1,0,0,0,1,1,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,2,0,0,1,0,0,0,1,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,1,0,1,0,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,1,0,1,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,0,1,0,1,0,0,1,0,1,0,1,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,0,1,0,0,1,0,0,0,1,0,0,3,0,0,0,0,0,1,0,1,0,0,0,0,0,0,2,1,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,1,0,1,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,0,0,0,0,0,0,0,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,2,1,0,0,1,0,0,1,0,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,2,0,0,0,0,0,0,1,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,1,0,1,1,0,1,0,0,2,1,0,0,1,1,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,2,1,0,0,1,0,0,1,0,0,0,1,0,1,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,2,0,0,0,2,1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,1,1,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,1,0,0,0,0,1,0,0,0,1,1,0,0,0,0,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,1,0,1,0,0,0,1,0,3,0,1,1,0,0,1,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,3,0,0,0,0,0,0,1,0,2,0,0,0,0,1,0,0,0,0,0,1,1,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,1,1,0,2,0,0,1,0,0,0,0,0,0,0,0,1,0,0,0,0,1,1,0,0,0,0,1,0,0,0,0,0,0,2,0,0,2,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,2,1,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,1,0,0,0,0,1,0,0,0,1,1,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,1,1,2,1,1,1,0,0,2,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,3,0,0,1,0,1,0,0,1,0,0,0,1,0,1,1,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,1,1,0,0,0,0,0,0,0,0,1,0,0,0,1,0,2,0,1,0,0,1,0,2,0,1,0,2,0,0,1,0,0,0,0,0,0,0,0,0,0,2,0,0,2,0,0,1,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,1,0,0,0,1,2,1,0,0,0,0,2,0,1,0,1,0,2,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,1,0,0,0,0,2,1,1,0,0,0,0,1,0,0,0,0,3,0,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,1,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,0,0,0,1,0,1,0,1,1,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0, }

	Data string = "version floatArchiveBegin token name ...ArchiveEndArchiveRecord token type string format ...AreaLightSource token name ...Atmosphere token name ...Attribute token name ...AttributeBeginAttributeEndBasis basis u int ustep basis v int vstepBegin token nameBlobby int nleaf int ninst int[] inst int nflt float[] flt int nstr token[] str ...Bound bound boundBxdf token name token handle ...Camera token name ...Clipping float nearplane float farplaneClippingPlane float Nx float Ny float Nz float Px float Py float PzColor float[] colorColorSamples int n float[] nRGB float[] RGBnConcatTransform matrix mCone float height float radius float tmax ...CoordSysTransform token fromspaceCoordinateSystem string nameCropWindow float left float right float bottom float topCurves token type int ncurves int[] nvertices token wrap ...Cylinder float radius float zmin float zmax float tmax ...Declare string name string declDeformation token name ...DepthOfField float fstop float length float distanceDetail bound boundDetailRange float minvis float lotrans float hitrans float maxvisDisk float height float radius float tmax ...Displacement token name ...Display string name token type token mode ...DisplayChannel token channel ...EditAttributeBegin token nameEditAttributeEndEditBegin token name ...EditEndEditWorldBegin token name ...EditWorldEndElseElseIf string expr ...EnableLightFilter lighthandle light token filter boolean onoffEndExposure float gain float gammaExterior token name ...Format int xres int yres float pixel_arFrameAspectRatio float arFrameBegin int frameFrameEndGeneralPolygon int nloops int[] nverts ...GeometricApproximation token type float valueGeometry token name ...Hider token name ...HierarchicalSubdivisionMesh token mask int nf int[] nverts int[] verts int nt token[] tags int[] nargs int[] intargs float[] floatargs token[] stringargs ...Hyperboloid point point1 point point2 float tmax ...IdentityIfBegin string expr ...IfEndIlluminate lighthandle light boolean onoffImager token name ...Integrator token name token handle ...Interior token name ...LightFilter token name token handle ...LightSource token name ...MakeBrickMap int nptcs string* ptcs string bkm ...MakeBump string pic string text token swrap token twrap filterfunc filt float swidth float twidth ...MakeCubeFaceEnvironment string px string nx string py string ny string pz string nz string text float fov filterfunc filt float swidth float twidth ...MakeLatLongEnvironment string pic string text filterfunc filt float swidth float twidth ...MakeShadow string pic string text ...MakeTexture string pic string text token swrap token twrap filterfunc filt float swidth float twidth ...Matte boolean onoffMotionBegin int n ...MotionEndNuPatch int nu int uorder float[] uknot float umin float umax int nv int vorder float[] vknot float vmin float vmax ...ObjectBeginObjectEndObjectInstance objecthandle hOpacity float[] colorOption token name ...Orientation token orientParaboloid float radius float zmin float zmax float tmax ...Patch token type ...PatchMesh token type int nu token uwrap int nv token vwrap ...Pattern token name token handle ...Perspective float fovPixelFilter filterfunc func float xwidth float ywidthPixelSampleImager token name ...PixelSamples float x float yPixelVariance float varPoints int nverts ...PointsGeneralPolygons int npolys int[] nloops int[] nverts int[] verts ...PointsPolygons int npolys int[] nverts int[] verts ...Polygon int nverts ...Procedural pointer data bound bound procsubdivfunc sfunc procfreefunc ffuncProcedural2 proc2subdivfunc sfunc proc2boundfunc bfunc ...Projection token name ...Quantize token type int one int min int max float ditherReadArchive token name archivecallback callback ...RelativeDetail float relResource token handle token type ...ResourceBeginResourceEndReverseOrientationRotate float angle float dx float dy float dzScale float sx float sy float szScopedCoordinateSystem string nameScreenWindow float left float right float bottom float topShader token name token handle ...ShadingInterpolation token typeShadingRate float sizeShutter float opentime float closetimeSides int nSkew float angle float d1x float d1y float d1z float d2x float d2y float d2zSolidBegin string opSolidEndSphere float radius float zmin float zmax float tmax ...SubdivisionMesh token mask int nf int[] nverts int[] verts int nt token[] tags int[] nargs int[] intargs float[] floatargs ...Surface token name ...System string nameTextureCoordinates float s1 float t1 float s2 float t2 float s3 float t3 float s4 float t4Torus float majrad float minrad float phimin float phimax float tmax ...Transform matrix mTransformBeginTransformEndTransformPoints token fromspace token tospace int n point[] pointsTranslate float dx float dy float dzTrimCurve int nloops int[] ncurves int[] order float[] knot float[] min float[] max int[] n float[] u float[] v float[] wVArchiveRecord token type string format va_list vapVPAtmosphere token name ...VPInterior token name ...VPSurface token name ...Volume token type bound bound int[] dimensions ...VolumePixelSamples float x float yWorldBeginWorldEnd"

	Indices = []int{ 0,13,40,50,92,122,147,171,185,197,238,254,337,354,386,407,446,513,532,576,600,645,678,706,762,822,880,911,937,989,1007,1072,1117,1144,1189,1221,1250,1266,1290,1297,1326,1338,1342,1364,1426,1429,1460,1483,1522,1547,1567,1575,1617,1662,1685,1705,1862,1914,1922,1945,1950,1992,2013,2051,2074,2113,2139,2189,2290,2441,2532,2569,2673,2692,2713,2722,2841,2852,2861,2890,2911,2932,2956,3016,3036,3098,3133,3154,3207,3239,3267,3290,3311,3385,3439,3461,3536,3594,3619,3675,3726,3750,3786,3799,3810,3828,3873,3905,3939,3997,4031,4062,4084,4122,4133,4209,4229,4237,4293,4419,4441,4459,4549,4621,4639,4653,4665,4731,4767,4888,4939,4966,4991,5015,5065,5099,5109, }
)


func RiBloomFilter() *BloomFilter {
	bits := make([]int,len(BloomFilterData))
	copy(bits,BloomFilterData)
	return &BloomFilter{bits,BloomFilterKeys}
}



func RiPrototypes() map[RtString]*Information {

	statements := make([]string,0)

	for i,_ := range Indices {
		if (i + 1) >= len(Indices) {
			idx := Indices[i]
			statements = append(statements,Data[ idx : ])
		} else {	
			idx := Indices[i]
			statements = append(statements,Data[ idx : Indices[i + 1] ])
		}
	}

	out := make(map[RtString]*Information,0)

	for _,statement := range statements {
		proto := Parse(statement)
		out[proto.Name] = proto
	}

	return out 
}	

/* EOF */
