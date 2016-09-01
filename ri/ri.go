package ri

import (
	"fmt"	
	"time"
	"io"
	"strings"

	. "github.com/mae-global/rigo2/ri/core"
)

/* This is just a wrapper, the guts are 
 * implemented/provided by RtContextHandle 
 */ 
type Ri struct {
	ctx RtContextHandle
}

/* RiArchiveBegin */
func (ri *Ri) ArchiveBegin(name RtToken, handle RtArchiveHandle, parameterlist ...RtPointer) RtArchiveHandle {
	h, err := ri.ctx.GenHandle(string(handle), "archive")
	if err != nil {
		if err := ri.ctx.HandleError(Error(2, 3, err.Error())); err != nil {
			panic(err)
		}
		h = string(handle)
	}

	ri.ctx.Handle(List("ArchiveBegin", []RtPointer{name, RtArchiveHandle(h)}, parameterlist))
	return RtArchiveHandle(h)
}

/* RiArchiveEnd */
func (ri *Ri) ArchiveEnd() {
	ri.ctx.Handle(List("ArchiveEnd", nil, nil))
}

/* RiArchiveRecord */
func (ri *Ri) ArchiveRecord(typeof RtToken, format string, parameterlist ...interface{}) {

	name := RtString("verbatim")

	switch string(typeof) {
	case "comment":
		name = RtString("#")
		break
	case "structure":
		name = RtString("##")
		break
	}

	out := fmt.Sprintf(format, parameterlist...)

	ri.ctx.Handle(List(name, []RtPointer{RtString(out)}, nil))
}

/* RiAreaLightSource */
func (ri *Ri) AreaLightSource(name RtToken, handle RtLightHandle, parameterlist ...RtPointer) RtLightHandle {

	h, err := ri.ctx.GenHandle(string(handle), "light")
	if err != nil {
		if err := ri.ctx.HandleError(Error(2, 3, err.Error())); err != nil {
			panic(err)
		}
		h = string(handle)
	}

	ri.ctx.Handle(AnnotatedList("AreaLightSource", []RtPointer{name, RtLightHandle(h)}, parameterlist,[]RtPointer{RtToken("lighthandle asset"),RtLightHandle(h)}))
	return RtLightHandle(h)
}

func (ri *Ri) AreaLightSourceV(name RtToken, handle RtLightHandle, n RtInt, tokens []RtToken, values []RtPointer) RtLightHandle {
	return ri.AreaLightSource(name, handle, ListParams(tokens, values)...)
}

/* RiAtmosphere */
func (ri *Ri) Atmosphere(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Atmosphere", []RtPointer{name}, parameterlist))
}

func (ri *Ri) AtmosphereV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Atmosphere(name, ListParams(tokens, values)...)
}

/* RiAttribute */
func (ri *Ri) Attribute(name string, parameterlist ...RtPointer) {
	
	/* check the parameterlist, if "identifier" is found then use that to annotate the block */
	var annotation []RtPointer

	if name == "identifier" && len(parameterlist) > 0 {
		params,values := Unmix(parameterlist)
		/* FIXME, check the balance of the parameterlist */
		for i,param := range params {	
			if token,ok := param.(RtToken); ok {
				info := Specification(string(token))	
				if info.Name == "name" {				
					annotation = []RtPointer{RtToken(info.ReplaceName("label")),values[i]}	
					break
				}
			}
		}
	}

	ri.ctx.Handle(AnnotatedList("Attribute", []RtPointer{RtString(name)}, parameterlist,annotation))
}

func (ri *Ri) AttributeV(name string, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Attribute(name, ListParams(tokens, values)...)
}

/* RiAttributeBegin */
func (ri *Ri) AttributeBegin() {
	ri.ctx.Handle(List("AttributeBegin", nil, nil))
}

/* RiAttributeEnd */
func (ri *Ri) AttributeEnd() {
	ri.ctx.Handle(List("AttributeEnd", nil, nil))
}

/* RiBasis */
func (ri *Ri) Basis(u RtBasis, ustep RtInt, v RtBasis, vstep RtInt) {

	out := make([]RtPointer,0)
	ub := false
	if u == BezierBasis && !ub {
		out = append(out,RtToken("bezier"))
		ub = true
	}

	if u == BSplineBasis && !ub {
		out = append(out,RtToken("b-spline"))
		ub = true
	}

	if u == CatmullRomBasis && !ub {
		out = append(out,RtToken("catmull-rom"))
		ub = true
	}

	if u == HermiteBasis && !ub {
		out = append(out,RtToken("hermite"))
		ub = true
	}

	if u == PowerBasis && !ub {
		out = append(out,RtToken("power"))
		ub = true
	}

	if !ub {
		out = append(out,u)
	}

	out = append(out,ustep)

	ub = false
	if v == BezierBasis && !ub {
		out = append(out,RtToken("bezier"))
		ub = true
	}

	if v == BSplineBasis && !ub {
		out = append(out,RtToken("b-spline"))
		ub = true
	}

	if v == CatmullRomBasis && !ub {
		out = append(out,RtToken("catmull-rom"))
		ub = true
	}

	if v == HermiteBasis && !ub {
		out = append(out,RtToken("hermite"))
		ub = true
	}

	if v == PowerBasis && !ub {
		out = append(out,RtToken("power"))
		ub = true
	}

	if !ub {
		out = append(out,v)
	}

	out = append(out,vstep)

	ri.ctx.Handle(List("Basis", out, nil))
}

/* RiBegin */
func (ri *Ri) Begin(name RtToken) {
	args, err := ParseBegin(name)
	if err != nil {
		if err := ri.ctx.HandleError(Error(2, 3, err.Error())); err != nil {
			panic(err)
		}
	}

	ri.ctx.Handle(List("Begin", args, nil))
}

/* RiBlobby */
func (ri *Ri) Blobby(nleaf RtInt, ninst RtInt, inst RtIntArray, nflt RtInt, flt RtFloatArray,
	nstr RtInt, str RtTokenArray, parameterlist ...RtPointer) {

	/* TODO: can check all the sizes -- maybe add a lazy mode where ninst = len(inst) */
	ri.ctx.Handle(List("Blobby", []RtPointer{nleaf, ninst, inst, nflt, flt, nstr, str}, parameterlist))
}

func (ri *Ri) BlobbyV(nleaf RtInt, ninst RtInt, inst RtIntArray, nflt RtInt, flt RtFloatArray,
	nstr RtInt, str RtTokenArray, n RtInt, tokens []RtToken, values []RtPointer) {

	ri.Blobby(nleaf, ninst, inst, nflt, flt, nstr, str, ListParams(tokens, values)...)
}

/* RiBound */
func (ri *Ri) Bound(bound RtBound) {
	ri.ctx.Handle(List("Bound", []RtPointer{bound}, nil))
}

/* RiBxdf */
func (ri *Ri) Bxdf(name, handle RtToken, parameterlist ...RtPointer) {

	h, err := ri.ctx.GenHandle(string(handle), "bxdf")
	if err != nil {
		if err := ri.ctx.HandleError(Error(2, 3, err.Error())); err != nil {
			panic(err)
		}
		h = string(handle)
	}

	ri.ctx.Handle(AnnotatedList("Bxdf", []RtPointer{name, RtToken(h)}, parameterlist,[]RtPointer{RtToken("token asset"),RtToken(h)}))
}

func (ri *Ri) BxdfV(name, handle RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Bxdf(name, handle, ListParams(tokens, values)...)
}

/* RiCamera */
func (ri *Ri) Camera(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Camera", []RtPointer{name}, parameterlist))
}

func (ri *Ri) CameraV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Camera(name, ListParams(tokens, values)...)
}

/* RiClipping */
func (ri *Ri) Clipping(nearplane RtFloat, farplane RtFloat) {
	ri.ctx.Handle(List("Clipping", []RtPointer{nearplane, farplane}, nil))
}

/* RiClippingPlane */
func (ri *Ri) ClippingPlane(Nx RtFloat, Ny RtFloat, Nz RtFloat, Px RtFloat, Py RtFloat, Pz RtFloat) {
	ri.ctx.Handle(List("ClippingPlane", []RtPointer{Nx, Ny, Nz, Px, Py, Pz}, nil))
}

/* RiColor */
func (ri *Ri) Color(color RtColor) {
	ri.ctx.Handle(List("Color", []RtPointer{color}, nil))
}

/* RiColorSamples */
func (ri *Ri) ColorSamples(n RtInt, nRGB RtFloatArray, RGBn RtFloatArray) {
	ri.ctx.Handle(List("ColorSamples", []RtPointer{nRGB, RGBn}, nil))
}

/* RiConcatTransform */
func (ri *Ri) ConcatTransform(m RtMatrix) {
	ri.ctx.Handle(List("ConcatTransform", []RtPointer{m}, nil))
}

/* RiCone */
func (ri *Ri) Cone(height RtFloat, radius RtFloat, tmax RtFloat, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Cone", []RtPointer{height, radius, tmax}, parameterlist))
}

func (ri *Ri) ConeV(height RtFloat, radius RtFloat, tmax RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Cone(height, radius, tmax, ListParams(tokens, values)...)
}

/* RiContext */
func (ri *Ri) Context(h RtContextHandle) {
	if h == nil {
		return
	}

	ri.ctx = h
}

func (ri *Ri) CoordSysTransform(fromspace RtToken) {
	ri.ctx.Handle(List("CoordSysTransform", []RtPointer{fromspace}, nil))
}

func (ri *Ri) CoordinateSystem(name string) {
	ri.ctx.Handle(List("CoordinateSystem", []RtPointer{RtString(name)}, nil))
}

func (ri *Ri) CropWindow(left, right, bottom, top RtFloat) {
	ri.ctx.Handle(List("CropWindow", []RtPointer{left, right, bottom, top}, nil))
}

func (ri *Ri) Curves(typeof RtToken, ncurves RtInt, nvertices RtIntArray, wrap RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Curves", []RtPointer{typeof, nvertices, wrap}, parameterlist))
}

func (ri *Ri) CurvesV(typeof RtToken, ncurves RtInt, nvertices RtIntArray, wrap RtToken,
	n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Curves(typeof, ncurves, nvertices, wrap, ListParams(tokens, values)...)
}

func (ri *Ri) Cylinder(radius, zmin, zmax, tmax RtFloat, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Cylinder", []RtPointer{radius, zmin, zmax, tmax}, parameterlist))
}

func (ri *Ri) CylinderV(radius, zmin, zmax, tmax RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Cylinder(radius, zmin, zmax, tmax, ListParams(tokens, values)...)
}

func (ri *Ri) Declare(name, decl string) RtToken {
	ri.ctx.Handle(List("Declare", []RtPointer{RtString(name), RtString(decl)}, nil))
	return ri.ctx.Set(name, decl)
}

func (ri *Ri) Deformation(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Deformation", []RtPointer{name}, parameterlist))
}

func (ri *Ri) DeformationV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Deformation(name, ListParams(tokens, values)...)
}

func (ri *Ri) DepthOfField(fstop, length, distance RtFloat) {
	ri.ctx.Handle(List("DepthOfField", []RtPointer{fstop, length, distance}, nil))
}

func (ri *Ri) Detail(bound RtBound) {
	ri.ctx.Handle(List("Detail", []RtPointer{bound}, nil))
}

func (ri *Ri) DetailRange(minvis, lotrans, hitrans, maxvis RtFloat) {
	ri.ctx.Handle(List("DetailRange", []RtPointer{minvis, lotrans, hitrans, maxvis}, nil))
}

func (ri *Ri) Disk(height, radius, tmax RtFloat, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Disk", []RtPointer{height, radius, tmax}, parameterlist))
}

func (ri *Ri) DiskV(height, radius, tmax RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Disk(height, radius, tmax, ListParams(tokens, values)...)
}

func (ri *Ri) Displacement(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Displacement", []RtPointer{name}, parameterlist))
}

func (ri *Ri) DisplacementV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Displacement(name, ListParams(tokens, values)...)
}

func (ri *Ri) Display(name string, typeof, mode RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Display", []RtPointer{RtString(name), typeof, mode}, parameterlist))
}

func (ri *Ri) DisplayChannel(channel RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("DisplayChannel", []RtPointer{channel}, parameterlist))
}

func (ri *Ri) DisplayChannelV(channel RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.DisplayChannel(channel, ListParams(tokens, values)...)
}

func (ri *Ri) EditAttributeBegin(name RtToken) {
	ri.ctx.Handle(List("EditAttributeBegin", []RtPointer{name}, nil))
}

func (ri *Ri) EditAttributeEnd() {
	ri.ctx.Handle(List("EditAttributeEnd", nil, nil))
}

func (ri *Ri) EditBegin(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("EditBegin", []RtPointer{name}, parameterlist))
}

func (ri *Ri) EditBeginV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.EditBegin(name, ListParams(tokens, values)...)
}

func (ri *Ri) EditEnd() {
	ri.ctx.Handle(List("EditEnd", nil, nil))
}

func (ri *Ri) EditWorldBegin(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("EditWorldBegin", []RtPointer{name}, parameterlist))
}

func (ri *Ri) EditWorldBeginV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.EditWorldBegin(name, ListParams(tokens, values)...)
}

func (ri *Ri) EditWorldEnd() {
	ri.ctx.Handle(List("EditWorldEnd", nil, nil))
}

func (ri *Ri) Else() {
	ri.ctx.Handle(List("Else", nil, nil))
}

func (ri *Ri) ElseIf(expr string, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("ElseIf", []RtPointer{RtString(expr)}, parameterlist))
}

func (ri *Ri) ElseIfV(expr string, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.ElseIf(expr, ListParams(tokens, values)...)
}

func (ri *Ri) EnableLightFilter(light RtLightHandle, filter RtToken, onoff RtBoolean) {
	ri.ctx.Handle(List("EnableLightFilter", []RtPointer{light, filter, onoff}, nil))
}
func (ri *Ri) End() {
	ri.ctx.Handle(List("End", nil, nil))
}
func (ri *Ri) ErrorHandler(handler RtErrorHandler) {
	ri.ctx.Handle(List("ErrorHandler", []RtPointer{handler}, nil))
}
func (ri *Ri) Exposure(gain, gamma RtFloat) {
	ri.ctx.Handle(List("Exposure", []RtPointer{gain, gamma}, nil))
}
func (ri *Ri) Exterior(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Exterior", []RtPointer{name}, parameterlist))
}

func (ri *Ri) ExteriorV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Exterior(name, ListParams(tokens, values)...)
}

func (ri *Ri) Format(xres, yres RtInt, pixel_ar RtFloat) {
	ri.ctx.Handle(List("Format", []RtPointer{xres, yres, pixel_ar}, nil))
}
func (ri *Ri) FrameAspectRatio(ar RtFloat) {
	ri.ctx.Handle(List("FrameAspectRatio", []RtPointer{ar}, nil))
}
func (ri *Ri) FrameBegin(frame RtInt) {

	/* NOTE, "label" is a special name for the block diagramming */
	ri.ctx.Handle(AnnotatedList("FrameBegin", []RtPointer{frame}, nil,[]RtPointer{RtToken("int label"),frame}))
}
func (ri *Ri) FrameEnd() {
	ri.ctx.Handle(List("FrameEnd", nil, nil))
}

func (ri *Ri) GeneralPolygon(nloops RtInt, nverts RtIntArray, parameterlist ...RtPointer) {
	if len(nverts) != int(nloops) {
		if err := ri.ctx.HandleError(Errorf(2,3,"value mismatch, expecting %d loops but only got %d",nloops,len(nverts))); err != nil {
			panic(err)
		}
	}
	ri.ctx.Handle(List("GeneralPolygon", []RtPointer{nverts}, parameterlist))
}

func (ri *Ri) GeneralPolygonV(nloops RtInt, nverts RtIntArray, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.GeneralPolygon(nloops, nverts, ListParams(tokens, values)...)
}

/* RiGetContext */
func (ri *Ri) GetContext() RtContextHandle {
	return ri.ctx
}

func (ri *Ri) GeometricApproximation(typeof RtToken, value RtFloat) {
	ri.ctx.Handle(List("GeometricApproximation", []RtPointer{typeof, value}, nil))
}

func (ri *Ri) Geometry(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Geometry", []RtPointer{name}, parameterlist))
}

func (ri *Ri) GeometryV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Geometry(name, ListParams(tokens, values)...)
}

func (ri *Ri) Hider(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Hider", []RtPointer{name}, parameterlist))
}

func (ri *Ri) HiderV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Hider(name, ListParams(tokens, values)...)
}

func (ri *Ri) HierarchicalSubdivisionMesh(mask RtToken, nf RtInt, nverts RtIntArray, verts RtIntArray,
	nt RtInt, tags RtTokenArray, nargs RtIntArray, intargs RtIntArray, floatargs RtFloatArray,
	stringargs RtTokenArray, parameterlist ...RtPointer) {

	ri.ctx.Handle(List("HieararchicalSubdivisionMesh", []RtPointer{mask, nf, nverts, verts, nt, tags,
		nargs, intargs, floatargs,
		stringargs}, parameterlist))
}

func (ri *Ri) HierarchicalSubdivisionMeshV(mask RtToken, nf RtInt, nverts RtIntArray, verts RtIntArray,
	nt RtInt, tags RtTokenArray, nargs RtIntArray, intargs RtIntArray, floatargs RtFloatArray,
	stringargs RtTokenArray, n RtInt, tokens []RtToken, values []RtPointer) {

	ri.HierarchicalSubdivisionMesh(mask, nf, nverts, verts, nt, tags, nargs, intargs,
		floatargs, stringargs, ListParams(tokens, values)...)
}

func (ri *Ri) Hyperboloid(point1, point2 RtPoint, tmax RtFloat, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Hyperboloid", []RtPointer{point1, point2, tmax}, parameterlist))
}

func (ri *Ri) HyperboloidV(point1, point2 RtPoint, tmax RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Hyperboloid(point1, point2, tmax, ListParams(tokens, values)...)
}

func (ri *Ri) Identity() {
	ri.ctx.Handle(List("Identity", nil, nil))
}

func (ri *Ri) IfBegin(expr string, parameterlist ...RtPointer) {
	ri.ctx.Handle(AnnotatedList("IfBegin", []RtPointer{RtString(expr)}, parameterlist,[]RtPointer{RtToken("string label"),RtString(expr)}))
}

func (ri *Ri) IfBeginV(expr string, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.IfBegin(expr, ListParams(tokens, values)...)
}

func (ri *Ri) IfEnd() {
	ri.ctx.Handle(List("IfEnd", nil, nil))
}

func (ri *Ri) Illuminate(light RtLightHandle, onoff RtBoolean) {
	ri.ctx.Handle(List("Illuminate", []RtPointer{light, onoff}, nil))
}

func (ri *Ri) Imager(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Imager", []RtPointer{name}, parameterlist))
}

func (ri *Ri) ImagerV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Imager(name, ListParams(tokens, values)...)
}

func (ri *Ri) Integrator(name, handle RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Integrator", []RtPointer{name, handle}, parameterlist))
}

func (ri *Ri) IntegratorV(name, handle RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Integrator(name, handle, ListParams(tokens, values)...)
}

func (ri *Ri) Interior(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Interior", []RtPointer{name}, parameterlist))
}

func (ri *Ri) InteriorV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Interior(name, ListParams(tokens, values)...)
}

func (ri *Ri) LightFilter(name, handle RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("LightFilter", []RtPointer{name, handle}, parameterlist))
}

func (ri *Ri) LightFilterV(name, handle RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.LightFilter(name, handle, ListParams(tokens, values)...)
}


/* LightSource -- depreciated in RenderMan 21.0 */
func (ri *Ri) LightSource(name RtToken, handle RtLightHandle, parameterlist ...RtPointer) RtLightHandle {
	h, err := ri.ctx.GenHandle(string(handle), "light")
	if err != nil {
		if err := ri.ctx.HandleError(Error(2, 3, err.Error())); err != nil {
			panic(err)
		}
		h = string(handle)
	}

	ri.ctx.Handle(AnnotatedList("LightSource", []RtPointer{name, RtLightHandle(h)}, parameterlist,[]RtPointer{RtToken("lighthandle asset"),RtLightHandle(h),RtToken("int depreciated"),RtInt(21)}))
	return RtLightHandle(h)
}

func (ri *Ri) LightSourceV(name RtToken, handle RtLightHandle, n RtInt, tokens []RtToken, values []RtPointer) RtLightHandle {
	return ri.LightSource(name, handle, ListParams(tokens, values)...)
}


/* Light -- added in RenderMan 21.0 */
func (ri *Ri) Light(name RtToken, handle RtLightHandle, parameterlist ...RtPointer) RtLightHandle {
	h, err := ri.ctx.GenHandle(string(handle),"light")
	if err != nil {
		if err := ri.ctx.HandleError(Error(2,3,err.Error())); err != nil {
			panic(err)
		}
		h = string(handle)
	}

	ri.ctx.Handle(AnnotatedList("Light",[]RtPointer{name,RtLightHandle(h)},parameterlist,[]RtPointer{RtToken("lighthandle asset"),RtLightHandle(h)}))
	return RtLightHandle(h)
}

func (ri *Ri) LightV(name RtToken, handle RtLightHandle, n RtInt, tokens []RtToken, values []RtPointer) RtLightHandle {
	return ri.Light(name,handle,ListParams(tokens,values)...)
}

func (ri *Ri) MakeBrickMap(nptcs RtInt, ptcs []string, bkm string, parameterlist ...RtPointer) {

	sa := make([]RtString, 0)
	for _, s := range ptcs {
		sa = append(sa, RtString(s))
	}

	ri.ctx.Handle(List("MakeBrickMap", []RtPointer{nptcs, RtStringArray(sa), RtString(bkm)}, parameterlist))
}

func (ri *Ri) MakeBrickMapV(nptcs RtInt, ptcs []string, bkm string, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.MakeBrickMap(nptcs, ptcs, bkm, ListParams(tokens, values)...)
}

func (ri *Ri) MakeBump(pic, text string, swrap, twrap RtToken, filt RtFilterFunc,
	swidth, twidth RtFloat, parameterlist ...RtPointer) {

	ri.ctx.Handle(List("MakeBump", []RtPointer{RtString(pic), RtString(text), twrap,
		filt, swidth, twidth}, parameterlist))
}

func (ri *Ri) MakeBumpV(pic, text string, swrap, twrap RtToken, filt RtFilterFunc,
	swidth, twidth RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {

	ri.MakeBump(pic, text, swrap, twrap, filt, swidth, twidth, ListParams(tokens, values)...)
}

func (ri *Ri) MakeCubeFaceEnvironment(px, nx, py, ny, pz, nz, text string,
	fov RtFloat, filt RtFilterFunc, swidth, twidth RtFloat, parameterlist ...RtPointer) {

	ri.ctx.Handle(List("MakeCubeFaceEnvironment", []RtPointer{RtString(px), RtString(nx), RtString(px),
		RtString(ny), RtString(pz), RtString(nz),
		RtString(text), fov, filt, swidth, twidth}, parameterlist))
}

func (ri *Ri) MakeCubeFaceEnvironmentV(px, nx, py, ny, pz, nz, text string,
	fov RtFloat, filt RtFilterFunc, swidth, twidth RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {

	ri.MakeCubeFaceEnvironment(px, nx, py, ny, pz, nz, text, fov, filt, swidth, twidth, ListParams(tokens, values)...)
}

func (ri *Ri) MakeLatLongEnvironment(pic, text string, filt RtFilterFunc, swidth, twidth RtFloat, parameterlist ...RtPointer) {

	ri.ctx.Handle(List("MakeLatLongEnvironment", []RtPointer{RtString(pic), RtString(text),
		filt, swidth, twidth}, parameterlist))
}

func (ri *Ri) MakeLatLongEnvironmentV(pic, text string, filt RtFilterFunc, swidth, twidth RtFloat,
	n RtInt, tokens []RtToken, values []RtPointer) {

	ri.MakeLatLongEnvironment(pic, text, filt, swidth, twidth, ListParams(tokens, values)...)
}

func (ri *Ri) MakeShadow(pic, text string, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("MakeShadow", []RtPointer{RtString(pic), RtString(text)}, parameterlist))
}

func (ri *Ri) MakeShadowV(pic, text string, n RtInt, tokens []RtToken, values []RtPointer) {

	ri.MakeShadow(pic, text, ListParams(tokens, values)...)
}

func (ri *Ri) MakeTexture(pic, text string, swrap, twrap RtToken,
	filt RtFilterFunc, swidth, twidth RtFloat, parameterlist ...RtPointer) {

	ri.ctx.Handle(List("MakeTexture", []RtPointer{RtString(pic), RtString(text), swrap,
		twrap, filt, swidth, twidth}, parameterlist))
}

func (ri *Ri) MakeTextureV(pic, text string, swrap, twrap RtToken,
	filt RtFilterFunc, swidth, twidth RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {

	ri.MakeTexture(pic, text, swrap, twrap, filt, swidth, twidth, ListParams(tokens, values)...)
}

func (ri *Ri) Matte(onoff RtBoolean) {
	ri.ctx.Handle(List("Matte", []RtPointer{onoff}, nil))
}

func (ri *Ri) MotionBegin(n RtInt, parameterlist ...RtPointer) {

	/* basically read the parameterlist early and drop all the args, we are
	 * only after the floatarray */
	array := RtFloatArray{}
	for _, param := range parameterlist {
		if a, ok := param.(RtFloatArray); ok {
			array = a
			break
		}
	}

	if len(array) != int(n) {
		if err := ri.ctx.HandleError(Errorf(2, 3, "data mismatch, expecting %d motion points, but got %d instead", int(n), len(array))); err != nil {
			panic(err)
		}
	}

	ri.ctx.Handle(List("MotionBegin", []RtPointer{array}, nil))
}

func (ri *Ri) MotionBeginV(n RtInt, values []RtPointer) {
	ri.MotionBegin(n, values...)
}

func (ri *Ri) MotionEnd() {
	ri.ctx.Handle(List("MotionEnd", nil, nil))
}

func (ri *Ri) NuPatch(nu, uorder RtInt, uknot RtFloatArray, umin, umax RtFloat,
	nv, vorder RtInt, vknot RtFloatArray, vmin, vmax RtFloat, parameterlist ...RtPointer) {

	ri.ctx.Handle(List("NuPatch", []RtPointer{nu, uorder, uknot, umin, umax, nv, vorder, vknot, vmin, vmax}, parameterlist))
}

func (ri *Ri) NuPatchV(nu, uorder RtInt, uknot RtFloatArray, umin, umax RtFloat,
	nv, vorder RtInt, vknot RtFloatArray, vmin, vmax RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {

	ri.NuPatch(nu, uorder, uknot, umin, umax, nv, vorder, vknot, vmin, vmax, ListParams(tokens, values)...)
}

func (ri *Ri) ObjectBegin() RtObjectHandle {	

	h, err := ri.ctx.GenHandle("-", "object")
	if err != nil {
		if err := ri.ctx.HandleError(Error(2, 3, err.Error())); err != nil {
			panic(err)
		}
		h = "-"
	}
	ri.ctx.Handle(List("ObjectBegin", nil, nil))

	return RtObjectHandle(h) /* FIXME */
}

func (ri *Ri) ObjectEnd() {
	ri.ctx.Handle(List("ObjectEnd", nil, nil))
}

func (ri *Ri) ObjectInstance(handle RtObjectHandle) {
	ri.ctx.Handle(List("ObjectInstance", []RtPointer{handle}, nil))
}

func (ri *Ri) Opacity(color RtColor) {
	ri.ctx.Handle(List("Opacity", []RtPointer{color}, nil))
}

func (ri *Ri) Option(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Option", []RtPointer{name}, parameterlist))
}

func (ri *Ri) OptionV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Option(name, ListParams(tokens, values)...)
}

func (ri *Ri) Orientation(orient RtToken) {
	ri.ctx.Handle(List("Orientation", []RtPointer{orient}, nil))
}

func (ri *Ri) Paraboloid(radius, zmin, zmax, tmax RtFloat, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Paraboloid", []RtPointer{radius, zmin, zmax, tmax}, parameterlist))
}

func (ri *Ri) ParaboloidV(radius, zmin, zmax, tmax RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Paraboloid(radius, zmin, zmax, tmax, ListParams(tokens, values)...)
}

func (ri *Ri) Patch(typeof RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Patch", []RtPointer{typeof}, parameterlist))
}

func (ri *Ri) PatchV(typeof RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Patch(typeof, ListParams(tokens, values)...)
}

func (ri *Ri) PatchMesh(typeof RtToken, nu RtInt,
	uwrap RtToken, nv RtInt, vwrap RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("PatchMesh", []RtPointer{typeof, nu, uwrap, nv, vwrap}, parameterlist))
}

func (ri *Ri) PatchMeshV(typeof RtToken, nu RtInt,
	uwrap RtToken, nv RtInt, vwrap RtToken, n RtInt, tokens []RtToken, values []RtPointer) {

	ri.PatchMesh(typeof, nu, uwrap, nv, vwrap, ListParams(tokens, values)...)
}

func (ri *Ri) Pattern(name, handle RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Pattern", []RtPointer{name, handle}, parameterlist))
}

func (ri *Ri) PatternV(name, handle RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Pattern(name, handle, ListParams(tokens, values)...)
}

func (ri *Ri) Perspective(fov RtFloat) {
	ri.ctx.Handle(List("Perspective", []RtPointer{fov}, nil))
}

func (ri *Ri) PixelFilter(f RtFilterFunc, xwidth, ywidth RtFloat) {
	ri.ctx.Handle(List("PixelFilter", []RtPointer{f, xwidth, ywidth}, nil))
}

func (ri *Ri) PixelSampleImager(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("PixelSampleImager", []RtPointer{name}, parameterlist))
}

func (ri *Ri) PixelSampleImagerV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.PixelSampleImager(name, ListParams(tokens, values)...)
}

func (ri *Ri) PixelSamples(x, y RtFloat) {
	ri.ctx.Handle(List("PixelSamples", []RtPointer{x, y}, nil))
}

func (ri *Ri) PixelVariance(variance RtFloat) {
	ri.ctx.Handle(List("PixelVariance", []RtPointer{variance}, nil))
}

func (ri *Ri) Points(nverts RtInt, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Points", nil, parameterlist))
}

func (ri *Ri) PointsV(nverts RtInt, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Points(nverts, ListParams(tokens, values)...)
}

func (ri *Ri) PointsGeneralPolygons(npolys RtInt, nloops, nverts, verts RtIntArray, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("PointsGeneralPolygons", []RtPointer{nloops, nverts, verts}, parameterlist))
}

func (ri *Ri) PointsGeneralPolygonsV(npolys RtInt, nloops, nverts, verts RtIntArray,
	n RtInt, tokens []RtToken, values []RtPointer) {

	ri.PointsGeneralPolygons(npolys, nloops, nverts, verts, ListParams(tokens, values)...)
}

func (ri *Ri) PointsPolygons(npolys RtInt, nverts, verts RtIntArray, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("PointsPolygons", []RtPointer{nverts, verts}, parameterlist))
}

func (ri *Ri) PointsPolygonsV(npolys RtInt, nverts, verts RtIntArray, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.PointsPolygons(npolys, nverts, verts, ListParams(tokens, values)...)
}

func (ri *Ri) Polygon(nverts RtInt, parameterlist ...RtPointer) {
	/* FIXME: check nverts agains input */
	ri.ctx.Handle(List("Polygon", nil, parameterlist))
}

func (ri *Ri) PolygonV(nverts RtInt, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Polygon(nverts, ListParams(tokens, values)...)
}

func (ri *Ri) Procedural(data RtPointer, bound RtBound, sfunc RtProcSubdivFunc, ffunc RtProcFreeFunc) {

	/*RtString args[] = { ”perl gencan.pl”, ”” };
	RtBound mybound = { -1, 1, -1, 1, 0, 6 };
	RiProcedural ((RtPointer)args, mybound, RiProcRunProgram, RiProcFree);
	Procedural ”RunProgram” [ ”perl gencan.pl” ”” ] [ -1 1 -1 1 0 6 ] */

	ri.ctx.Handle(AnnotatedList("Procedural", []RtPointer{sfunc, data, bound}, nil,[]RtPointer{RtToken("subdivfunc asset"),sfunc})) //[]RtPointer{data,bound,sfunc,ffunc},nil))
}

func (ri *Ri) Procedural2(sfunc RtProc2SubdivFunc, bfunc RtProc2BoundFunc, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Procedural2", []RtPointer{sfunc, bfunc}, parameterlist))
}

func (ri *Ri) Procedural2V(sfunc RtProc2SubdivFunc, bfunc RtProc2BoundFunc, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Procedural2(sfunc, bfunc, ListParams(tokens, values)...)
}

func (ri *Ri) Projection(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Projection", []RtPointer{name}, parameterlist))
}

func (ri *Ri) ProjectionV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Projection(name, ListParams(tokens, values)...)
}

func (ri *Ri) Quantize(typeof RtToken, one, min, max RtInt, dither RtFloat) {
	ri.ctx.Handle(List("Quantize", []RtPointer{typeof, one, min, max,dither}, nil))
}

func (ri *Ri) ReadArchive(name RtToken, callback RtArchiveCallback, parameterlist ...RtPointer) {
	
	params := []RtPointer{name}

	if callback != RtArchiveCallback("") && callback != RtArchiveCallback("-") {
		params = append(params,callback)
	}

	ri.ctx.Handle(List("ReadArchive", params, parameterlist))
}

func (ri *Ri) ReadArchiveV(name RtToken, callback RtArchiveCallback, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.ReadArchive(name, callback, ListParams(tokens, values)...)
}

func (ri *Ri) RelativeDetail(rel RtFloat) {
	ri.ctx.Handle(List("RelativeDetail", []RtPointer{rel}, nil))
}

func (ri *Ri) Resource(handle, typeof RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Resource", []RtPointer{handle, typeof}, parameterlist))
}

func (ri *Ri) ResourceV(handle, typeof RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Resource(handle, typeof, ListParams(tokens, values)...)
}

func (ri *Ri) ResourceBegin() {
	ri.ctx.Handle(List("ResourceBegin", nil, nil))
}

func (ri *Ri) ResourceEnd() {
	ri.ctx.Handle(List("ResourceEnd", nil, nil))
}

func (ri *Ri) ReverseOrientation() {
	ri.ctx.Handle(List("ReverseOrientation", nil, nil))
}

func (ri *Ri) Rotate(angle, dx, dy, dz RtFloat) {
	ri.ctx.Handle(List("Rotate", []RtPointer{angle, dx, dy, dz}, nil))
}

func (ri *Ri) Scale(sx, sy, sz RtFloat) {
	ri.ctx.Handle(List("Scale", []RtPointer{sx, sy, sz}, nil))
}

func (ri *Ri) ScopedCoordinateSystem(name string) {
	ri.ctx.Handle(List("ScopedCoordinateSystem", []RtPointer{RtString(name)}, nil))
}

func (ri *Ri) ScreenWindow(left, right, bottom, top RtFloat) {
	ri.ctx.Handle(List("ScreenWindow", []RtPointer{left, right, bottom, top}, nil))
}

func (ri *Ri) Shader(name, handle RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Shader", []RtPointer{name, handle}, parameterlist))
}

func (ri *Ri) ShaderV(name, handle RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Shader(name, handle, ListParams(tokens, values)...)
}

func (ri *Ri) ShadingInterpolation(typeof RtToken) {
	ri.ctx.Handle(List("ShadingInterpolation", []RtPointer{typeof}, nil))
}

func (ri *Ri) ShadingRate(size RtFloat) {
	ri.ctx.Handle(List("ShadingRate", []RtPointer{size}, nil))
}

func (ri *Ri) Shutter(opentime, closetime RtFloat) {
	ri.ctx.Handle(List("Shutter", []RtPointer{opentime, closetime}, nil))
}

func (ri *Ri) Sides(n RtInt) {
	ri.ctx.Handle(List("Sides", []RtPointer{n}, nil))
}

func (ri *Ri) Skew(angle, d1x, d1y, d1z, d2x, d2y, d2z RtFloat) {
	ri.ctx.Handle(List("Skew", []RtPointer{angle, d1x, d1y, d1z, d2x, d2y, d2z}, nil))
}

func (ri *Ri) SolidBegin(op string) {
	ri.ctx.Handle(List("SolidBegin", []RtPointer{RtString(op)}, nil))
}

func (ri *Ri) SolidEnd() {
	ri.ctx.Handle(List("SolidEnd", nil, nil))
}

func (ri *Ri) Sphere(radius, zmin, zmax, tmax RtFloat, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Sphere", []RtPointer{radius, zmin, zmax, tmax}, parameterlist))
}

func (ri *Ri) SphereV(radius, zmin, zmax, tmax RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Sphere(radius, zmin, zmax, tmax, ListParams(tokens, values)...)
}

func (ri *Ri) SubdivisionMesh(mask RtToken, nf RtInt, nverts, verts RtIntArray, nt RtInt,
	tags RtTokenArray, nargs, intargs RtIntArray, floatargs RtFloatArray, parameterlist ...RtPointer) {

	ri.ctx.Handle(List("SubdivisionMesh", []RtPointer{mask, nf, verts, nt, tags, nargs, intargs, floatargs}, parameterlist))
}

func (ri *Ri) SubdivisionMeshV(mask RtToken, nf RtInt, nverts, verts RtIntArray, nt RtInt,
	tags RtTokenArray, nargs, intargs RtIntArray, floatargs RtFloatArray, n RtInt, tokens []RtToken, values []RtPointer) {

	ri.SubdivisionMesh(mask, nf, nverts, verts, nt, tags, nargs, intargs, floatargs, ListParams(tokens, values)...)
}

func (ri *Ri) Surface(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Surface", []RtPointer{name}, parameterlist))
}

func (ri *Ri) SurfaceV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Surface(name, ListParams(tokens, values)...)
}

func (ri *Ri) System(name string) {
	ri.ctx.Handle(List("System", []RtPointer{RtString(name)}, nil))
}

func (ri *Ri) TextureCoordinates(s1, t1, s2, t2, s3, t3, s4, t4 RtFloat) {
	ri.ctx.Handle(List("TextureCoordinates", []RtPointer{s1, t1, s2, t2, s3, t3, s4, t4}, nil))
}
func (ri *Ri) Torus(majrad, minrad, phimin, phimax, tmax RtFloat, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Torus", []RtPointer{majrad, minrad, phimin, phimax, tmax}, parameterlist))
}
func (ri *Ri) TorusV(majrad, minrad, phimin, phimax, tmax RtFloat, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Torus(majrad, minrad, phimin, phimax, tmax, ListParams(tokens, values)...)
}

func (ri *Ri) Transform(m RtMatrix) {
	ri.ctx.Handle(List("Transform", []RtPointer{m}, nil))
}

func (ri *Ri) TransformBegin() {
	ri.ctx.Handle(List("TransformBegin", nil, nil))
}

func (ri *Ri) TransformEnd() {
	ri.ctx.Handle(List("TransformEnd", nil, nil))
}

func (ri *Ri) TransformPoints(fromspace, tospace RtToken, n RtInt, points RtPointArray) RtPointArray {
	ri.ctx.Handle(List("TransformPoints", []RtPointer{fromspace, tospace, n, points}, nil))
	return RtPointArray{}
}

func (ri *Ri) Translate(dx, dy, dz RtFloat) {
	ri.ctx.Handle(List("Translate", []RtPointer{dx, dy, dz}, nil))
}

func (ri *Ri) TrimCurve(nloops RtInt, ncurves, order RtIntArray, knot, min,
	max RtFloatArray, n RtIntArray, u, v, w RtFloatArray) {

	ri.ctx.Handle(List("TrimCurve", []RtPointer{nloops, ncurves, order, knot, min, max, n, u, v, w}, nil))
}

func (ri *Ri) VPAtmosphere(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("VPAtmosphere", []RtPointer{name}, parameterlist))
}

func (ri *Ri) VPAtmosphereV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.VPAtmosphere(name, ListParams(tokens, values)...)
}

func (ri *Ri) VPInterior(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("VPInterior", []RtPointer{name}, parameterlist))
}

func (ri *Ri) VPInteriorV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.VPInterior(name, ListParams(tokens, values)...)
}

func (ri *Ri) VPSurface(name RtToken, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("VPSurface", []RtPointer{name}, parameterlist))
}

func (ri *Ri) VPSurfaceV(name RtToken, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.VPSurface(name, ListParams(tokens, values)...)
}

func (ri *Ri) Volume(typeof RtToken, bound RtBound, dimensions RtIntArray, parameterlist ...RtPointer) {
	ri.ctx.Handle(List("Volume", []RtPointer{typeof, bound, dimensions}, parameterlist))
}

func (ri *Ri) VolumeV(typeof RtToken, bound RtBound, dimensions RtIntArray, n RtInt, tokens []RtToken, values []RtPointer) {
	ri.Volume(typeof, bound, dimensions, ListParams(tokens, values)...)
}

func (ri *Ri) VolumePixelSamples(x, y RtFloat) {
	ri.ctx.Handle(List("VolumePixelSamples", []RtPointer{x, y}, nil))
}

func (ri *Ri) WorldBegin() {
	ri.ctx.Handle(List("WorldBegin", nil, nil))
}

func (ri *Ri) WorldEnd() {
	ri.ctx.Handle(List("WorldEnd", nil, nil))
}

type RiContext struct {
	*Ri            /* Renderman Interface */
	Utils *Utility /* Utility Interface */
}

type Utility struct {
	ctx RtContextHandle
}

/* grab the serialised (RIB) line from the driver */
func (utils *Utility) GetLastRIB() string {
	return utils.ctx.GetLastRIB()
}

/* RicFlush */
func (utils *Utility) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {

	/* FIXME */
}

/* RicGetProgress */
func (utils *Utility) GetProgress() RtInt {
	return utils.ctx.GetProgress()
}

/* A basic wait for progress function -- wait will every 10ms check the status via GetProgress
 * if a different value from the last check we send that value to the user supplied function (f), 
 * the return from the user function determines if we continue or not. The function always quits
 * on progress 100 regardless. 
 *
 * NOTE: if you render using "multires" in the Display statement, the progress will halt on 99 
 * until the display is closed. In this case this function will block the program control. 
 * 
 */
func (utils *Utility) Wait(f func(RtInt) bool) {

	var last RtInt

	for {
		progress := utils.GetProgress()

		if progress != last && f != nil {
			if f(progress) {
				break
			}
		}

		if progress == RtInt(100) {
			break
		}

		time.Sleep(10 * time.Millisecond)
		last = progress
	}
}

func (utils *Utility) WaitCh() chan RtInt {

	ch := make(chan RtInt,100)
	
	go func() {
		var last RtInt

		for {
			progress := utils.GetProgress()

			if progress != last {
				last = progress
				ch <- progress
			}

			if progress == RtInt(100) {
				break
			}

			time.Sleep(10 * time.Millisecond)
		}
	}()

	return ch
}

/* RIB parsing code */
func (utils *Utility) RIBString(stream string) error {
	return parse(strings.NewReader(stream),utils.ctx)
}

func (utils *Utility) RIB(reader io.Reader) error {
	return parse(reader,utils.ctx)
}


func Wrap(ctx RtContextHandle) *RiContext {
	return &RiContext{&Ri{ctx}, &Utility{ctx}}
}

