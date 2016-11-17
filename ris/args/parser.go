package args

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"

	. "github.com/mae-global/rigo2/ri/core"
)

var (
	ErrSkipParam    = errors.New("skip param")
	ErrInvalidParam = errors.New("invalid param")
)

/* InfoOutput */
type InfoOutput struct {
	Name  RtString
	Types []RtToken
}

/* InfoParam */
type InfoParam struct {
	Label       RtString
	Name        RtString
	Default     RtPointer /* we can get the type of the param from the default */
	Min         RtPointer
	Max         RtPointer
	Widget      RtToken
	Help        RtString
	Connectable RtBoolean

	/* FIXME: add in hintdict */
}

func (p *InfoParam) String() string {
	return string(p.Name)
}

func (p *InfoParam) Type() string {
	return "param"
}

func (p *InfoParam) Break() []RtPointer {
	return []RtPointer{RtToken(p.Default.Type() + " " + string(p.Name)), p.Default}
}

type InfoPage struct {
	Name   RtString
	Open   RtBoolean
	Params []InfoParam
}

/* Information */
type Information struct {
	ShaderType     RtString
	NodeId         RtToken
	Name           RtString
	Classification RtString
	Help           RtString

	Pages []InfoPage

	Params  []InfoParam
	Outputs []InfoOutput
}

func Str2Normal(str string) RtNormal {

	parts := strings.Split(strings.TrimSpace(str), " ")

	if len(parts) != 3 {
		return RtNormal{0, 0, 0}
	}
	out := RtNormal{0, 0, 0}

	for i, part := range parts {
		if f, err := strconv.ParseFloat(part, 64); err != nil {
			/* eat error */
			continue
		} else {
			out[i] = RtFloat(f)
		}
	}

	return out
}

func Str2Point(str string) RtPoint {

	parts := strings.Split(strings.TrimSpace(str)," ")
	
	if len(parts) != 3 {
		return RtPoint{0,0,0}
	}
	out := RtPoint{0,0,0}

	for i,part := range parts {
		if f,err := strconv.ParseFloat(part,64); err != nil {
			continue
		} else {
			out[i] = RtFloat(f)
		}
	}

	return out
}

func Str2Vector(str string) RtVector {

	parts := strings.Split(strings.TrimSpace(str), " ")

	if len(parts) != 3 {
		return RtVector{0, 0, 0}
	}
	out := RtVector{0, 0, 0}

	for i, part := range parts {
		if f, err := strconv.ParseFloat(part, 64); err != nil {
			/* eat error */
			continue
		} else {
			out[i] = RtFloat(f)
		}
	}

	return out
}

func Str2Color(str string) RtColor {

	parts := strings.Split(strings.TrimSpace(str), " ")
	out := make([]RtFloat, 0)

	for _, part := range parts {
		if f, err := strconv.ParseFloat(part, 64); err != nil {
			/* eat error */
			continue
		} else {
			out = append(out, RtFloat(f))
		}
	}

	return RtColor(out)
}

func ParseArgsParam(param *Param) (InfoParam, error) {
	if param == nil {
		return InfoParam{}, ErrInvalidParam
	}

	var info InfoParam

	var def RtPointer
	var min RtPointer
	var max RtPointer

	switch param.Type {
	case "float":
		def = RtFloat(0.0)
		min = RtFloat(0.0)
		max = RtFloat(0.0)

		param.Default = strings.Replace(param.Default, "f", "", -1)

		if len(param.Default) > 0 {
			if f, err := strconv.ParseFloat(param.Default, 64); err != nil {
				return info, err
			} else {
				def = RtFloat(f)
			}
		}
		if len(param.Min) > 0 {
			if f, err := strconv.ParseFloat(param.Min, 64); err != nil {
				return info, err
			} else {
				min = RtFloat(f)
			}
		}
		if len(param.Max) > 0 {
			if f, err := strconv.ParseFloat(param.Max, 64); err != nil {
				return info, err
			} else {
				max = RtFloat(f)
			}
		}
		break
	case "int":
		def = RtInt(0)
		min = RtInt(0)
		max = RtInt(0)

		if len(param.Default) > 0 {
			if i, err := strconv.Atoi(param.Default); err != nil {
				return info, err
			} else {
				def = RtInt(i)
			}
		}
		if len(param.Min) > 0 {
			if i, err := strconv.Atoi(param.Min); err != nil {
				return info, err
			} else {
				min = RtInt(i)
			}
		}
		if len(param.Max) > 0 {
			if i, err := strconv.Atoi(param.Max); err != nil {
				return info, err
			} else {
				max = RtInt(i)
			}
		}

		break
	case "color":
		def = RtColor{0, 0, 0}
		min = RtColor{0, 0, 0}
		max = RtColor{0, 0, 0}

		if len(param.Default) > 0 {
			def = Str2Color(param.Default)
		}
		if len(param.Min) > 0 {
			min = Str2Color(param.Min)
		}
		if len(param.Max) > 0 {
			max = Str2Color(param.Max)
		}

		break
	case "normal":
		def = RtNormal{0, 0, 0}
		min = RtNormal{0, 0, 0}
		max = RtNormal{0, 0, 0}

		if len(param.Default) > 0 {
			def = Str2Normal(param.Default)
		}
		if len(param.Min) > 0 {
			min = Str2Normal(param.Min)
		}
		if len(param.Max) > 0 {
			max = Str2Normal(param.Max)
		}
		break
	case "vector":
		def = RtVector{0, 0, 0}
		min = RtVector{0, 0, 0}
		max = RtVector{0, 0, 0}

		if len(param.Default) > 0 {
			def = Str2Vector(param.Default)
		}
		if len(param.Min) > 0 {
			min = Str2Vector(param.Min)
		}
		if len(param.Max) > 0 {
			max = Str2Vector(param.Max)
		}
		break
	case "point":
		def = RtPoint{0,0,0}
		min = RtPoint{0,0,0}
		max = RtPoint{0,0,0}

		if len(param.Default) > 0 {
			def = Str2Point(param.Default)
		}
		if len(param.Min) > 0 {
			min = Str2Point(param.Min)
		}
		if len(param.Max) > 0 {
			max = Str2Point(param.Max)
		}
		break
	case "string":
		def = RtString(param.Default)
		min = RtString(param.Min)
		max = RtString(param.Max)
		break
	case "struct": /* FIXME, don't know how to handle this type ? */
		return info, ErrSkipParam
		break
	default:
		return info, fmt.Errorf("Unknown Type %s=[%s]", param.Name, param.Type)
		break
	}

	info.Label = RtString(param.Label)
	info.Name = RtString(param.Name)
	info.Widget = RtToken(param.Widget)
	info.Help = RtString(strings.TrimSpace(param.Help.Value))
	info.Default = def
	info.Min = min
	info.Max = max
	info.Connectable = RtBoolean(false)
	if param.Connectable == "True" {
		info.Connectable = RtBoolean(true)
	}

	return info, nil
}

func ParseArgsXML(data []byte) (*Args, error) {

	var a Args
	if err := xml.Unmarshal(data, &a); err != nil {
		return nil, err
	}
	return &a, nil
}

func Parse(name string, data []byte) (*Information, error) {

	args, err := ParseArgsXML(data)
	if err != nil {
		return nil, err
	}

	info := new(Information)
	info.ShaderType = RtString(args.Shader.Tag.Value)
	info.NodeId = RtToken(args.Rfmdata.NodeId)
	info.Name = RtString(name)
	info.Classification = RtString(args.Rfmdata.Classification)
	info.Help = RtString(strings.TrimSpace(args.Help.Value))

	/* start with the pages */
	info.Pages = make([]InfoPage, 0)

	for _, page := range args.Pages {
		p := InfoPage{}
		p.Name = RtString(page.Name)
		p.Open = RtBoolean(true)
		if page.Open == "False" {
			p.Open = RtBoolean(false)
		}
		p.Params = make([]InfoParam, 0)

		for _, param := range page.Params {

			infoparam, err := ParseArgsParam(&param)
			if err != nil {
				if err == ErrSkipParam {
					continue
				}
				return nil, err
			}

			p.Params = append(p.Params, infoparam)
		}

		info.Pages = append(info.Pages, p)
	}

	info.Params = make([]InfoParam, 0)

	for _, param := range args.Params {
		infoparam, err := ParseArgsParam(&param)
		if err != nil {
			if err == ErrSkipParam {
				continue
			}
			return nil, err
		}

		info.Params = append(info.Params, infoparam)
	}

	info.Outputs = make([]InfoOutput, 0)

	for _, output := range args.Outputs {
		op := InfoOutput{}
		op.Name = RtString(output.Name)
		op.Types = make([]RtToken, 0)
		for _, t := range output.Tags.Tags {
			op.Types = append(op.Types, RtToken(t.Value))
		}

		info.Outputs = append(info.Outputs, op)
	}

	return info, nil
}
