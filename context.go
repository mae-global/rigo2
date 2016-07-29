package rigo

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"

	"github.com/mae-global/rigo2/rie"
)

/* HandleGeneratorHandler -- generate handles, RtLightHandle, ... */
type HandleGeneratorHandler func(name, typeof string) (string, error)

func DefaultHandleGeneratorHandler(name, typeof string) (string, error) {

	if name == "" || name == "-" {

		b := make([]byte, 8)
		n, err := io.ReadFull(rand.Reader, b)
		if err != nil {
			return "-", err
		}
		return typeof + "_" + hex.EncodeToString(b[:n]), nil
	}

	return name, nil
}

type ErrorHandler func(code, severity int, msg string) error

func AbortErrorHandler(code, severity int, msg string) error {
	err := fmt.Errorf("%05d:%d -- %s\n", code, severity, msg)
	fmt.Fprintf(os.Stdout, "%v\n", err)
	/* loop the msg back forcing a panic */
	return err
}

func PrintErrorHandler(code, severity int, msg string) error {
	err := fmt.Errorf("%05d:%d -- %s\n", code, severity, msg)
	fmt.Fprintf(os.Stdout, "%v\n", err)
	return nil
}

/* Configuration */
type Configuration struct {
	Debug       bool /* use debug sequence */
	PrettyPrint bool /* use the pretty print for RIB output */
	ReadAhead   bool /* use type readahead */
	Strict      bool /* use strict checking of types */
	Fragment    bool /* don't do any automagic stuff */

	Errorf  ErrorHandler /* Error Handler function */
	Handlef HandleGeneratorHandler
	Logger  *log.Logger
}

type BlockInfo struct {
	Type string
	Annotation string
	Statements int
	Tokens int
	InlineTokens int
	Assets []string         /* lights, shaders etc */
}

func (info *BlockInfo) String() string {
	assets := func(n int) string {
		if n == 1 {
			return "1 asset"
		}
		return fmt.Sprintf("%d assets",n)
	}

	return fmt.Sprintf("%s %s\t-- (statements %d, tokens [%d/%d], %s)",info.Type,info.Annotation,info.Statements,info.InlineTokens,info.Tokens,assets(len(info.Assets)))
}

type Block struct {
	Info *BlockInfo

	Parent *Block    /* if nil then root */
	Children []*Block
}





/* Statistics -- run per begin/end statements TODO: move to utilities and standardise */
type Statistics struct {
	sync.RWMutex

	Tokens map[RtToken] int    /* usage of tokens (by name only) */
	Root *Block
	Current *Block
}

func (s *Statistics) IncrementStatements() {
	s.Lock()
	defer s.Unlock()

	if s.Current == nil {
		s.Current = new(Block)
		if s.Root == nil {	
			s.Root = s.Current
		}

		s.Current.Info = new(BlockInfo)
	}

	s.Current.Info.Statements ++
}

func (s *Statistics) IncrementTokens(n,inline int) {
	s.Lock()
	defer s.Unlock()

	if s.Current == nil {
		s.Current = new(Block)
		if s.Root == nil {
			s.Root = s.Current
		}
	
		s.Current.Info = new(BlockInfo)
	}

	s.Current.Info.Tokens += n
	s.Current.Info.InlineTokens += inline
}

func (s *Statistics) AnnotateLabel(token RtToken,value RtPointer) {
	s.Lock()
	defer s.Unlock()

	if s.Current == nil {
		s.Current = new(Block)
		if s.Root == nil {
			s.Root = s.Current
		}
		s.Current.Info = new(BlockInfo)
	}

	s.Current.Info.Annotation = value.String()
}

func (s *Statistics) AddAsset(token RtToken,value RtPointer) {
	s.Lock()
	defer s.Unlock()

	if s.Current == nil {
		s.Current = new(Block)
		if s.Root == nil {
			s.Root = s.Current
		}
		s.Current.Info = new(BlockInfo)
	}

	if s.Current.Info.Assets == nil {
		s.Current.Info.Assets = make([]string,0)
	}

	s.Current.Info.Assets = append(s.Current.Info.Assets,value.String())
}


func (s *Statistics) CurrentBlock() string {
	s.RLock()
	defer s.RUnlock()

	if s.Current == nil {
		return ""
	}

	return s.Current.Info.Type
}

func (s *Statistics) Begin(name string) {
	s.Lock()
	defer s.Unlock()

	block := new(Block)
	block.Info = new(BlockInfo)
	block.Info.Type = name

	block.Parent = s.Current	
	
	if s.Current != nil {
		if s.Current.Children == nil {
			s.Current.Children = make([]*Block,0)
		}
		s.Current.Children = append(s.Current.Children,block)
	}

	s.Current = block
}

func (s *Statistics) End() {
	s.Lock()
	defer s.Unlock()

	if s.Current != nil {
		s.Current = s.Current.Parent
	}
}
		
func blockprint(depth int,root *Block) string {
	if root == nil {
		return ""
	}

	out := ""
	tabs := ""
	for i := 0; i < depth; i++ {
		tabs += "\t"
	}
	
	if root.Children != nil {		
		for _,block := range root.Children {	
			out += blockprint(depth + 1,block)
		}
	}

	return tabs + root.Info.String() + "\n" + out
}


/* currently prettyprint to string */
func (s *Statistics) PrettyPrint() string {
	s.RLock()
	defer s.RUnlock()

	times := func(n int) string {
		if n == 1 {
			return fmt.Sprintf("seen once")
		}
		return fmt.Sprintf("seen %d times",n)
	}

	out := fmt.Sprintf("statistics -- %d tokens\n",len(s.Tokens))
	for t,v := range s.Tokens {
		out += fmt.Sprintf("\t%s, %s\n",t,times(v))
	}	

	out += "\n"

	if s.Root == nil {
		out += "blocks -- none\n"
	} else {
		out += "blocks --\n"

		out += blockprint(0,s.Root)
	}

	return out
}

func (s *Statistics) IncrementToken(name RtToken) {
	s.Lock()
	defer s.Unlock()

	if _,exists := s.Tokens[name]; !exists {
			s.Tokens[name] = 0
	}

	s.Tokens[name] ++
}


/* Context */
type Context struct {
	sync.RWMutex
	dict          map[string]string /* technically the dictionary can be implemented seperately from the context */
	open          bool              /* has Begin been called? */
	errorhandler  ErrorHandler
	handlehandler HandleGeneratorHandler

	seen   bool /* track if seen RenderMan RIB */
	logger *log.Logger

	config Configuration
	driver Driver

	statistics *Statistics
}

/* NewContext */
func NewContext(config *Configuration) *Context {

	if config == nil {
		config = new(Configuration)
	}

	ctx := new(Context)
	ctx.dict = make(map[string]string, 0)

	ctx.config.Debug = config.Debug
	ctx.config.PrettyPrint = config.PrettyPrint
	ctx.config.ReadAhead = config.ReadAhead
	ctx.config.Strict = config.Strict
	ctx.config.Fragment = config.Fragment

	if config.Errorf != nil {
		ctx.errorhandler = config.Errorf
	}

	if config.Handlef != nil {
		ctx.handlehandler = config.Handlef
	} else {
		ctx.handlehandler = DefaultHandleGeneratorHandler
	}

	if config.Logger != nil {
		ctx.logger = config.Logger
	} else {
		ctx.logger = log.New(os.Stderr, "rigo2: ", log.Lshortfile)
	}

	ctx.statistics = new(Statistics)
	ctx.statistics.Tokens = make(map[RtToken]int,0)


	return ctx
}

/* peel a copy of the statistics */
func (ctx *Context) PeelStatistics() *Statistics {
	return ctx.statistics
}



func (ctx *Context) fatal(err interface{}) {
	if ctx.logger != nil {
		ctx.logger.Fatal(err)
	} else {
		panic(err)
	}
}

/* Set */
func (ctx *Context) Set(name, declaration string) RtToken {
	ctx.Lock()
	defer ctx.Unlock()
	ctx.dict[name] = declaration
	return RtToken(declaration + " " + name)
}

/* Get */
func (ctx *Context) Get(name string) string {
	ctx.RLock()
	defer ctx.RUnlock()
	if declaration, exists := ctx.dict[name]; exists {
		return declaration
	}
	return ""
}

func (ctx *Context) GenHandle(name, typeof string) (string, error) {
	return ctx.handlehandler(name, typeof)
}

/* HandleError */
func (ctx *Context) HandleError(err *RtError) *RtError {
	if ctx.errorhandler != nil {
		if err2 := ctx.errorhandler(err.Code, err.Severity, err.Msg); err2 != nil {
			return Error(err.Code, err.Severity, err2.Error())
		}
		return nil
	}

	/* else use abort handler */
	if ctx.config.Debug {
		if err2 := AbortErrorHandler(err.Code, err.Severity, err.Msg); err2 != nil {
			return Error(err.Code, err.Severity, err2.Error())
		}
	}

	if err2 := PrintErrorHandler(err.Code, err.Severity, err.Msg); err2 != nil {
		return Error(err.Code, err.Severity, err2.Error())
	}

	return nil
}

/* Handle */
func (ctx *Context) Handle(list []RtPointer) {

	name := list[0].(RtString)
	args := make([]RtPointer, 0)
	params := make([]RtPointer, 0)
	annotations := make([]RtPointer,0)	
	tokens2 := make([]RtToken, 0)
	
	inargs := false
	inparams := false
	inannotation := false
	trigger := -1

	for i, param := range list {
		/* RtToken ---- */
		if p, ok := param.(RtToken); ok {

			switch string(p) {
			case string(ARGUMENTS):
				inparams = false
				inargs = true
				inannotation = false
				trigger = i
				break
			case string(PARAMETERLIST):
				inparams = true
				inargs = false
				inannotation = false
				trigger = i
				break
			case string(ANNOTATIONS):
				inannotation = true
				inargs = false	
				inparams = false
				trigger = i
				break
			default:
				if string(p) != "|" && !inannotation {
					/* record in the statistics */
					ctx.statistics.IncrementToken(p)
				}
			
				break
			}
		}

		if trigger == i {
			continue
		}

		if inargs {
			args = append(args, param)
		}

		if inparams {
			params = append(params, param)
		}

		if inannotation {
			annotations = append(annotations, param)
		}
	}

	tokens, values := Unmix(params)

	/* go through all the parameters and check
		 * the specification of the token supplied
	   * using the dictionary to improve the
	   * information for later in the sequence
	*/

	/* Copy global dictionary */
	dict := make(map[string]string, 0)
	ctx.RLock()
	for key, value := range ctx.dict {
		dict[key] = value
	}
	ctx.RUnlock()

	inlineparams := 0 /* number of tokens that are inline */

	for i, param := range tokens {

		if token, ok := param.(RtToken); ok {

			/* break the token down into the specification/declaration parts */
			info := Specification(string(token))
			if len(info.Name) == 0 {
				if err := ctx.HandleError(Errorf(rie.BadToken, rie.Error, "bad parameter, invalid token -- %s", token)); err != nil {
					log.Fatal(err)
				}
				return
			}
	
			if info.IsInline() {
				inlineparams ++
			}

			if declaration, exists := dict[info.Name]; exists {
				/* lookup in dictionary, if found then get the global
				 * specification. Then merge both local and global
				 * specifications into a final one.
				 */
				info2 := Specification(declaration + " " + info.Name)
				info = Merge(info2, info) /* info2 provides the base, whilst info will override with inline parts */
			}

			/* DEBUG: use for debug and development purposes only */
			if info.Type == "" && ctx.config.ReadAhead { /* read the actual value and get the type information */
				value := values[i]
				info3 := Specification(value.Type() + " " + info.Name)
				info = Merge(info, info3)
			}

			tokens2 = append(tokens2, RtToken(info.String()))

		} else { /* in error */
			if err := ctx.HandleError(Errorf(rie.BadToken, rie.Error, "bad parameter, expecting token got value instead -- %s", param)); err != nil {
				log.Fatal(err)
			}
			return
		}
	}


	switch string(name) {
		case "AttributeBegin", "FrameBegin", "MotionBegin", "ObjectBegin", "SolidBegin", "TransformBegin", "WorldBegin","IfBegin":

			ctx.statistics.Begin(string(name))
			
		break
		case "AttributeEnd", "FrameEnd", "MotionEnd", "ObjectEnd", "SolidEnd", "TransformEnd", "WorldEnd","IfEnd":
		
			/* check if the End and Begin match */
			if ctx.config.Strict {
				block := strings.TrimSuffix(ctx.statistics.CurrentBlock(),"Begin")
				if block != strings.TrimSuffix(string(name),"End") {
					if err := ctx.HandleError(Errorf(rie.BadToken,rie.Error,"block mismatch, expecting %s, but got %s instead",block + "End", string(name))); err != nil {
						log.Fatal(err)
					}
				}
			}
			ctx.statistics.End()
	
		break
		default:
			/* only count non-structural commands */
			ctx.statistics.IncrementStatements()
			ctx.statistics.IncrementTokens(len(params),inlineparams)
		break
	}

	if len(annotations) > 0 {
		aparams,avalues := Unmix(annotations)
		/* check for annotation, and then apply to the current block, TODO: check the balance of aparams = avalues  */
		for i,anno := range aparams {
			if token,ok := anno.(RtToken); ok {
				info := Specification(string(token))
				value := avalues[i]

				switch info.Name {
					case "label":
						ctx.statistics.AnnotateLabel(token,value)
					break
					case "asset":
						ctx.statistics.AddAsset(token,value)
					break
				}
			}
		}
	}
				


	
	/* go through all and check the tokens against the values */
	if ctx.config.Strict {
		for i, token := range tokens2 {
			value := values[i]
			info := Specification(string(token))

			if info.Type == "" { /* nothing set */
				if err := ctx.HandleError(Error(rie.BadToken, rie.Error, "bad parameter|token, unknown value type")); err != nil {
					if ctx.config.Debug {
						fmt.Printf("!! info = %s\n!! value = %s %v\n", info.String2(), value.Type(), value)
					}
					log.Fatal(err)
				}
			}

			if value.Type() != info.LongType() && info.Class != "vertex" {
				/* FIXME: */
				if err := ctx.HandleError(Errorf(rie.Consistency, rie.Error, "bad parameter, expecting %s, but was %s", info.LongType(), value.Type())); err != nil {
					if ctx.config.Debug {
						fmt.Printf("!! info = %s\n!! value = %s %v\n", info.String2(), value.Type(), value)
					}
					log.Fatal(err)
				}
			}
		}
	} /* /end strict */

	if ctx.open && name == RtString("Begin") {
		if err := ctx.HandleError(Error(rie.NotStarted, rie.Error, "already begun")); err != nil {
			ctx.fatal(err)
		}
	}

	/* TODO: everything below here should be in the Driver code,
		 * concept: a prototype-driver sits in the drivers place, once begin has been called
	   *          the actual driver takes over. This allows a very specific error handler to
	   *          exist, reducing the more complex code below
	*/

	/* here we only handle the prototype driver code which consists of begin */

	switch string(name) {
	case "Begin", "begin":
		ctx.open = true
		ctx.dict = make(map[string]string, 0)
		ctx.seen = false
		ctx.driver = nil

		statement, ok := args[0].(RtToken)
		if !ok {
			if err := ctx.HandleError(Errorf(rie.System, rie.Error, "expecting token, got %s instead", args[0].Type())); err != nil {
				ctx.fatal(err)
			}
		}

		options := &DriverOptions{}
		options.PrettyPrint = ctx.config.PrettyPrint

		if string(statement) != "|" {
			sargs := make([]string, 0)
			sargs = append(sargs, string(statement))
			for _, arg := range args {
				if str, ok := arg.(RtToken); ok {
					sargs = append(sargs, string(str))
				}
			}

			d, err := BuildRIBFileDriver(ctx.logger, options, sargs...)
			if err != nil {
				ctx.fatal(err)
			}

			ctx.driver = d
			return
		}

		/* start a pipe */
		target, ok := args[1].(RtToken)
		if !ok {
			if err := ctx.HandleError(Errorf(rie.System, rie.Error, "expecting token, got %s instead", args[1].Type())); err != nil {
				ctx.fatal(err)
			}
		}

		switch string(target) {
		case "stdout": /* Standard Out */

			sargs := make([]string, 0)
			for i, arg := range args {
				if i > 1 {

					if str, ok := arg.(RtToken); ok {
						sargs = append(sargs, string(str))
					}
				}
			}

			d, err := BuildRIBStdoutDriver(ctx.logger, options, sargs...)
			if err != nil {
				ctx.fatal(err)
			}

			ctx.driver = d

			break
		case "render": /* Render */

			sargs := make([]string, 0)
			for i, arg := range args {
				if i > 1 {
					if str, ok := arg.(RtToken); ok {
						sargs = append(sargs, string(str))
					}
				}
			}

			d, err := BuildRenderDriver(ctx.logger, options, sargs...)
			if err != nil {
				ctx.fatal(err)
			}

			ctx.driver = d

			break
		case "catrib": /* Catrib */

			sargs := make([]string, 0)
			for i, arg := range args {
				if i > 1 {
					if str, ok := arg.(RtToken); ok {
						sargs = append(sargs, string(str))
					}
				}
			}

			d, err := BuildCatribDriver(ctx.logger, options, sargs...)
			if err != nil {
				ctx.fatal(err)
			}

			ctx.driver = d
			break
		case "debug": /* Debug */

			sargs := make([]string, 0)
			for i, arg := range args {
				if i > 1 {
					if str, ok := arg.(RtToken); ok {
						sargs = append(sargs, string(str))
					}
				}
			}

			d, err := BuildDebugDriver(ctx.logger, options, sargs...)
			if err != nil {
				ctx.fatal(err)
			}

			ctx.driver = d
			break
		case "block": /* Block Diagramming */

			sargs := make([]string,0)
			for i,arg := range args {
				if i > 1 {
					if str, ok := arg.(RtToken); ok {
						sargs = append(sargs,string(str))
					}
				}
			}

			d, err := BuildBlockDiagrammingDriver(ctx.logger,options,sargs...)
			if err != nil {
				ctx.fatal(err)
			}

			ctx.driver = d
		break
	}
	break /* end of Begin */
	case "End", "end":
		if ctx.open {
			ctx.open = false

			if ctx.driver != nil {
				ctx.driver.Close()
			}

			return
		} else {
			if err := ctx.HandleError(Error(rie.NotStarted, rie.Error, "not begun yet")); err != nil {
				ctx.fatal(err)
			}
		}
	default:
		if !ctx.open {
			if err := ctx.HandleError(Error(rie.NotStarted, rie.Error, "must call begin first")); err != nil {
				ctx.fatal(err)
			}
			return
		}

		/* RIB output --> pass down data sequence */
		if !ctx.seen {
			/* check if the current data is RenderMan RIB, if not then inject before */
			if string(name) == "##" {
				if len(args) > 0 {
					if str, ok := args[0].(RtString); ok {
						if strings.Contains(string(str), "RenderMan") {
							ctx.seen = true
						}
					}
				}
			}

			if !ctx.seen {
				/* inject into pipe ##RenderMan RIB */
				if !ctx.config.Fragment {
					/* if we are a fragment then ignore */
					ctx.Handle(List("##", []RtPointer{RtString("RenderMan RIB")}, nil))
				}
				ctx.seen = true
			}
		}

		if ctx.driver != nil {
			ctx.driver.Handle(name, args, tokens, values)
		}

		break
	}
}

func (ctx *Context) GetProgress() RtInt {
	ctx.RLock()
	defer ctx.RUnlock()
	if ctx.driver != nil {
		return ctx.driver.GetProgress()
	}
	return RtInt(100)
}

func (ctx *Context) GetLastRIB() string {
	ctx.RLock()
	defer ctx.RUnlock()
	if ctx.driver != nil {
		return ctx.driver.GetLastRIB()
	}
	return ""
}

func New(config *Configuration) *RiContext {

	return Wrap(NewContext(config))
}
