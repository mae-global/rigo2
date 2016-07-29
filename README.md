#RiGO2  ![version](https://img.shields.io/badge/version-2-red.svg)

Implementation of the RenderMan Interface for the Go programming language. This is currently 
based on Pixar's RenderMan Specification version 3.2.1 (November 2005). This implementation 
is still under *active development*, so *expect* holes and bugs. 

[Online Documentation](https://godoc.org/github.com/mae-global/rigo2)

Install with:

    go get github.com/mae-global/rigo2


Example usage: 

```go

import (
  . "github.com/mae-global/rigo2/ri/core"
  . "github.com/mae-global/rigo2/ri"
  "github.com/mae-global/rigo2"
)

/* create a context to work with */
ri := rigo.New(&rigo.Configuration{PrettyPrint:true})

ri.Begin("unitcube.rib")
ri.AttributeBegin()
	ri.Attribute("identifier", RtToken("name"), RtToken("unitcube"))
	ri.Bound(RtBound{-.5, .5, -.5, .5, -.5, .5})
	ri.TransformBegin()

		points := RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5}

		ri.ArchiveRecord(COMMENT,"far face")
		ri.Polygon(4, P, points)
		ri.Rotate(90, 0, 1, 0)

		ri.ArchiveRecord(COMMENT,"right face")
		ri.Polygon(4, P, points)
		ri.Rotate(90, 0, 1, 0)

		ri.ArchiveRecord(COMMENT,"near face")
		ri.Polygon(4, P, points)
		ri.Rotate(90, 0, 1, 0)

		ri.ArchiveRecord(COMMENT,"left face")
		ri.Polygon(4, P, points)

	ri.TransformEnd()
	ri.TransformBegin()

		ri.ArchiveRecord(COMMENT,"bottom face")
		ri.Rotate(90, 1, 0, 0)
		ri.Polygon(4, P, points)

		ri.TransformEnd()
		ri.TransformBegin()

		ri.ArchiveRecord(COMMENT,"top face")
		ri.Rotate(-90, 1, 0, 0)
		ri.Polygon(4, P, points)

	ri.TransformEnd()
ri.AttributeEnd()
ri.End()	
```

RIB output of *unitcube.rib* is thus :-

```
##RenderMan RIB
AttributeBegin 
	Attribute "identifier" "name" "unitcube"
	Bound [-.5 .5 -.5 .5 -.5 .5]
	TransformBegin 
		# far face
		Polygon "P" [.5 .5 .5 -.5 .5 .5 -.5 -.5 .5 .5 -.5 .5]
		Rotate 90. 0 1. 0
		# right face
		Polygon "P" [.5 .5 .5 -.5 .5 .5 -.5 -.5 .5 .5 -.5 .5]
		Rotate 90. 0 1. 0
		# near face
		Polygon "P" [.5 .5 .5 -.5 .5 .5 -.5 -.5 .5 .5 -.5 .5]
		Rotate 90. 0 1. 0
		# left face
		Polygon "P" [.5 .5 .5 -.5 .5 .5 -.5 -.5 .5 .5 -.5 .5]
	TransformEnd 
	TransformBegin 
		# bottom face
		Rotate 90. 1. 0 0
		Polygon "P" [.5 .5 .5 -.5 .5 .5 -.5 -.5 .5 .5 -.5 .5]
	TransformEnd 
	TransformBegin 
		# top face
		Rotate -90. 1. 0 0
		Polygon "P" [.5 .5 .5 -.5 .5 .5 -.5 -.5 .5 .5 -.5 .5]
	TransformEnd 
AttributeEnd 
```




###Information

RenderMan Interface Specification is Copyright © 2005-2016 Pixar.
RenderMan © is a registered trademark of Pixar.

