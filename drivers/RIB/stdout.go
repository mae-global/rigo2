package ribdriver

import (
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/mae-global/rigo2/drivers"
)

type stdoutDriver struct {
	last    string
	depth   int
	proc    bool /* is used by procedural calls, will append \377 to stdout */
	indent  bool /* tab */
	wide    bool /* do not auto carriage return long statements */
}

func (d *stdoutDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* do nothing */
}

func (d *stdoutDriver) GetProgress() RtInt {
	return RtInt(100)
}

func (d *stdoutDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {

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
	/* FIXME, take into account "wide" */
	fmt.Fprintf(os.Stdout, "%s\n", out)
	d.last = out
	return nil
}

func (d *stdoutDriver) Close() *RtError {
	if d.proc {
		fmt.Fprintf(os.Stdout, "\377")
	}
	return nil
}

func (d *stdoutDriver) GetLastRIB() string {
	return d.last
}

func BuildStdoutDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {

	logger.Printf("Building RIB Stdout Driver, options=%s, args=%s\n",options,args)

	d := &stdoutDriver{}

	for _, arg := range args {
		if arg == "proc" {
			d.proc = true
		}
	}

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

