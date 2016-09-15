package ri

import (
	"fmt"
	
	"github.com/golang/protobuf/proto"
	. "github.com/mae-global/rigo2/ri/core"
)

/* debug test */
func stream(name RtToken,value []RtPointer) ([]byte,error) {

	values := make([]*PBiValue,0)

	for _,v := range value {

		switch v.Type() {
			case "float":
				r,_ := v.(RtFloat)
				values = append(values,&PBiValue{Type:PBt_FLOAT,Float:float64(r)})
			break
			case "int":
				r,_ := v.(RtInt)
				values = append(values,&PBiValue{Type:PBt_INT,Int:int64(r)})
			break
			default:
				return nil,fmt.Errorf("unsupported type \"%s\"\n",v.Type())
			break
		}

	}
	

	pbi := &PBi{
		Name: string(name),
		Values: values,
	}			
	

	return proto.Marshal(pbi)
}


func unstream(data []byte) (RtToken,[]RtPointer,error) {
	
	var pbi PBi

	if err := proto.Unmarshal(data, &pbi); err != nil {
		return RtToken(""),nil,err
	}

	name := RtToken(pbi.Name)
	values := make([]RtPointer,0)	

	for _,v := range pbi.Values {
		
		switch v.Type {
			case PBt_FLOAT:
				values = append(values,RtFloat(v.Float))
			break
			case PBt_INT:
				values = append(values,RtInt(v.Int))
			break
			default:
				return RtToken(""),nil,fmt.Errorf("unsupported type \"%s\"\n",PBt_name[int32(v.Type)])
			break
		}
	
	}

	return name,values,nil
}




	
