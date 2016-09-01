package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
	"text/template"

	. "../"
)


const Version = 0

func main() {

	target := flag.String("target","/opt/pixar/RenderManProServer-20.8/include/ri.h","ri.h to generate from")
	pname  := flag.String("package","prototype","package name")

	flag.Parse()

	fmt.Printf("Generator Version %d\n",Version)
	fmt.Printf("Parsing [%s]...\n",*target)

	file,err := ioutil.ReadFile(*target)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error reading [%s] -- %v\n",*target,err)
		os.Exit(1)
	}

	buffer := strings.Split(string(file),"\n")
	ri := new(Ri)

	pos := 0

	for i, line := range buffer {
		if strings.Contains(line,"#define RI_VERSION") {
			parts := strings.Split(line," ")
			if len(parts) == 3 {
				ver,err := strconv.Atoi(parts[2])
				if err != nil {
					fmt.Fprintf(os.Stderr,"Error parsing RI_VERSION -- %v\n",err)
					os.Exit(1)
				}
				ri.Version = ver
				fmt.Printf("[%05d] RI_VERSION = %d\n",i,ri.Version)
			}
			continue
		}

		if strings.Contains(line,"RI_EXPORT RtToken") {
			parts := strings.Split(line," ")
			if len(parts) == 3 {
				token := strings.TrimPrefix(strings.Replace(parts[2],";","",1),"RI_")
				ri.Tokens = append(ri.Tokens,token)
				if len(token) > ri.Longest {
					ri.Longest = len(token) /* FIXME: is this still used? */
				}
			}
		}

		if strings.Contains(line,"RI_EXPORT") {
			parts := strings.Split(line," ")
			var name string
			for _,p := range parts {
				if strings.Contains(p,"Ri") {
					name = p
					break
				}
			}
			if name == "" {
				continue
			}

			nameparts := strings.Split(name,"(")
			name = strings.TrimPrefix(nameparts[0],"Ri")
			name = strings.Replace(name,";","",-1)

			if name == "ArchiveBegin" { /* skip until we find the actual Ri C API */
			
				pos = i 
				/* rewind */
				break
			}
		}
	} /* end of range buffer */

	fmt.Printf("Ri C prototype interface found from line %d\n",pos)

	hashes := make(map[string][]int,0)
	
	hashes["version"] = Hash("version")

	/* first fix the neat freak */
	buffer = buffer[pos:]
	var buf string	
	statements := make([]string,0)
	statements = append(statements,"version float")

	for _,line := range buffer {
		buf += strings.Replace(line,"\n","",-1)
	}

	buffer = strings.Split(buf,";")

	
	for i,line := range buffer {
		if strings.Contains(line,"Ri") {
			prototypes := ParseCPrototype(line)
			if len(prototypes) == 0 {
				continue
			}
			
			name := prototypes[0]
			if name[(len(name) - 1)] == 'V' || name == "GetContext" || name == "Context" || name == "ErrorHandler" {
				continue
			}
			fmt.Printf("[%5d] -- \"%s\"\n",i,strings.Join(prototypes," "))
			statements = append(statements,strings.Join(prototypes," "))
			hashes[name] = Hash(name)			
		}
	}
	


	fmt.Printf("\n\n")

	size := 128
	step := size
	var bloom []int
	var filter *BloomFilter

	for {
		bloom = make([]int,size)
		
		for _, h := range hashes {
			for i := 0; i < len(h); i++ {
				bloom[ h[i] % size ]++
			}
		}

		//filter = &BloomFilter{bloom,len(hashes)}
		filter = SetBloomFilter(bloom,len(hashes))
		fmt.Printf("\n============= [%05d] BloomFilter\n%s\n",size,filter.Print())

		_,_, sparse := filter.Stats()

		if !filter.IsMember("Alice","Fred","Eve") && filter.IsMember("Sphere","Translate") {
			if sparse >= 0.8 {
				break
			}
		} else {
			fmt.Printf(">>> failed membership test\n")
		}

		size = size + step
	}
	
	keys,bits := filter.Raw()
	ri.FilterData = bits
	ri.FilterKeys = keys
	
	/* buffer all c prototyped strings into an indexed string */
	indices := make([]int,0)
	var prototypes string
	index := 0	
	for _,statement := range statements {
		prototypes += statement 
		indices = append(indices,index)
		index = len(prototypes) 
	}

	fmt.Printf("%d prototypes = %dbytes, %d indices\n",len(statements),len(prototypes),len(indices))
	fmt.Printf("%s\n\n",prototypes)
	
	for i,_ := range indices {
		if (i + 1) >= len(indices) {
			idx := indices[i]
			fmt.Printf("[%s]\n",prototypes[ idx : ])
		} else {	
			idx := indices[i]
			fmt.Printf("[%s]\n",prototypes[ idx : indices[i + 1] ])
		}
	}

	ri.Prototypes = prototypes
	ri.PrototypesIndices = indices

	o := &T{}
	o.Version = Version
	o.Date = time.Now().UTC()
	o.Source = *target
	o.Name = *pname
	o.Ri = ri

	t,err := template.New("temp").Funcs(template.FuncMap{"Token":Token,"Value":Value}).Parse(TokensTemplate)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error parsing template -- %v\n",err)
		os.Exit(1)
	}

	b := bytes.NewBuffer(nil)

	if err := t.Execute(b,o); err != nil {
		fmt.Fprintf(os.Stderr,"Error executing template -- %v\n",err)
		os.Exit(1)
	}
	
	if err := ioutil.WriteFile("out.txt",b.Bytes(),0664); err != nil {
		fmt.Fprintf(os.Stderr,"Error writing file -- %v\n",err)
		os.Exit(1)
	}

	t,err = template.New("temp").Funcs(template.FuncMap{"Token":Token,"Value":Value}).Parse(PrototypesTemplate)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error parsing template -- %v\n",err)
		os.Exit(1)
	}

	b = bytes.NewBuffer(nil)

	if err := t.Execute(b,o); err != nil {
		fmt.Fprintf(os.Stderr,"Error executing template -- %v\n",err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile("out1.txt",b.Bytes(),0664); err != nil {
		fmt.Fprintf(os.Stderr,"Error writing file -- %v\n",err)
		os.Exit(1)
	}

}

type Ri struct {
	Version int
	Tokens []string
	Longest int
	FilterData []int
	FilterKeys int
	Prototypes string
	PrototypesIndices []int
}

type T struct {
	Version int
	Date time.Time
	Source string
	Name string
	Ri *Ri
}

func Token(t string) string {
	return fmt.Sprintf("\t%s",t)
}

func Value(v string) string {
	return strings.ToLower(v)
}

const TokensTemplate = `/* machine generated
 * build tool version {{.Version}}
 * generated on {{.Date}}
 * source {{.Source}}
 * RiVersion {{.Ri.Version}}
 */

package {{.Name}}

const (
{{range $token := .Ri.Tokens}}{{$token | Token}} RtToken = "{{$token | Value}}"
{{end}}
)

/* EOF */
`

const PrototypesTemplate = `/* machine generated
 * build tool version {{.Version}}
 * generated on {{.Date}}
 * source {{.Source}}
 * RiVersion {{.Ri.Version}}
 */

package {{.Name}}

import (
	. "github.com/mae-global/rigo2/ri/core"
)

const (
	RIBVersion RtFloat = 3.04
	BloomFilterKeys int = {{.Ri.FilterKeys}}
)


var (
	BloomFilterData = []int{ {{range $i := .Ri.FilterData}}{{$i}},{{end}} }

	Data string = "{{.Ri.Prototypes}}"

	Indices = []int{ {{range $i := .Ri.PrototypesIndices}}{{$i}},{{end}} }
)


func RiBloomFilter() *BloomFilter {
	bits := make([]int,len(BloomFilterData))
	copy(bits,BloomFilterData)
	return &BloomFilter{bits,BloomFilterKeys}
}



func RiPrototypes() map[RtString]*Information {

	statements := make([]string,0)

	for i,_ := range Indices {
		if (i + 1) >= len(Indices) {
			idx := Indices[i]
			statements = append(statements,Data[ idx : ])
		} else {	
			idx := Indices[i]
			statements = append(statements,Data[ idx : Indices[i + 1] ])
		}
	}

	out := make(map[RtString]*Information,0)

	for _,statement := range statements {
		proto := Parse(statement)
		out[proto.Name] = proto
	}

	return out 
}	

/* EOF */
`


func ParseCPrototype(line string) []string {

	line = line[strings.Index(line,"Ri"):]
	
	prototype := make([]string,0) /* {name,"type(class) token"...} */
	var word string
	var curtype string /* current type */
	
	line = strings.Replace(line,"char*","RtString",-1)
	line = strings.Replace(line,"char *","RtString",-1)
	line = strings.Replace(line,"RtFloat*","RtFloat[]",-1)
	line = strings.Replace(line,"RtFloat *","RtFloat[]",-1)
	line = strings.Replace(line,"RtInt*","RtInt[]",-1)
	line = strings.Replace(line,"RtInt *","RtInt[]",-1)
	line = strings.Replace(line,"RtPoint*","RtPoint[]",-1)
	line = strings.Replace(line,"RtPoint *","RtPoint[]",-1)
	line = strings.Replace(line,"RtToken *","RtToken[]",-1)
	line = strings.Replace(line,"RtToken*","RtToken[]",-1)

	/* Drunken C Parsing */
	for _,c := range line {
		if c == ' ' || c == '(' || c == '\n' || c == '\t' || c == ',' || c == ')' {
			if len(word) > 0 {
				word = strings.TrimSpace(strings.TrimPrefix(word,"Ri"))
			
				if word == "..." {
					prototype = append(prototype,"...")
					word = ""
					continue
				}

				if strings.HasPrefix(word,"Rt") {
					curtype = strings.ToLower(strings.TrimPrefix(word,"Rt"))
					word = ""
					continue
				}
	
				toadd := word
				if len(curtype) > 0 {
					toadd = curtype + " " + toadd
				}

				if len(toadd) > 0 {	
					prototype = append(prototype,toadd)
				}

				word = ""
				curtype = ""
				continue
			}
		}
		word += string(c)
	}

	return prototype
}





