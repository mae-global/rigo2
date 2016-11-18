package parser

import (
	"encoding/json"
)



func Parse(name string,data []byte) (*RenderManAsset,error) {

	var outer *Outer
	
	if err := json.Unmarshal(data,&outer); err != nil {
		return nil,err
	}

	return outer.RenderManAsset,nil
}




