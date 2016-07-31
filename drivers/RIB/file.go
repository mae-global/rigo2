package ribdriver

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/mae-global/rigo2/drivers"
)

type fileDriver struct {
	sync.RWMutex

	file    *os.File
	last    string
	depth   int
	indent 	bool
	wide 		bool
}

func (d *fileDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* do nothing */
}

func (d *fileDriver) GetProgress() RtInt {
	return RtInt(100)
}

func (d *fileDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {

	out := ""

	if d.indent {
		switch string(name) {
		case "AttributeBegin", "FrameBegin", "MotionBegin", "ObjectBegin", "SolidBegin", "TransformBegin", "WorldBegin":

			defer func() { d.depth++ }()
			break
		case "AttributeEnd", "FrameEnd", "MotionEnd", "ObjectEnd", "SolidEnd", "TransformEnd", "WorldEnd":
			d.depth--
			break
		}

		for i := 0; i < d.depth; i++ {
			out += "\t"
		}
	}

	out += RIBStream(name, args, tokens, values)
	fmt.Fprintf(d.file, "%s\n", out)
	d.last = out
	return nil
}

func (d *fileDriver) Close() *RtError {
	d.file.Close()
	return nil
}

func (d *fileDriver) GetLastRIB() string {
	return d.last
}

func BuildFileDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {
	
	logger.Printf("Building RIB File Driver, options=%s, args=%s\n",options,args)
	
	filename := args[0]
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	d := &fileDriver{}
	d.file = f

	params,values := Unmix(options)
	for i,param := range params {
		value := values[i]
		tkn,ok := param.(RtToken)
		if !ok {
			continue
		}
		switch string(tkn) {
			case "asciistyle":
				/* comma seperated list */
				val,ok := value.(RtString)
				if !ok {
					continue
				}
				parts := strings.Split(string(val),",")
				/* go through the parts and setup */
				for _,part := range parts {
					switch part {
						case "indented":
							d.indent = true
						break
						case "wide":
							d.wide = true
						break
					}
			}
			break
			/* TODO: add the rest in */
		}
	}
	
	return d, nil
}

