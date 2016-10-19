package core

import (
	"fmt"
	
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)



func RtShouldEqual(actual interface{}, expected ...interface{}) string {
	a,ok := actual.(RtPointer)
	if !ok {
		return "invalid RtPointer"
	}
	b,ok := expected[0].(RtPointer)
	if !ok {
		return "invalid RtPointer"
	}
	if a.String() == b.String() {
		return ""
	}
	
	return fmt.Sprintf("Expected \"%s\", got \"%s\"",b.String(),a.String())
}


func Test_Net(t *testing.T) {

	Convey("Test Net using Protobuffers -- simple",t,func() {

		data,err := Encode(RtString("Foo"),[]RtPointer{RtString("bar")})
		So(err,ShouldBeNil)
		So(data,ShouldNotBeNil)

		fmt.Printf("length of encoded data is %dbytes\n",len(data))

		name,list,err := Decode(data)
		So(err,ShouldBeNil)
		So(name,ShouldEqual,RtString("Foo"))
		So(len(list),ShouldEqual,1)
		So(list[0],ShouldEqual,RtString("bar"))
	})

	Convey("Test Net using protobuffers -- actual command",t,func() {

		data,err := Encode(RtString("Projection"),[]RtPointer{RtToken("perspective"),RtToken("float fov"),RtFloat(30.0)})
		So(err,ShouldBeNil)
		So(data,ShouldNotBeNil)

		fmt.Printf("length of encoded data is %dbytes\n",len(data))
		
		name,list,err := Decode(data)
		So(err,ShouldBeNil)
		So(name,ShouldEqual,RtString("Projection"))
		So(len(list),ShouldEqual,3)	
	
		So(list[0],RtShouldEqual,RtToken("perspective"))
		So(list[1],RtShouldEqual,RtToken("float fov"))
		So(list[2],RtShouldEqual,RtFloat(30.0))	
	})

	Convey("Test Net using protobuffers -- complex command",t,func() {

		data,err := Encode(RtString("Attribute"),[]RtPointer{RtToken("float[2] angle"),RtFloatArray{1.23,32.1},
																												 RtToken("float tint"),RtFloat(1.0)})

		So(err,ShouldBeNil)
		So(data,ShouldNotBeNil)

		fmt.Printf("length of encoded data is %dbytes\n",len(data))

		name,list,err := Decode(data)
		So(err,ShouldBeNil)
		So(name,ShouldEqual,RtString("Attribute"))
		So(len(list),ShouldEqual,4)
		
		So(list[0],RtShouldEqual,RtToken("float[2] angle"))
		So(list[1],RtShouldEqual,RtFloatArray{1.23,32.1})
		So(list[2],RtShouldEqual,RtToken("float tint"))
		So(list[3],RtShouldEqual,RtFloat(1.0))

	})
}
