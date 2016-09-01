package ri

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/mae-global/rigo2/ris/args"
)

const (
	ARGUMENTS     = RtToken("__rigo2_arguments")
	PARAMETERLIST = RtToken("__rigo2_parameterlist")
	ANNOTATIONS   = RtToken("__rigo2_annotations")
	Null          = RtToken("")
)

/* -- /ri/* -- */

/* create a list */
func List(name RtString, args, params []RtPointer) []RtPointer {
	list := []RtPointer{name, ARGUMENTS}

	if args != nil {
		list = append(list, args...)
	}

	list = append(list, PARAMETERLIST)

	if params != nil {
		for _, param := range params {
			if param.Type() == "param" {
				if attr, ok := param.(*InfoParam); ok {
					list = append(list, attr.Break()...)
				}
				continue
			}
			/* check for array type then decant to stream */
			if param.Type() == "array" {
				if array, ok := param.(RtArray); ok {
					list = append(list, array...)
				}
			} else {
				list = append(list, param)
			}
		}
	}
	return list
}

func AnnotatedList(name RtString,args,params,annotations []RtPointer) []RtPointer {
	list := List(name,args,params)

	if annotations != nil {
		list = append(list,ANNOTATIONS)
		list = append(list,annotations...)
	}
	return list
}
		

func ListParams(tokens RtTokenArray,values []RtPointer) []RtPointer {
	params := make([]RtPointer, 0)

	for i, token := range tokens {
		params = append(params, token)
		if i >= len(values) {
			break
		}
		params = append(params, values[i])
	}
	return params
}






type SpecificationInfo struct {
	Class string
	Type  string
	Name  string
	Size  int
}

func (info SpecificationInfo) IsInline() bool {
	if info.Type == "" {
		return false
	}
	return true
}

func (info SpecificationInfo) LongType() string {
	if info.Size == 1 {
		if info.Class == "reference" {
			return "string"
		} else {
			return info.Type
		}
	}

	return fmt.Sprintf("%s[%d]", info.Type, info.Size)
}

func (info SpecificationInfo) Declaration() string {
	out := ""
	if len(info.Class) > 0 {
		out = info.Class + " "
	}
	if len(info.Type) > 0 {
		out += info.Type
	}

	if info.Size > 1 {
		out += fmt.Sprintf("[%d]", info.Size)
	}
	return out
}

func (info SpecificationInfo) String() string {

	out := ""
	if len(info.Class) > 0 {
		out = info.Class + " "
	}

	if len(info.Type) > 0 {
		out += info.Type
	}

	if info.Size > 1 {
		out += fmt.Sprintf("[%d]", info.Size)
	}

	if len(info.Name) > 0 {
		if len(out) > 0 {
			out += " "
		}
		out += info.Name
	}
	return out
}

func (info SpecificationInfo) String2() string {
	out := ""
	if len(info.Class) > 0 {
		out = info.Class + " "
	} else {
		out = "unknown-class "
	}

	if len(info.Type) > 0 {
		out += info.Type
	} else {
		out += "unknown-type"
	}

	if info.Size > 1 {
		out += fmt.Sprintf("[%d]", info.Size)
	}

	if len(info.Name) > 0 {
		out += " " + info.Name
	} else {
		out += " unknown-name"
	}
	return out
}

func (info SpecificationInfo) ReplaceName(name string) string {
	out := ""
	if len(info.Class) > 0 {
		out = info.Class + " "
	}

	if len(info.Type) > 0 {
		out += info.Type
	}

	if info.Size > 1 {
		out += fmt.Sprintf("[%d]", info.Size)
	}

	if len(name) > 0 {
		if len(out) > 0 {
			out += " "
		}
		out += name
	}
	return out
}

func Specification(token string) SpecificationInfo {
	/* TODO: this is _very_ messy code, make a small DSL
	 * parser for this.
	 */
	class := ""
	typeof := ""
	size := 1

	parts := strings.Split(strings.TrimSpace(token), " ")

	if len(parts) == 1 {
		return SpecificationInfo{Class: class, Type: typeof, Name: token, Size: size}
	}

	if len(parts) == 2 {
		typeof = parts[0]
		/* TODO: check valid type */

		if strings.Contains(parts[0], "[") && strings.Contains(parts[0], "]") {
			typeparts := strings.Split(parts[0], "[")
			out := strings.Replace(typeparts[1], "]", "", -1)
			if c, err := strconv.Atoi(out); err == nil {
				size = c
			}
			typeof = typeparts[0]
		}
		return SpecificationInfo{Class: class, Type: typeof, Name: parts[1], Size: size}
	}

	class = parts[0]
	typeof = parts[1]

	if strings.Contains(parts[1], "[") && strings.Contains(parts[1], "]") {
		typeparts := strings.Split(parts[1], "[")
		out := strings.Replace(typeparts[1], "]", "", -1)
		if c, err := strconv.Atoi(out); err == nil {
			size = c
		}
		typeof = typeparts[0]
	}
	return SpecificationInfo{Class: class, Type: typeof, Name: parts[2], Size: size}
}

func usedefault(d, a string) string {
	if len(a) > 0 {
		return a
	}
	return d
}

func Merge(base, local SpecificationInfo) SpecificationInfo {

	class := usedefault("uniform", base.Class)
	typeof := usedefault("string", base.Type)
	size := base.Size

	if class != local.Class && len(local.Class) > 0 {
		class = local.Class
	}

	if typeof != local.Type && len(local.Type) > 0 {
		typeof = local.Type
	}

	if size != local.Size && local.Size > 0 {
		size = local.Size
	}

	return SpecificationInfo{Class: class, Type: typeof, Name: base.Name, Size: size}
}

func RIBStream(name RtString, args, tokens, values []RtPointer) string {

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

func ParseBegin(statement RtToken) ([]RtPointer, error) {
	/* examples :-
	 * RI_NULL : in-which case out.rib should be used
	 * [filename].rib : use the default fileout SEQ
	 * stdout : use standard out for RIB output
	 * prman : pipe to renderer -- launches the renderer first */

	out := []RtPointer{}

	parts := strings.Split(string(statement), " ")

	switch parts[0] {
	case "", "-":
		out = []RtPointer{RtToken("out.rib")}
		break
	/*
	case "stdout":
		out = []RtPointer{RtToken("|"), RtToken("stdout")}
		if len(parts) > 1 {
			for _, str := range parts[1:] {
				out = append(out, RtToken(str))
			}
		}
		break
	case "render", "prman":
		out = []RtPointer{RtToken("|"), RtToken("render")}
		if len(parts) > 1 {
			for _, str := range parts[1:] {
				out = append(out, RtToken(str))
			}
		}
		break
	case "catrib", "cat":
		out = []RtPointer{RtToken("|"), RtToken("catrib")}
		if len(parts) > 1 {
			for _, str := range parts[1:] {
				out = append(out, RtToken(str))
			}
		}		
		break
	case "debug":
		out = []RtPointer{RtToken("|"), RtToken("debug")}
		if len(parts) > 1 {
			for _, str := range parts[1:] {
				out = append(out, RtToken(str))
			}
		}
		break
	case "block":
		out = []RtPointer{RtToken("|"),RtToken("block")}
		if len(parts) > 1 {
			for _,str := range parts[1:] {
				out = append(out,RtToken(str))
			}
		}
		break
	*/
	default:
		if strings.HasSuffix(parts[0],".rib") || strings.HasSuffix(parts[0],".rib.gz") {
			out = []RtPointer{statement}
		} else {

			out = []RtPointer{RtToken("|"),RtToken(parts[0])}
			if len(parts) > 1 {
				for _,str := range parts[1:] {
					out = append(out,RtToken(str))
				}
			}
		}
		break
	
	}

	return out, nil
}
