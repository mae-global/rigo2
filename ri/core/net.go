package core

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)


/* Encode -- encode a command sequence to proto-buffers */
func Encode(name RtString,list []RtPointer) ([]byte, error) {

	p := &PBi{ Name: string(name), Values: make([]*PBiValue,0) }

	for _,value := range list {
	
		var v *PBiValue

		switch ToType(value) {
			case "string":
				av,ok := value.(RtString)
				if !ok {
					return nil,fmt.Errorf("invalid string")
				}

				v = &PBiValue{Type:PBt_STRING,String_:string(av)}			

			break
			case "token":
				av,ok := value.(RtToken)
				if !ok {
					return nil,fmt.Errorf("invalid token")
				}

				v = &PBiValue{Type:PBt_TOKEN,String_:string(av)}
			break
			case "int":
				av,ok := value.(RtInt)
				if !ok {
					return nil,fmt.Errorf("invalid integer")
				}
				v = &PBiValue{Type:PBt_INT,Int:int64(av)}
			break
			case "float":
				av,ok := value.(RtFloat)
				if !ok {
					return nil,fmt.Errorf("invalid float")
				}
				v = &PBiValue{Type:PBt_FLOAT,Float:float64(av)}
			break
			case "string_array":
				av,ok := value.(RtStringArray)
				if !ok {
					return nil,fmt.Errorf("invalid string array")
				}
				nav := make([]string,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = string(av[i])
				}

				v = &PBiValue{Type:PBt_STRINGARRAY,StringArray:nav}
			break
			case "token_array":
				av,ok := value.(RtTokenArray)
				if !ok {
					return nil,fmt.Errorf("invalid token array")
				}
				nav := make([]string,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = string(av[i])
				}
			
				v = &PBiValue{Type:PBt_TOKENARRAY,StringArray:nav}
			break
			case "int_array":
				av,ok := value.(RtIntArray)
				if !ok {
					return nil,fmt.Errorf("invalid int array")
				}
				nav := make([]int64,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = int64(av[i])
				}

				v = &PBiValue{Type:PBt_INTARRAY,IntArray:nav}
			break
			case "float_array":
				av,ok := value.(RtFloatArray)
				if !ok {
					return nil,fmt.Errorf("invalid float array")
				}
				nav := make([]float64,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = float64(av[i])
				}
					
				v = &PBiValue{Type:PBt_FLOATARRAY,FloatArray:nav}
			break
			case "boolean":
				av,ok := value.(RtBoolean)
				if !ok {
					return nil,fmt.Errorf("invalid boolean")
				}
				v = &PBiValue{Type:PBt_BOOLEAN,Boolean:bool(av)}
			break
			case "color":
				av,ok := value.(RtColor)
				if !ok {
					return nil,fmt.Errorf("invalid color")
				}
				nav := make([]float64,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = float64(av[i])
				}

				v = &PBiValue{Type:PBt_COLOR,FloatArray:nav}
			break
			case "point":
				av,ok := value.(RtPoint)
				if !ok {
					return nil,fmt.Errorf("invalid point")
				}
				nav := make([]float64,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = float64(av[i])
				}

				v = &PBiValue{Type:PBt_POINT,FloatArray:nav}
			break
			/*
			case "color_array":
				av,ok := value.(RtColorArray)
				if !ok {
					return nil,fmt.Errorf("invalid color array")
				}
			break
			*/
			case "normal":
				av,ok := value.(RtNormal)
				if !ok {
					return nil,fmt.Errorf("invalid normal")
				}
				nav := make([]float64,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = float64(av[i])
				}

				v = &PBiValue{Type:PBt_NORMAL,FloatArray:nav}
			break
			case "hpoint":
				av,ok := value.(RtHpoint)				
				if !ok {
					return nil,fmt.Errorf("invalid hpoint")
				}
				nav := make([]float64,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = float64(av[i])
				}

				v = &PBiValue{Type:PBt_HPOINT,FloatArray:nav}
			break
			case "matrix":
				av,ok := value.(RtMatrix)
				if !ok {
					return nil,fmt.Errorf("invalid matrix")
				}
				nav := make([]float64,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = float64(av[i])
				}
	
				v = &PBiValue{Type:PBt_MATRIX,FloatArray:nav}
			break
			case "basis":
				av,ok := value.(RtBasis)
				if !ok {
					return nil,fmt.Errorf("invalid basis")
				}
				nav := make([]float64,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = float64(av[i])
				}
	
				v = &PBiValue{Type:PBt_BASIS,FloatArray:nav}
			break
			case "bound":
				av,ok := value.(RtBound)
				if !ok {
					return nil,fmt.Errorf("invalid bound")
				}
				nav := make([]float64,len(av))
				for i := 0; i < len(av); i++ {
					nav[i] = float64(av[i])
				}

				v = &PBiValue{Type:PBt_BOUND,FloatArray:nav}
			break
			case "lighthandle":
				av,ok := value.(RtLightHandle)
				if !ok {
					return nil,fmt.Errorf("invalid light handle")
				}
				v = &PBiValue{Type:PBt_LIGHTHANDLE,String_:string(av)}
			break
			case "objecthandle":
				av,ok := value.(RtObjectHandle)
				if !ok {
					return nil,fmt.Errorf("invalid object handle")
				}
				v = &PBiValue{Type:PBt_OBJECTHANDLE,String_:string(av)}
			break
			case "shaderhandler":
				av,ok := value.(RtShaderHandle)
				if !ok {
					return nil,fmt.Errorf("invalid shader handle")
				}
				v = &PBiValue{Type:PBt_SHADERHANDLE,String_:string(av)}
			break							
	/*
	PBt_POINTARRAY          PBt = 12
	PBt_INTERVAL            PBt = 13
	PBt_NORMAL              PBt = 14
	PBt_HPOINT              PBt = 15
	PBt_MATRIX              PBt = 16
	PBt_BASIS               PBt = 17
	PBt_BOUND               PBt = 18
	PBt_LIGHTHANDLE         PBt = 19
	PBt_OBJECTHANDLE        PBt = 20
	PBt_SHADERHANDLE        PBt = 21
	PBt_ARCHIVEHANDLE       PBt = 22
	PBt_FILTERFUNC          PBt = 23
	PBt_ERRORHANDLE         PBt = 24
	PBt_PROCSUBDIVFUNC      PBt = 25
	PBt_PROC2SUBDIVFUNC     PBt = 26
	PBt_PROC2BOUNDFUNC      PBt = 27
	PBt_PROCFREEFUNC        PBt = 28
	PBt_ARCHIVECALLBACKFUNC PBt = 29
*/
			default:
				return nil,fmt.Errorf("not implemented %s",value.Type())
			break
		}

		p.Values = append(p.Values,v)
	}

	return proto.Marshal(p)
}

/* Decode -- decode a command sequence from proto-buffers */
func Decode(data []byte) (RtString,[]RtPointer,error) {

	array := func(v *PBiValue) []RtFloat {
		if len(v.FloatArray) == 0 {
			return nil
		}
		out := make([]RtFloat,len(v.FloatArray))
		for i := 0; i < len(v.FloatArray); i++ {
			out[i] = RtFloat(v.FloatArray[i])
		}
		return out
	}
		

	out := &PBi{}
	if err := proto.Unmarshal(data,out); err != nil {
		return "",nil,err
	}

	list := make([]RtPointer,0)

	for _,value := range out.Values {
		switch value.Type {
			case PBt_STRING:
				list = append(list,RtString(value.String_))
			break
			case PBt_TOKEN:
				list = append(list,RtToken(value.String_))
			break
			case PBt_INT:
				list = append(list,RtInt(value.Int))
			break
			case PBt_FLOAT:
				list = append(list,RtFloat(value.Float))
			break
			case PBt_STRINGARRAY:
				nav := make([]RtString,len(value.StringArray))
				for i := 0; i < len(value.StringArray); i++ {
					nav[i] = RtString(value.StringArray[i])
				}
				list = append(list,RtStringArray(nav))
			break
			case PBt_TOKENARRAY:
				nav := make([]RtToken,len(value.StringArray))
				for i := 0; i < len(value.StringArray); i++ {
					nav[i] = RtToken(value.StringArray[i])
				}
				list = append(list,RtTokenArray(nav))
			break
			case PBt_INTARRAY:
				nav := make([]RtInt,len(value.IntArray))
				for i := 0; i < len(value.IntArray); i++ {
					nav[i] = RtInt(value.IntArray[i])
				}
				list = append(list,RtIntArray(nav))
			break
			case PBt_FLOATARRAY:
				list = append(list,RtFloatArray(array(value)))
			break
			case PBt_BOOLEAN:
				list = append(list,RtBoolean(value.Boolean))
			break
			case PBt_COLOR:
				list = append(list,RtColor(array(value)))
			break
			case PBt_POINT:
				list = append(list,RtFloatArray(array(value)).ToPoint())
			break
			/*		
			case PBt_COLORARRAY:

			break
			case PBt_POINTARRAY:

			break
			case PBt_INTERVAL:

			break
			*/
			case PBt_NORMAL:
				list = append(list,RtFloatArray(array(value)).ToNormal())
			break
			case PBt_HPOINT:
				list = append(list,RtFloatArray(array(value)).ToHpoint())

			break
			case PBt_MATRIX:
				list = append(list,RtFloatArray(array(value)).ToMatrix())
			break
			case PBt_BASIS:
				list = append(list,RtFloatArray(array(value)).ToBasis())
			break
			case PBt_BOUND:
				list = append(list,RtFloatArray(array(value)).ToBound())
			break
			case PBt_LIGHTHANDLE:
				list = append(list,RtLightHandle(value.String_))
			break
			case PBt_OBJECTHANDLE:
				list = append(list,RtObjectHandle(value.String_))
			break
			case PBt_SHADERHANDLE:
				list = append(list,RtShaderHandle(value.String_))
			break

			default:
				return "",nil,fmt.Errorf("not implemented %s",value.Type)
			break
		}
	}


	return RtString(out.Name),list,nil
}
