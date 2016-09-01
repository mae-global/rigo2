package prototype

import (
	"fmt"
	"strings"

	. "github.com/mae-global/rigo2/ri/core"
)

const (
	PARAMETERLIST = RtString("__parameterlist__")
	BesselFilter RtFilterFunc = "bessel"
	ProcDelayedReadArchive RtProcSubdivFunc = "DelayedReadArchive"
	ProcFree RtProcFreeFunc = "free"	
	Proc2DelayedReadArchive RtProc2SubdivFunc = "DelayedReadArchive"
)

type Argument struct {
	Type string
	Example RtPointer
	Name string
}

func (arg *Argument) String() string {
	return fmt.Sprintf("name=\"%s\", type=\"%s\", example=%s",arg.Name,arg.Type,arg.Example)
}

type Information struct {
	Name RtString
	Arguments []*Argument
	Parameterlist bool
}

func (info *Information) String() string {
	out := fmt.Sprintf("Information \"%s\", parameterlist=%v\n",info.Name,info.Parameterlist)
	if len(info.Arguments) > 0 {
		out += fmt.Sprintf("\t%d arguments\n",len(info.Arguments))
		for i,arg := range info.Arguments {
			out += fmt.Sprintf("\t\t[%03d] \"%s\" -- type=\"%s\", example=%s\n",i,arg.Name,arg.Type,arg.Example)
		}
	}
	return out
}


func Parse(stream string) *Information {

	list := strings.Split(stream," ")
	if len(list) == 0 {
		return &Information{Name:RtString(stream)}
	}

	proto := &Information{}
	
	proto.Name = RtString(list[0])
	proto.Arguments = make([]*Argument,0)	

	var arg *Argument

	/* example :- "Shader token name token handle ..." */	
	var r RtPointer
	flipflop := true

	for i := 1; i < len(list); i++ {
		r = nil

		switch list[i] {
			case "string":
				r = RtString("string")
			break
			case "string[]":
				r = RtStringArray{}
			break
			case "float":
				r = RtFloat(1)
			break
			case "float[]":
				r = RtFloatArray{1,2,3}
			break
			case "int":
				r = RtInt(1)
			break
			case "int[]":
				r = RtIntArray{1,2,3}
			break
			case "token":
				r = RtToken("name")
			break
			case "tokenarray":
				r = RtTokenArray{}
			break
			case "lighthandle":
				r = RtLightHandle("light")
			break
			case "objecthandle":
				r = RtObjectHandle("object")
			break
			case "filterfunc":
				r = BesselFilter
			break
			case "boolean":
				r = RtBoolean(true)
			break
			case "color":
				r = RtColor{1,1,1}
			break
			case "point":
				r = RtPoint{1,1,1}
			break
			case "point[]":
				r = RtPointArray{}
			break
			case "basis":
				r = RtBasis{}
			break
			case "bound":
				r = RtBound{}
			break
			case "matrix":
				r = RtMatrix{}
			break
			case "pointer":
				r = RtStringArray{}
			break
			case "procsubdivfunc":
				r = ProcDelayedReadArchive
			break
			case "procfreefunc":
				r = ProcFree
			break
			case "proc2subdivfunc":
				r = Proc2DelayedReadArchive
			break
			case "...":
				r = PARAMETERLIST
			break
		}

		if !flipflop {

			/* then it is a name */
			arg.Name = list[i]
		} else {

			if arg != nil && len(arg.Type) > 0 {
				proto.Arguments = append(proto.Arguments,arg)
				arg = new(Argument)
			}

			if r == PARAMETERLIST {
				
				proto.Parameterlist = true
				
			} else {

				arg = new(Argument)
				arg.Type = list[i]
				arg.Example = r
			}
		}

		flipflop = !flipflop
	}
	if arg != nil && len(arg.Type) > 0 {
		proto.Arguments = append(proto.Arguments,arg)
	}	

	
	
	return proto
}


