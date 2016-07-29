package args

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func concat(list []string) string {
	out := ""
	for i, ele := range list {
		if i > 0 {
			out += ","
		}
		out += ele
	}
	return out
}

func Test_ArgsParser(t *testing.T) {

	Convey("Inline Args Parse", t, func() {

		info, err := Parse("PxrCamera", []byte(argsExample1))
		So(err, ShouldBeNil)
		So(info, ShouldNotBeNil)

		So(string(info.Name), ShouldEqual, "PxrCamera")
		So(string(info.ShaderType), ShouldEqual, "projection")
		So(string(info.NodeId), ShouldEqual, "")
		So(string(info.Classification), ShouldEqual, "")

		So(len(info.Pages), ShouldEqual, 6)
		So(len(info.Params), ShouldEqual, 0)
		So(len(info.Outputs), ShouldEqual, 0)

		So(string(info.Help), ShouldEqual, `A camera model that approximates a number of real world physical
    effects.  This supports all of the traditional prman perspective camera
    settings including shaped motion blur and bokeh.`)

	})
}

/* Taken directly from PxrCamera.args, Copyright Pixar */
const argsExample1 = `
<args format="1.0">
  <shaderType>
    <tag value="projection"/>
  </shaderType>
  <help>
    A camera model that approximates a number of real world physical
    effects.  This supports all of the traditional prman perspective camera
    settings including shaped motion blur and bokeh.
  </help>
  <page name="Standard Perspective" open="True">
    <param name="fov" label="Field of View" type="float" widget="default"
           default="90.0" min="1.0" max="135.0"
           connectable="False">      
    </param>
  </page>
  <page name="Tilt-Shift" open="False">
    <param name="tilt" label="Tilt Angle" type="float" widget="default"
           default="0.0" min="-20.0" max="20.0"
           connectable="False">      
    </param>
    <param name="roll" label="Roll Angle" type="float" widget="default"
           default="0.0" min="-180.0" max="180.0"
           connectable="False">
      <help>
        Roll the lens clockwise.  If the lens tilt is non-zero this can be
        used to rotate the plane of focus around the image center.
      </help>
    </param>
    <param name="shiftX" label="Shift X" type="float" widget="default"
           default="0.0" min="-1.0" max="1.0"
           connectable="False">
      <help>
        Shift the lens horizontally.  This can be used to correct for
        perspective distortion.  Positive values shift towards the right.
      </help>
    </param>
    <param name="shiftY" label="Shift Y" type="float" widget="default"
           default="0.0" min="-1.0" max="1.0"
           connectable="False">      
    </param>
  </page>
  <page name="Lens Distortion" open="False">
    <param name="radial1" label="Radial Distortion 1" type="float" widget="default"
           default="0.0" min="-0.3" max="0.3"
           connectable="False">      
    </param>
    <param name="radial2" label="Radial Distortion 2" type="float" widget="default"
           default="0.0" min="-0.3" max="0.3"
           connectable="False">      
    </param>
    <param name="assymX" label="Assymetric Distortion X" type="float" widget="default"
           default="0.0" min="-0.3" max="0.3"
           connectable="False">      
    </param>
    <param name="assymY" label="Assymetric Distortion Y" type="float" widget="default"
           default="0.0" min="-0.3" max="0.3"
           connectable="False">     
    </param>
    <param name="squeeze" label="Anamorphic Squeeze" type="float" widget="default"
           default="1.0" min="0.5" max="2.0"
           connectable="False">     
    </param>
  </page>
  <page name="Chromatic Aberration" open="False">
    <param name="transverse" label="Transverse" type="color" widget="default"
           default="1.0 1.0 1.0"
           connectable="False">     
    </param>
    <param name="axial" label="Axial" type="color" widget="default"
           default="0.0 0.0 0.0"
           connectable="False">     
    </param>
  </page>
  <page name="Vignetting" open="False">
    <param name="natural" label="Natural" type="float" widget="default"
           default="0.0" min="0.0" max="1.0"
           connectable="False">     
    </param>
    <param name="optical" label="Optical" type="float" widget="default"
           default="0.0" min="0.0" max="1.0"
           connectable="False">      
    </param>
  </page>
  <page name="Shutter" open="False">
    <param name="sweep" label="Sweep" type="string" widget="default"
           default="global"
           connectable="False">
      <hintdict name="options">
        <string name="down" value="down"/>
        <string name="right" value="right"/>
        <string name="up" value="up"/>
        <string name="left" value="left"/>
      </hintdict>     
    </param>
    <param name="duration" label="Duration" type="float" widget="default"
           default="1.0" min="0.0" max="1.0"
           connectable="False">     
    </param>
  </page>
</args>
`
