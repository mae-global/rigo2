/* rigo/core/definitions.go */
package core

import (
	"fmt"
)

type RtContextHandle interface {
	HandleError(*RtError) *RtError
	Handle([]RtPointer)
	HandleV(RtString,[]RtPointer,[]RtPointer,[]RtPointer)
	GenHandle(string, string) (string, error)
	Set(string, string) RtToken /* Dictionary, entries set with RiDeclare() */
	GetProgress() RtInt         /* Get the current progress (as from a pipe output when rendering) */
	GetLastRIB() string
}

type RtError struct {
	Code     int
	Severity int
	Msg      string
}

func ToType(t RtPointer) string {

	foo := ""
	for _,c := range t.Type() { 
	
		if c == '[' {
			foo += "_array"
			break
		}
		foo += string(c)
	}
	return foo
}


func (err *RtError) String() string {
	return fmt.Sprintf("%05d, %d -- %s", err.Code, err.Severity, err.Msg)
}

func (err *RtError) Error() error {
	return fmt.Errorf("%05d, %d -- %s", err.Code, err.Severity, err.Msg)
}

func Error(code, severity int, msg string) *RtError {
	return &RtError{code, severity, msg}
}

func Errorf(code, severity int, format string, params ...interface{}) *RtError {
	return &RtError{code, severity, fmt.Sprintf(format, params...)}
}

/* subvert RtPointer into a more Go useful interface */
type RtPointer interface {
	String() string
	Type() string
}

/* used to push lists of pointers */
type RtArray []RtPointer

func (s RtArray) String() string {
	return ""
}

func (s RtArray) Type() string { return "array" }

type RtAttributer interface {
	Break() []RtPointer
}

type RtString string

func (s RtString) String() string {
	return fmt.Sprintf("\"%s\"", string(s))
}
func (s RtString) Type() string { return "string" }

type RtStringArray []RtString

func (s RtStringArray) String() string {
	out := "["
	for i, r := range s {
		if i > 0 {
			out += " "
		}
		out += r.String()
	}
	return out + "]"
}
func (s RtStringArray) Type() string { return fmt.Sprintf("string[%d]", len(s)) }

/* RtBoolean */
type RtBoolean bool

func (s RtBoolean) String() string {
	if bool(s) == false {
		return "0"
	}
	return "1"
}
func (s RtBoolean) Type() string { return "boolean" }

/* RtToken -- token */
type RtToken string

func (s RtToken) String() string {
	return fmt.Sprintf("\"%s\"", string(s))
}
func (s RtToken) Type() string { return "token" }

/* RtTokenArray -- token[] */
type RtTokenArray []RtToken

func (s RtTokenArray) String() string {
	out := "["
	for i, r := range s {
		if i > 0 {
			out += " "
		}
		out += r.String()
	}
	return out + "]"
}
func (s RtTokenArray) Type() string { return fmt.Sprintf("token[%d]", len(s)) }

/* RtInt -- int */
type RtInt int64

func (s RtInt) String() string {
	return fmt.Sprintf("%d", int64(s))
}
func (s RtInt) Type() string { return "int" }

/* RtIntArray -- int[] */
type RtIntArray []RtInt

func (s RtIntArray) String() string {
	out := "["
	for i, r := range s {
		if i > 0 {
			out += " "
		}
		out += r.String()
	}
	return out + "]"
}
func (s RtIntArray) Type() string { return fmt.Sprintf("int[%d]", len(s)) }

/* RtFloat -- float */
type RtFloat float64

func (s RtFloat) String() string {
	return Reduce(s)
}
func (s RtFloat) Type() string { return "float" }

/* RtFloatArray -- float[] */
type RtFloatArray []RtFloat

func (s RtFloatArray) String() string {
	return "[" + Reducev(s) + "]"
}
func (s RtFloatArray) Type() string { return fmt.Sprintf("float[%d]", len(s)) }

func (s RtFloatArray) ToPoint() RtPoint {

	out := [3]RtFloat{0,0,0}
	if len(s) == 0 {
		return out
	}	

	sc := s[0]
	
	for i := 0; i < 3; i++ {
		out[i] = sc
		if i < len(s) {  
			sc = s[i]		
		}
	}
	return out
}

func (s RtFloatArray) ToNormal() RtNormal {

	out := [3]RtFloat{0,0,0}
	if len(s) == 0 {
		return out
	}

	sc := s[0]
	for i := 0; i < 3; i++ {
		out[i] = sc
		if i < len(s) {
			sc = s[i]
		}
	}
	return out
}


func (s RtFloatArray) ToHpoint() RtHpoint {

	out := [4]RtFloat{0,0,0,0}
	if len(s) == 0 {
		return out
	}
	
	sc := s[0]
	for i := 0; i < 4; i++ {
		out[i] = sc
		if i < len(s) {
			sc = s[i]
		}
	}
	return out
}

func (s RtFloatArray) ToMatrix() RtMatrix {

	out := [16]RtFloat{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	if len(s) == 0 {
		return out
	}

	sc := s[0]
	for i := 0; i < 16; i++ {
		out[i] = sc
		if i < len(s) {
			sc = s[i]
		}
	}
	return out
}

func (s RtFloatArray) ToBasis() RtBasis {

	out := [16]RtFloat{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	if len(s) == 0 {
		return out
	}

	sc := s[0]
	for i := 0; i < 16; i++ {
		out[i] = sc
		if i < len(s) {
			sc = s[i]
		}
	}
	return out
}

func (s RtFloatArray) ToBound() RtBound {

	out := [6]RtFloat{0,0,0,0,0,0}
	if len(s) == 0 {
		return out
	}

	sc := s[0]
	for i := 0; i < 6; i++ {
		out[i] = sc
		if i < len(s) {
			sc = s[i]
		}
	}
	return out
}


/* RtColor -- color */
type RtColor []RtFloat /* 3 components is normal */
func (s RtColor) String() string {
	return "[" + Reducev(s) + "]"
}
func (s RtColor) Type() string { return "color" }

/* RtPoint */
type RtPoint [3]RtFloat

func (s RtPoint) String() string {
	return Reduce(s[0]) + " " + Reduce(s[1]) + " " + Reduce(2)
}
func (s RtPoint) Type() string { return "point" }


/* RtPointArray */
type RtPointArray []RtPoint

func (s RtPointArray) String() string {
	out := ""
	for i, p := range s {
		if i > 0 && i < len(s)-1 {
			out += " "
		}
		out += p.String()
	}
	return out
}
func (s RtPointArray) Type() string { return fmt.Sprintf("point[%d]", len(s)) }

/* RtInterval */
type RtInterval [2]RtFloat

func (s RtInterval) String() string {
	return Reduce(s[0]) + " " + Reduce(s[1])
}
func (s RtInterval) Type() string { return "interval" }

/* RtVector */
type RtVector [3]RtFloat

func (s RtVector) String() string {
	return Reduce(s[0]) + " " + Reduce(s[1]) + " " + Reduce(s[2])
}
func (s RtVector) Type() string { return "vector" }

/* RtNormal */
type RtNormal [3]RtFloat

func (s RtNormal) String() string {
	return Reduce(s[0]) + " " + Reduce(s[1]) + " " + Reduce(s[2])
}
func (s RtNormal) Type() string { return "type" }

/* RtHpoint */
type RtHpoint [4]RtFloat

func (s RtHpoint) String() string {
	return Reduce(s[0]) + " " + Reduce(s[1]) + " " + Reduce(s[2]) + " " + Reduce(s[3])
}
func (s RtHpoint) Type() string { return "hpoint" }

/* RtMatrix */
type RtMatrix [16]RtFloat

func (s RtMatrix) String() string {
	out := ""
	for i := 0; i < 16; i++ {
		out += Reduce(s[i])
		if i < 15 {
			out += " "
		}
	}
	return "[" + out + "]"
}
func (s RtMatrix) Type() string { return "matrix" }

/* RtBasis */
type RtBasis [16]RtFloat

func (s RtBasis) String() string {
	out := ""
	for i := 0; i < 16; i++ {
		out += Reduce(s[i])
		if i < 15 {
			out += " "
		}
	}
	return "[" + out + "]"
}
func (s RtBasis) Type() string { return "basis" }

/* RtBound */
type RtBound [6]RtFloat

func (s RtBound) String() string {
	out := ""
	for i := 0; i < 6; i++ {
		out += Reduce(s[i])
		if i < 5 {
			out += " "
		}
	}
	return "[" + out + "]"
}
func (s RtBound) Type() string { return "bound" }

/* RtLightHandle */
type RtLightHandle string

func (s RtLightHandle) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtLightHandle) Type() string   { return "lighthandle" }

/* RtObjectHandle */
type RtObjectHandle string

func (s RtObjectHandle) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtObjectHandle) Type() string   { return "objecthandle" }

/* RtShaderHandle */
type RtShaderHandle string

func (s RtShaderHandle) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtShaderHandle) Type() string   { return "shaderhandle" }

/* RtArchiveHandle */
type RtArchiveHandle string

func (s RtArchiveHandle) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtArchiveHandle) Type() string   { return "archivehandle" }

/* RtArchiveCallback */
type RtArchiveCallback string

func (s RtArchiveCallback) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtArchiveCallback) Type() string   { return "archivecallback" }

/* RtFilterFunc */
type RtFilterFunc string

func (s RtFilterFunc) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtFilterFunc) Type() string   { return "filterfunc" }

/* RtErrorHandler */
type RtErrorHandler string

func (s RtErrorHandler) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtErrorHandler) Type() string   { return "errorhandler" }

/* RtProcSubdivFunc subdivision function */
type RtProcSubdivFunc string

func (s RtProcSubdivFunc) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtProcSubdivFunc) Type() string   { return "procsubdivfunc" }

/* RtProc2SubdivFunc */
type RtProc2SubdivFunc string

func (s RtProc2SubdivFunc) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtProc2SubdivFunc) Type() string   { return "procs2ubdivfunc" }

/* RtProc2BoundFunc */
type RtProc2BoundFunc string

func (s RtProc2BoundFunc) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtProc2BoundFunc) Type() string   { return "proc2boundfunc" }

/* RtProcFreeFunc */
type RtProcFreeFunc string

func (s RtProcFreeFunc) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtProcFreeFunc) Type() string   { return "procfreefunc" }

/* RtArchiveCallbackFunc */
type RtArchiveCallbackFunc string

func (s RtArchiveCallbackFunc) String() string { return fmt.Sprintf("\"%s\"", string(s)) }
func (s RtArchiveCallbackFunc) Type() string   { return "archivecallbackfunc" }
