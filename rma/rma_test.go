package rma

import (
	"os"	
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"fmt"
)

var localtest bool

func init() {
	debug := os.Getenv("RIGO2_DEBUG")
	if debug == "testing" {
		localtest = false
	} else {
		localtest = true
	}
}

func Test_RMA(t *testing.T) {

	Convey("RMA test",t,func() {

		if !localtest {
			t.Skip()
		}

		asset,err := Load("/Materials/Fabric/beige_fabric.rma")
		So(err,ShouldBeNil)
		So(asset,ShouldNotBeNil)

	})

	Convey("RMA load library",t,func() {

		if !localtest {
			t.Skip()
		}

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
