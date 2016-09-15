package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	. "github.com/mae-global/rigo2/ri/core"

	"fmt"
)

func Test_Stream(t *testing.T) {

	Convey("Stream via protobuffers (3)",t,func() {

		data,err := stream("RiSphere",[]RtPointer{RtFloat(1),RtFloat(-1),RtFloat(1),RtFloat(360)})
		So(err,ShouldBeNil)
		So(data,ShouldNotBeNil)

		name,values,err := unstream(data)
		So(err,ShouldBeNil)
		So(name,ShouldEqual,"RiSphere")
		So(len(values),ShouldEqual,4)
		So(values[0],ShouldEqual,RtFloat(1))
		So(values[1],ShouldEqual,RtFloat(-1))
		So(values[2],ShouldEqual,RtFloat(1))
		So(values[3],ShouldEqual,RtFloat(360))
		
		fmt.Printf("%s %s\n",name,values)

	})
}
