package rma

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"fmt"
)

func Test_RMA(t *testing.T) {

	Convey("RMA test",t,func() {

		asset,err := Load("/Materials/Fabric/beige_fabric.rma")
		So(err,ShouldBeNil)
		So(asset,ShouldNotBeNil)

	})

	Convey("RMA load library",t,func() {

		lib,err := LoadLibrary()
		So(err,ShouldBeNil)
		So(lib,ShouldNotBeNil)

		lib.PrettyPrint()
		
		fmt.Printf("depth of library is %d\n",lib.Depth())

		node := lib.Find("yellow matte plastic")
		So(node,ShouldNotBeNil)

		node.PrettyPrint()
	})
}
