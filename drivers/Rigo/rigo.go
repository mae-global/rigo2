package rigodriver

import (
	"sync"
	"log"
	"os"
	"fmt"
	"strings"
	
	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/mae-global/rigo2/drivers"
)

type rigoDriver struct {
	sync.RWMutex

	file *os.File
	last string
	header bool

	progress *ProtectedInteger
}

func (d *rigoDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {

}

func (d *rigoDriver) GetProgress() RtInt {
	d.progress.RLock()
	defer d.progress.RUnlock()
	return RtInt(d.progress.Value)
}

func (d *rigoDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {

	d.last = RIBStream(name,args,tokens,values)
	
	if !d.header {

		fmt.Fprintf(d.file,"%s",program_header)
		d.header = true
	}

	switch string(name) {
		case "#":
			out := ""
			if len(args) > 0 {
				out = args[0].String()
			}			
			fmt.Fprintf(d.file,"\tri.ArchiveRecord(\"comment\",\"%%s\",%s)\n",out)		
		break
		case "##":
			out := ""
			if len(args) > 0 {
				out = args[0].String()
			}
			if out == "\"RenderMan RIB\"" {
				return nil /* skip this */
			}

			fmt.Fprintf(d.file,"\tri.ArchiveRecord(\"structure\",\"%%s\",%s)\n",out)		
		break
		case "verbatim":
			out := ""
			if len(args) > 0 {
				out = args[0].String()
			}
			fmt.Fprintf(d.file,"\tri.ArchiveRecord(\"verbatim\",\"%%s\",%s)\n",out)		
		break
		default:

			fmt.Fprintf(d.file,"\tri.%s(",string(name))
			for i,arg := range args {
				if i > 0 {
					fmt.Fprintf(d.file,",")
				}
				fmt.Fprintf(d.file,"%s",Formalise(arg))
			}


			for i,tkn := range tokens {
				value := values[i]
				fmt.Fprintf(d.file,",RtToken(%s),Rt%s(%s)",tkn,strings.Title(value.Type()),value)
			}

		fmt.Fprintf(d.file,")\n")
		break
	}

	return nil
}

func (d *rigoDriver) Close() *RtError {

	if !d.header {
		fmt.Fprintf(d.file,"%s",program_header)
	}

	fmt.Fprintf(d.file,"%s",program_footer)

	d.file.Close()
	
	return nil
}

func (d *rigoDriver) GetLastRIB() string {
	return d.last
}

func BuildDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver,error) {

	d := &rigoDriver{}
	d.progress = new(ProtectedInteger)
	d.progress.Value = 100
	d.header = false

	
	filename := args[0]
	f,err := os.Create(filename)
	if err != nil {
		return nil,err
	}

	d.file = f

	/* FIXME, add in the Options */





	return d,nil
}

/* TODO: add the rest of the array classes here */
func Formalise(v RtPointer) string {

	var out string

	switch v.Type() {
		case "color":
			color := v.(RtColor)

			out = "RtColor{"
			for i,c := range color {
				if i > 0 {
					out += ","
				}
				out += c.String()
			}
			out += "}"
		break
		default: 
			out = v.String()
		break
	}

	return out
}





const program_header = `package main

import (
	. "github.com/mae-global/rigo2/ri/core"
	"github.com/mae-global/rigo2"
)

func main() {

	ri := rigo.New(nil)
	ri.Begin("-")

`

const program_footer = `	ri.End()
}
`






	
