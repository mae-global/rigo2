package ri

import (
	"io"
	"fmt"
	"strconv"

	. "github.com/mae-global/rigo2/ri/core"

	"github.com/mae-global/rigo2/rib"
	"github.com/mae-global/rigo2/ri/prototype"
)

type tokenio struct {
	tokens []rib.Token
	position int
}

func (w *tokenio) Write(t rib.Token) {
	if w.tokens == nil {
		w.tokens = make([]rib.Token, 0)
	}
	w.tokens = append(w.tokens, t)
}

func (w *tokenio) Read() (rib.Token, error) {
	if w.tokens == nil || len(w.tokens) == 0 {
		return rib.EmptyToken, io.EOF
	}

	if w.position >= len(w.tokens) {
		return rib.EmptyToken, io.EOF
	}
	t := w.tokens[w.position]
	w.position++
	return t, nil
}

func (w *tokenio) Print() string {
	out := ""
	for _, token := range w.tokens {
		tag := "content"
		if token.Type == rib.Tokeniser {
			tag = "tokeniser"
		}

		ritype := token.RiType
		if token.Error != nil {
			ritype = token.Error.Error()
		}
		out += fmt.Sprintf("%04d:%03d --%30s\t(%s)\tL:%10s\tRi:%10s\n",
			token.Line, token.Pos, token.Word, tag, token.Lex, ritype)
	}

	return out
}



func parse(reader io.Reader, writer RtContextHandle) error {
	
	tw := new(tokenio)
	if err := rib.Tokenise(reader, tw); err != nil {
		return err
	}

	tw1 := new(tokenio)
	if err := rib.Lexer(tw, tw1, prototype.RiBloomFilter()); err != nil {
		return err
	}

	tw2 := new(tokenio)
	if err := rib.Parser(tw1, tw2); err != nil {
		return err
	}

	lookup := prototype.RiPrototypes()

	fmt.Printf("\n%s\n\n", tw2.Print())

	args := make([]RtPointer, 0)
	tokens := make([]RtPointer, 0)
	values := make([]RtPointer, 0)
		

	farray := make([]RtFloat, 0)
	isarray := false
	
	var proto *prototype.Information

	for {
		token, err := tw2.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if token.Type != rib.Content {
			continue
		}

		switch token.RiType {
		case "func":
			if proto != nil { /* write to out with the current */

				/* FIXME: remove the debug printing */
				fmt.Printf("%s (%d) %d args, %d tokens & %d values\n",
										proto.Name, len(proto.Arguments), len(args), len(tokens), len(values))
				fmt.Printf("tokens = %v, values = %v\n",tokens,values)


				/* Due to the dumbness of the parser we now correct the parser with the prototype information */			
				nargs,ntokens,nvalues,err := CorrectParser(proto,args,tokens,values)

				if err != nil {
					return err
				}
				
				/* Write out */
				if writer != nil {

					if len(ntokens) != len(nvalues) {
						return fmt.Errorf("Tokens - Values mismatch, %d tokens != %d values",len(ntokens),len(nvalues))
					}

					writer.HandleV(proto.Name, nargs, ntokens, nvalues)					
				}
		

				/* Reset ready for next function */
				args = make([]RtPointer, 0)
				tokens = make([]RtPointer, 0)
				values = make([]RtPointer, 0)
				farray = make([]RtFloat, 0)
			}

			/* start new func, lookup the information required */
			if p, ok := lookup[RtString(token.Word)]; ok {
					proto = p						
			} else {
				return fmt.Errorf("unknown function [\"%s\"]",token.Word)
			}
			break

		case "float":
			if f, err := strconv.ParseFloat(token.Word, 64); err != nil {
				return err
			} else {
				farray = append(farray, RtFloat(f))
				if !isarray {
					if len(args) >= len(proto.Arguments) {
						values = append(values, RtFloat(f)) //RtFloatArray(farray))
					} else {
						args = append(args, RtFloat(f)) //RtFloatArray(farray))
					}
				}
			}
			break
		case "string":
			if len(args) >= len(proto.Arguments) {
				if len(tokens) % 2 != 0 {
					values = append(values,RtString(token.Word))
				} else {
					tokens = append(tokens,RtToken(token.Word))	
				}
			} else {
				args = append(args,RtString(token.Word))
			}
		break
		case "token":
			if len(args) >= len(proto.Arguments) {
				if len(tokens) % 2 != 0 {
					values = append(values,RtToken(token.Word))
				} else {
					tokens = append(tokens, RtToken(token.Word))
				}
			} else {
				args = append(args, RtString(token.Word))
			}
			break

		case "array_begin":
			farray = make([]RtFloat, 0)
			isarray = true
			break

		case "array_end":
			isarray = false
			if len(args) >= len(proto.Arguments) {
				values = append(values, RtFloatArray(farray))
			} else {
				args = append(args, RtFloatArray(farray))
			}
			break
		}
	}
	/* tail */
	if proto != nil { /* write to out with the current */

		/* FIXME: remove the debug printing */
		fmt.Printf("T %s (%d) %d args, %d tokens & %d values\n",
								proto.Name, len(proto.Arguments), len(args), len(tokens), len(values))

		/* Due to the dumbness of the parser we now correct the parser with the prototype information */			
		nargs,ntokens,nvalues,err := CorrectParser(proto,args,tokens,values)
		if err != nil {
			return err
		}
				
		/* Write out */
		if writer != nil {
			writer.HandleV(proto.Name, nargs, ntokens, nvalues)
		}
	}

	return nil
}




func CorrectParser(proto *prototype.Information, args, tokens, values []RtPointer) ([]RtPointer,[]RtPointer,[]RtPointer,error) {

	if len(args) != len(proto.Arguments) {
		fmt.Printf("%s\n",proto)
		return nil,nil,nil,fmt.Errorf("Invalid count of arguments (%d != %d) for \"%s\"",len(proto.Arguments),len(args),proto.Name)
	}

	if len(tokens) > 0 || len(values) > 0 {
		if !proto.Parameterlist {
			fmt.Printf("%s\n",proto)
			return nil,nil,nil,fmt.Errorf("\"%s\" has no parameterlist list",proto.Name)
		}
	}

	nargs := make([]RtPointer,0)

	/* go through the arguments first and correct types as needed */	
	for i := 0; i < len(proto.Arguments); i++ {

		arg0 := proto.Arguments[i]
		arg1 := args[i]

		/* types are decomposed into these basic types; RtToken, RtFloat and RtFloatArray */
		switch arg1.Type() {
			case "token": 
				v := arg1.(RtToken)

				switch arg0.Type {
				case "token":
					nargs = append(nargs,v)	
				break
				case "string":
					nargs = append(nargs,RtString(string(v)))
				break
				case "lighthandle":
					nargs = append(nargs,RtLightHandle(string(v)))
				break			

	
				default:
					return nil,nil,nil,fmt.Errorf("Invalid Type -- \"%s\" (%s), should be \"%s\"",arg1.Type(),arg1,arg0.Type)
				break
			}
			break
			case "float":
				v := arg1.(RtFloat)
				
				switch arg0.Type {
					case "float":
						nargs = append(nargs,v)
					break					
					case "int":
						nargs = append(nargs,v)
					break
					/* TODO add rest here */
					default:
						return nil,nil,nil,fmt.Errorf("Invalid Type -- \"%s\" (%s), should be \"%s\"",arg1.Type(),arg1,arg0.Type)
					break
				}
			break
			case "float[]":
				v := arg1.(RtFloatArray)

				switch arg0.Type {
				case "float":
					/* check that v is singular then add -- otherwise in error */
					if len(v) != 1 {
						fmt.Printf("v = %v\n",v)
						fmt.Printf("args = %v\n",args)
						return nil,nil,nil,fmt.Errorf("Invalid Type %s(%s) -- expecting singular float but have %d floats",
																					proto.Name,arg0.Name,len(v))
					}
					
					nargs = append(nargs,RtFloat(v[0]))
				break				
				case "float[]":
					nargs = append(nargs,v)
				break
			
				/* TODO add rest here */
				default:
					return nil,nil,nil,fmt.Errorf("Invalid Type -- \"%s\" (%s), should be \"%s\"",arg1.Type(),arg1,arg0.Type)
				break
			}
			break
		}
		
	}	
	/* go through the parameterlist list is present and attempt to correct the types */


	return nargs,tokens,values,nil
}







