/* machine generated
 * build tool version 0
 * generated on 2016-11-17 18:00:37.653664413 +0000 UTC
 * source /opt/pixar/RenderManProServer-21.2/include/ri.h
 * RiVersion 5
 */

package prototype

import (
	. "../core"
)

const (
	RIBVersion RtFloat = 3.04
	BloomFilterKeys int = 141
)


var (
	BloomFilterData = []int{ 0,0,0,0,0,0,0,0,1,1,0,0,0,0,1,0,0,0,1,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,2,0,0,1,0,1,0,0,0,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,0,1,0,1,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,1,0,2,0,0,0,0,0,0,0,2,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,1,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,2,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,2,0,1,0,0,0,0,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,1,1,2,0,0,0,1,0,0,0,1,1,0,1,3,1,0,0,1,0,0,1,0,0,0,0,0,1,0,0,2,0,0,1,0,0,0,0,0,0,1,1,0,0,0,0,0,0,1,0,0,0,0,1,1,1,1,0,1,2,0,2,1,0,0,1,0,0,0,0,1,0,1,0,0,1,0,1,0,1,0,0,0,2,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,1,0,0,1,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,1,0,1,0,0,1,1,0,0,0,0,0,0,0,1,0,1,1,0,0,0,0,0,1,1,0,1,1,1,0,0,0,1,0,1,0,0,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,1,0,1,0,0,1,0,0,1,0,1,0,0,0,0,0,0,1,0,0,0,1,1,1,1,0,0,0,0,0,0,2,0,0,0,2,0,0,1,0,0,1,2,1,1,0,1,0,0,0,0,0,1,1,0,1,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,1,0,0,0,0,0,0,1,2,0,0,0,1,1,0,1,0,0,0,2,0,1,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,0,1,0,0,0,1,0,0,1,0,0,1,0,0,0,0,0,1,0,1,1,1,0,0,1,0,0,0,0,1,0,0,1,0,0,1,0,0,1,1,0,0,1,0,0,0,0,1,0,0,0,1,0,2,0,1,0,0,1,1,0,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,1,0,0,2,1,1,1,0,0,0,1,0,1,0,1,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,2,0,0,1,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,1,0,0,0,0,1,0,0,0,1,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,1,0,0,0,0,0,0,1,0,0,1,1,0,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,1,0,0,0,1,0,0,1,0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,2,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,1,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,2,0,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,1,0,0,1,1,0,0,1,0,0,1,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,1,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,1,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,1,0,0,1,0,1,0,0,0,0,0,0,1,2,1,0,0,1,0,0,1,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,2,0,0,0,1,0,1,1,0,0,0,1,0,0,0,0,1,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,1,0,0,0,1,0,1,1,0,2,0,0,0,1,1,0,0,0,0,0,1,0,0,0,1,1,0,2,2,1,0,0,1,0,0,0,0,1,1,0,1,0,0,0,4,0,0,2,0,0,2,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,2,0,0,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,1,0,0,1,1,1,0,3,0,0,0,0,0,0,1,0,2,0,0,1,2,0,0,0,0,1,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,1,0,2,0,0,1,0,0,0,0,1,0,0,1,0,1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,1,2,1,0,0,0,0,1,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,1,0,0,0,1,0,0,1,0,1,0,0,0,1,0,0,0,0,0,2,1,0,1,0,1,0,2,0,0,1,1,0,1,0,0,0,0,0,1,0,1,0,2,0,0,0,0,0,0,1,0,1,0,0,1,0,1,1,0,1,0,0,0,2,0,0,0,1,0,1,0,1,0,0,0,1,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,1,0,0,0,1,0,0,1,1,0,0,1,0,1,0,1,0,2,0,0,0,0,1,0,0,1,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,1,1,0,0,1,0,0,0,1,0,1,0,1,2,2,0,0,0,0,1,0,2,0,1,0,1,1,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,1,1,0,1,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,1,1,0,0,0,0,0,0,1,2,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,2,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,2,0,0,0,1,2,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,1,1,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,1,1,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,1,2,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,1,1,1,0,0,0,0,0,0,0,0,0,1,0,0,2,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0, }

	Data string = "version floatArchiveBegin token name ...ArchiveEndArchiveRecord token type string format ...AreaLightSource token name ...Atmosphere token name ...Attribute token name ...AttributeBeginAttributeEndBasis basis u int ustep basis v int vstepBegin token nameBlobby int nleaf int ninst int* inst int nflt float* flt int nstr token[] str ...Bound bound boundBxdf token name token handle ...Camera token name ...Clipping float nearplane float farplaneClippingPlane float Nx float Ny float Nz float Px float Py float PzColor float* colorColorSamples int n float* nRGB float* RGBnConcatTransform matrix mCone float height float radius float tmax ...CoordSysTransform token fromspaceCoordinateSystem string nameCropWindow float left float right float bottom float topCurves token type int ncurves int* nvertices token wrap ...Cylinder float radius float zmin float zmax float tmax ...Declare string name string declDeformation token name ...DepthOfField float fstop float length float distanceDetail bound boundDetailRange float minvis float lotrans float hitrans float maxvisDisk float height float radius float tmax ...Displace token name token handle ...Displacement token name ...Display string name token type token mode ...DisplayChannel token channel ...DisplayFilter token name token handle ...EditAttributeBegin token nameEditAttributeEndEditBegin token name ...EditEndEditWorldBegin token name ...EditWorldEndElseElseIf string expr ...EnableLightFilter lighthandle light token filter boolean onoffEndExposure float gain float gammaExterior token name ...Format int xres int yres float pixel_arFrameAspectRatio float arFrameBegin int frameFrameEndGeneralPolygon int nloops int* nverts ...GeometricApproximation token type float valueGeometry token name ...Hider token name ...HierarchicalSubdivisionMesh token mask int nf int* nverts int* verts int nt token[] tags int* nargs int* intargs float* floatargs token[] stringargs ...Hyperboloid point point1 point point2 float tmax ...IdentityIfBegin string expr ...IfEndIlluminate lighthandle light boolean onoffImager token name ...Integrator token name token handle ...Interior token name ...Light token name token handle ...LightFilter token name token handle ...LightSource token name ...MakeBrickMap int nptcs string* ptcs string bkm ...MakeBump string pic string text token swrap token twrap filterfunc filt float swidth float twidth ...MakeCubeFaceEnvironment string px string nx string py string ny string pz string nz string text float fov filterfunc filt float swidth float twidth ...MakeLatLongEnvironment string pic string text filterfunc filt float swidth float twidth ...MakeShadow string pic string text ...MakeTexture string pic string text token swrap token twrap filterfunc filt float swidth float twidth ...Matte boolean onoffMotionBegin int n ...MotionEndNuPatch int nu int uorder float* uknot float umin float umax int nv int vorder float* vknot float vmin float vmax ...ObjectBeginObjectEndObjectInstance objecthandle hOpacity float* colorOption token name ...Orientation token orientParaboloid float radius float zmin float zmax float tmax ...Patch token type ...PatchMesh token type int nu token uwrap int nv token vwrap ...Pattern token name token handle ...Perspective float fovPixelFilter filterfunc func float xwidth float ywidthPixelSampleImager token name ...PixelSamples float x float yPixelVariance float varPoints int nverts ...PointsGeneralPolygons int npolys int* nloops int* nverts int* verts ...PointsPolygons int npolys int* nverts int* verts ...Polygon int nverts ...Procedural pointer data bound bound procsubdivfunc sfunc procfreefunc ffuncProcedural2 proc2subdivfunc sfunc proc2boundfunc bfunc ...Projection token name ...Quantize token type int one int min int max float ditherReadArchive token name archivecallback callback ...RelativeDetail float relResource token handle token type ...ResourceBeginResourceEndReverseOrientationRotate float angle float dx float dy float dzSampleFilter token name token handle ...Scale float sx float sy float szScopedCoordinateSystem string nameScreenWindow float left float right float bottom float topShader token name token handle ...ShadingInterpolation token typeShadingRate float sizeShutter float opentime float closetimeSides int nSkew float angle float d1x float d1y float d1z float d2x float d2y float d2zSolidBegin string opSolidEndSphere float radius float zmin float zmax float tmax ...SubdivisionMesh token mask int nf int* nverts int* verts int nt token[] tags int* nargs int* intargs float* floatargs ...Surface token name ...System string nameTextureCoordinates float s1 float t1 float s2 float t2 float s3 float t3 float s4 float t4Torus float majrad float minrad float phimin float phimax float tmax ...Transform matrix mTransformBeginTransformEndTransformPoints token fromspace token tospace int n point[] pointsTranslate float dx float dy float dzTrimCurve int nloops int* ncurves int* order float* knot float* min float* max int* n float* u float* v float* wVArchiveRecord token type string format va_list vapVPAtmosphere token name ...VPInterior token name ...VPSurface token name ...Volume token type bound bound int* dimensions ...VolumePixelSamples float x float yWorldBeginWorldEnd"

	Indices = []int{ 0,13,40,50,92,122,147,171,185,197,238,254,335,352,384,405,444,511,529,571,595,640,673,701,757,816,874,905,931,983,1001,1066,1111,1147,1174,1219,1251,1292,1321,1337,1361,1368,1397,1409,1413,1435,1497,1500,1531,1554,1593,1618,1638,1646,1687,1732,1755,1775,1927,1979,1987,2010,2015,2057,2078,2116,2139,2172,2211,2237,2287,2388,2539,2630,2667,2771,2790,2811,2820,2937,2948,2957,2986,3006,3027,3051,3111,3131,3193,3228,3249,3302,3334,3362,3385,3406,3477,3529,3551,3626,3684,3709,3765,3816,3840,3876,3889,3900,3918,3963,4003,4035,4069,4127,4161,4192,4214,4252,4263,4339,4359,4367,4423,4544,4566,4584,4674,4746,4764,4778,4790,4856,4892,5004,5055,5082,5107,5131,5180,5214,5224, }
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
