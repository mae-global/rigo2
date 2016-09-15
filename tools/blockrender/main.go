package main

import (
	"fmt"
	"os"
	"log"
	"bytes"
	"strings"
	"text/template"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	"github.com/mae-global/rigo2/drivers"
	"github.com/mae-global/rigo2"
)

/* TODO: at the moment this is just development code, in the future this will
 * be a useful tool for render block diagrams of RIB files 
 */

func main() {

	rigo.AddDriver("block",builder)

	ri := rigo.New(nil) 

	/* Nice Simple example, should contain a single block, with two assets */
	ri.Begin("block simple.html")
	
		ri.Display("simplesphere.tiff", "file", "rgba")
		ri.Format(320, 240, 1)
		ri.Projection(PERSPECTIVE, RtToken("float fov"), RtFloat(30))
		ri.Translate(0, 0, 6)
		ri.WorldBegin()
			ri.LightSource("ambientlight", "-", RtToken("float intensity"), RtFloat(0.5))
			ri.Color(RtColor{1, 0, 0})
			ri.Sphere(1, -1, 1, 360)
		ri.WorldEnd()

	ri.End()	




	
	/* More involved example */	
	ri.Begin("block involved.html")
		ri.Display("redsphere.tiff", "file", "rgba")
		ri.Format(320, 240, 1)
		ri.ShadingRate(100)
		ri.ShadingInterpolation("smooth")
		ri.PixelVariance(0.001)
		ri.PixelFilter(GaussianFilter, RtFloat(3), RtFloat(3))
		ri.Projection("perspective", RtToken("float fov"), RtFloat(30))

		ri.Imager("background", RtToken("color color"), RtColor{1, 1, 1}, RtToken("float aplha"), RtFloat(1))

		ri.Hider("raytrace", RtToken("int maxsamples"), RtInt(64), RtToken("int minsamples"), RtInt(16),
			RtToken("int incremental"), RtInt(12), RtToken("float[4] aperture"), RtFloatArray{0, 0, 0, 0},
			RtToken("string integrationmode"), RtString("path"))

		ri.Integrator("PxrPathTracer", "redbox", RtToken("int maxPathLength"), RtInt(20),
			RtToken("string sampleMOD"), RtString("bxdf"), RtToken("int rouletteDepth"), RtInt(12),
			RtToken("float rouletteThreshold"), RtFloat(0.2), RtToken("int clampDepth"), RtInt(2),
			RtToken("int clampLuminance"), RtInt(10))

		ri.FrameBegin(1)
			ri.WorldBegin()
				ri.Translate(0, 0, 15)

				ri.AttributeBegin()
					ri.Attribute("identifier",RtToken("name"),RtString("light"))
					ri.Translate(5, 5, 5)

						light := ri.AreaLightSource("PxrStdEnvDayLight", RtLightHandle("-"), RtToken("float importance"), RtFloat(2),
							RtToken("float exposure"), RtFloat(1), RtToken("vector directionVector"),
							RtVector{1, 1, 1}, RtToken("color specAmount"), RtColor{.5, .5, .5},
							RtToken("float haziness"), RtFloat(1.7), RtToken("float enableShadows"),
							RtFloat(1))

					ri.Bxdf("PxrLightEmission", RtToken(light))
					ri.Geometry("envsphere", RtToken("constant float[2] resolution"), RtFloatArray{1024, 1024})
				ri.AttributeEnd()

				ri.Illuminate(light, true)

				ri.AttributeBegin()
					ri.Attribute("identifier",RtToken("name"),RtString("red-sphere"))
					ri.Bxdf("PxrDisney", "-", RtToken("color baseColor"), RtColor{1, 0, 0})
					ri.Sphere(1, -1, 1, 360)
				ri.AttributeEnd()
			ri.WorldEnd()
		ri.FrameEnd()

		/* wait on the render progress to end */
		ri.Utils.Wait(func(progress RtInt) bool {
			fmt.Fprintf(os.Stderr,"progress %03d%%\n", progress)
			if int(progress) < 100 {
				return false
			}
			return true
		})

	ri.End()

	rigo.RemoveDriver("block")
}

type Asset struct {
	Name RtString
}


type Block struct {
	Name RtString
	Label RtString /* generally taken from any Attribute "name" */
	Instructions RtInt
	Assets []*Asset	
	Parent *Block
	Children []*Block
}

func NewBlock(name RtString) *Block {
	b := &Block{}
	b.Name = name
	b.Instructions = 0
	b.Assets = make([]*Asset,0)
	b.Children = make([]*Block,0)
	return b
}


type custom struct {
	depth int
	last string
	output string /* output html file */

	file *os.File

	root *Block
	current *Block
}

func (c *custom) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) { }
func (c *custom) GetProgress() RtInt { return RtInt(100) }
func (c *custom) Handle(name RtString,args,tokens,values []RtPointer) *RtError {

	/* TODO: recreate the Ri spec, or the important commands that we are interested in */
		
	switch string(name) {
		/* TODO: FrameBegin [frame] can be used as a label */
		case "AttributeBegin", "FrameBegin", "MotionBegin", "ObjectBegin", "SolidBegin", "TransformBegin", "WorldBegin","IfBegin":
		
			for i := 0; i < c.depth; i++ {
				fmt.Printf("\t")
			}
			
			fmt.Printf("+ %s\n",name)
			
			c.depth ++

			b := NewBlock(name)
			b.Parent = c.current
			c.current.Children = append(c.current.Children,b)
			c.current = b

			out,_ := Render(TemplateBeginBlock,nil)
			c.file.Write(out)

		break
		case "AttributeEnd", "FrameEnd", "MotionEnd", "ObjectEnd", "SolidEnd", "TransformEnd", "WorldEnd","IfEnd":

			c.depth --
			for i := 0; i < c.depth; i++ {
				fmt.Printf("\t")
			}
	
			fmt.Printf("%d instructions, %d assets label=%s\n",c.current.Instructions,len(c.current.Assets),c.current.Label)

			c.current = c.current.Parent

			/* TODO: actually render block before popping */
			out,_ := Render(TemplateEndBlock,nil)
			c.file.Write(out)	

		break
		case "Attribute":
			if s,ok := args[0].(RtString); ok {
				if string(s) == "identifier" {
				
					/* go through and find the "name" token if present */
					for i,token := range tokens {
						if t,ok := token.(RtToken); ok {
							info := Specification(string(t))
							if info.Name == "name" {
								if v,ok := values[i].(RtString); ok {
									c.current.Label = v
								}
							}
						}
					}
				}
			}
		break
		/*
		case "#","##":

			for i := 0; i < c.depth; i++ {
				fmt.Printf("\t")
			}

			fmt.Printf("%s\n",name)

			c.current.Instructions ++
		break
		*/

		case "Sphere":
			
			c.current.Assets = append(c.current.Assets,&Asset{Name:name})
			out,_ := Render(TemplateAsset,nil)
			c.file.Write(out)			


		break				
		default:
			c.current.Instructions ++
		
		break
	}

	return nil
} 

func (c *custom) Close() *RtError { 
	
	if c.file != nil {

		b,_ := Render(TemplateFooter,nil)
		c.file.Write(b)
		c.file.Close()
	}

	return nil 
}

func (c *custom) GetLastRIB() string { return c.last }


/* custom driver builder */
func builder(logger *log.Logger, options []RtPointer, args ...string) (drivers.Driver,error) {
	
	logger.Printf("custom driver loaded args=%v\n",args)

	c := &custom{}
	c.root = NewBlock("root")
	c.current = c.root
	c.output = "out.html"

	if len(args) > 0 {
		if strings.HasSuffix(args[0],".html") {
			c.output = args[0]
		}
	} 

	/* open file to output */
	file,err := os.Create(c.output)
	if err != nil {
		return nil,err
	}

	c.file = file

	b,err := Render(TemplateHeader,nil)
	if err != nil {
		c.file.Close()
		return nil,err
	}

	c.file.Write(b)


	return c,nil
}

func Render(f string,v interface{}) ([]byte,error) {

	t,err := template.New("temp").Funcs(template.FuncMap{}).Parse(f)
	if err != nil {
		return nil,err
	}

	b := bytes.NewBuffer(nil)

	if err := t.Execute(b,v); err != nil {
		return nil,err
	}

	return b.Bytes(),nil
}
		



const TemplateHeader string = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="">
  <meta name="author" content="">
  <link rel="icon" href="favicon.ico">

  <title>blockrender</title>

	<style>
	
		div { display: block; }
		div.block { margin-left: 10px; border-top: 1px solid #000; padding-left: 5px; }
		div.asset { background-color: pink; padding: 30px; }


	</style>
</head>
<body>
`

const TemplateFooter string = `</body>
</head>
`

const TemplateBeginBlock string = `<div class="block">
	<p>block</p>
`

const TemplateEndBlock string = `</div>`

const TemplateAsset string = `<div class="asset"></div>`




























