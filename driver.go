package rigo

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
)
const (
	DefaultBlockFile = "out.rib.block"
)



type Driver interface {
	Flush(marker RtString, synchronous RtBoolean, flushmode RtToken)
	GetProgress() RtInt
	Handle(RtString, []RtPointer, []RtPointer, []RtPointer) *RtError
	Close() *RtError
	GetLastRIB() string
}

/* construction function for creating a new driver */
type DriverBuilder func(logger *log.Logger, options []RtPointer, args ...string) (Driver,error)

var internal struct {
	sync.RWMutex
	Drivers map[string]DriverBuilder
}

func init() {
	
	/* build the default drivers table */
	dd := make(map[string]DriverBuilder,0)
	dd["block"] = BuildBlockDiagrammingDriver
	dd["stdout"] = BuildRIBStdoutDriver
	dd["catrib"] = BuildCatribDriver
	dd["render"] = BuildRenderDriver
	dd["file"]   = BuildRIBFileDriver
	
	internal.Drivers = dd
}


func AddDriver(name string, builder DriverBuilder) error {
	internal.Lock()
	defer internal.Unlock()
	internal.Drivers[name] = builder
	return nil /* TODO */
}

func RemoveDriver(name string) error {
	internal.Lock()
	defer internal.Unlock()
	delete(internal.Drivers,name)
	return nil
}


type ProtectedInteger struct {
	sync.RWMutex
	Value int
}

/* block */
type BlockDiagrammingDriver struct {
	sync.RWMutex

	last string

	file io.WriteCloser

	root,current *BdiaBlock
}

func (d *BlockDiagrammingDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* do nothing */
}

func (d *BlockDiagrammingDriver) GetProgress() RtInt {
	return RtInt(100)
}

/* Block Diagramming information structs */
type BdiaAsset struct {
	Name string
}

type BdiaBlock struct {
	Name string
	
	Parent *BdiaBlock
	Children []interface{}
}

func (d *BlockDiagrammingDriver) Handle(name RtString,args []RtPointer,tokens []RtPointer,values []RtPointer) *RtError {

	/* TODO: build up a block diagramming set, output at logical points in the parsing
   *       such as WorldEnd and FrameEnd 
	 */

	switch string(name) {
		case "AttributeBegin", "FrameBegin", "MotionBegin", "ObjectBegin", "SolidBegin", "TransformBegin", "WorldBegin","IfBegin":

			/* start a new block as a child of the current */
			block := new(BdiaBlock)
			block.Parent = d.current
			block.Children = make([]interface{},0)

			block.Name = strings.TrimSuffix(string(name),"Begin")

			d.current.Children = append(d.current.Children,block)
			d.current = block
					
		break
		case "AttributeEnd", "FrameEnd", "MotionEnd", "ObjectEnd", "SolidEnd", "TransformEnd", "WorldEnd","IfEnd":
			/* Block-wise we write the attribute out when we reach the end, thus knowing what is inside */
	
			d.current = d.current.Parent

		break
		case "Sphere","Cone","Cylinder","hyberboloid","Paraboloid","Disk","Torus":
			
			asset := new(BdiaAsset)
			asset.Name = string(name)
	
			d.current.Children = append(d.current.Children,asset)

		break
		default:

		break
	}

	return nil
}

/* ASCII renderer === */

func writehelper(file io.WriteCloser,node interface{}) {
	if node == nil {
		return
	}

	/* asset : */
	if asset,ok := node.(*BdiaAsset); ok {
		fmt.Fprintf(file,"---> %s\n",asset.Name)
		return
	}

	/* block : */
	block,ok := node.(*BdiaBlock)
	if !ok {
		return
	}


	fmt.Fprintf(file,"%s\n",block.Name)		

	for _,node := range block.Children {
		writehelper(file,node)
	}
}	

func (d *BlockDiagrammingDriver) Close() *RtError {

	/* do all the writing here */
	for _,block := range d.root.Children {
		writehelper(d.file,block)
	}

	if d.file != nil {
		d.file.Close()
	}
	return nil
}

func (d *BlockDiagrammingDriver) GetLastRIB() string {
	return d.last
}



func BuildBlockDiagrammingDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver,error) {

	d := &BlockDiagrammingDriver{}

	out := DefaultBlockFile
	if len(args) > 0 {
		out = args[0]
	}
	/* FIXME: do a check of the args */

	filename := out
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	d.file = f
	d.root = new(BdiaBlock)
	d.root.Parent = nil
	d.root.Children = make([]interface{},0)

	d.current = d.root
	
	return d,nil
}


/* Debug */

type DebugDriver struct {
	sync.RWMutex

	last    string
	depth   int
}

func (d *DebugDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* do nothing */
}

func (d *DebugDriver) GetProgress() RtInt {
	return RtInt(100)
}

func (d *DebugDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {

	out := ""

	/* FIXME: check indent is on */
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
	

	out += RIBStream(name, args, tokens, values)
	d.last = out
	return nil
}

func (d *DebugDriver) Close() *RtError {
	return nil
}

func (d *DebugDriver) GetLastRIB() string {
	return d.last
}

func BuildDebugDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {

	d := &DebugDriver{}
	d.last = ""

	return d, nil
}

type RIBFileDriver struct {
	sync.RWMutex

	file    *os.File
	last    string
	depth   int
	indent 	bool
	wide 		bool
}

func (d *RIBFileDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* do nothing */
}

func (d *RIBFileDriver) GetProgress() RtInt {
	return RtInt(100)
}

func (d *RIBFileDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {

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

func (d *RIBFileDriver) Close() *RtError {
	d.file.Close()
	return nil
}

func (d *RIBFileDriver) GetLastRIB() string {
	return d.last
}

func BuildRIBFileDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {
	
	logger.Printf("Building RIB File Driver, options=%s, args=%s\n",options,args)
	
	filename := args[0]
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	d := &RIBFileDriver{}
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
						case "indent":
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

type RIBStdoutDriver struct {
	last    string
	depth   int
	proc    bool /* is used by procedural calls, will append \377 to stdout */
	indent  bool /* tab */
	wide    bool /* do not auto carriage return long statements */
}

func (d *RIBStdoutDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* do nothing */
}

func (d *RIBStdoutDriver) GetProgress() RtInt {
	return RtInt(100)
}

func (d *RIBStdoutDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {

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

func (d *RIBStdoutDriver) Close() *RtError {
	if d.proc {
		fmt.Fprintf(os.Stdout, "\377")
	}
	return nil
}

func (d *RIBStdoutDriver) GetLastRIB() string {
	return d.last
}

func BuildRIBStdoutDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {

	logger.Printf("Building RIB Stdout Driver, options=%s, args=%s\n",options,args)

	d := &RIBStdoutDriver{}

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
						case "indent":
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

type CatribDriver struct {
	sync.RWMutex

	last string
	cmd  *exec.Cmd
	in   io.WriteCloser
	out  io.ReadCloser
	err  io.ReadCloser
}

func (d *CatribDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* do nothing */
}

func (d *CatribDriver) GetProgress() RtInt {
	return RtInt(100)
}

func (d *CatribDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {
	d.last = RIBStream(name, args, tokens, values)
	fmt.Fprintf(d.in, "%s\n", d.last)
	return nil
}

func (d *CatribDriver) Close() *RtError {

	d.out.Close()
	//d.cmd.Wait()
	return nil
}

func (d *CatribDriver) GetLastRIB() string {
	return d.last
}

func BuildCatribDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {

	//logger.Printf("CatribDriver options = %s\n",options)

	cmd := exec.Command("catrib", args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	d := &CatribDriver{}
	d.cmd = cmd
	d.in = stdin
	d.out = stdout
	d.err = stderr

	return d, d.cmd.Start()
}

type RenderDriver struct {
	sync.RWMutex

	last string
	cmd  *exec.Cmd
	in   io.WriteCloser
	out  io.ReadCloser
	err  io.ReadCloser

	progress *ProtectedInteger
}

func (d *RenderDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* TODO: talk to renderer via Ric */
}

func (d *RenderDriver) GetProgress() RtInt {
	d.progress.RLock()
	defer d.progress.RUnlock()
	return RtInt(d.progress.Value)
}

func (d *RenderDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {
	d.last = RIBStream(name, args, tokens, values)
	fmt.Fprintf(d.in, "%s\n", d.last)
	return nil
}

func (d *RenderDriver) Close() *RtError {

	d.out.Close()

	/* wait on the process to finish */
	//d.cmd.Wait()

	return nil
}

func (d *RenderDriver) GetLastRIB() string {
	return d.last
}

func BuildRenderDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {

	//logger.Printf("RenderDriver options = %s\n",options)

	cmd := exec.Command("render", args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	d := &RenderDriver{}
	d.cmd = cmd
	d.in = stdin
	d.out = stdout
	d.err = stderr
	d.progress = new(ProtectedInteger)

	go func(pi *ProtectedInteger) {
		buffer := make([]byte, 256)
		for {
			n, err := stderr.Read(buffer)
			if err != nil {
				if err == io.EOF {
					return
				}

				logger.Fatal(err)
			}

			str := string(buffer[:n])
			/* basic progress parser */
			end := -1
			for i, c := range str {
				if c == '%' {
					end = i
					break
				}
			}

			if end != -1 && end > 0 {
				p, err := strconv.Atoi(strings.TrimSpace(str[:end]))
				if err != nil {
					logger.Printf("unable to parse progress -- %v (%s)\n", err, str)
					continue
				}

				pi.Lock()
				pi.Value = p
				pi.Unlock()
			}
		}
	}(d.progress)

	return d, d.cmd.Start()
}





