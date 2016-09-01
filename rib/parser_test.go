/* rigo2/rib/parser_test.go */
package rib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"fmt"
	"io"
	"strings"
)

type TestTokenWriter struct {
	tokens   []Token
	position int
}

func (w *TestTokenWriter) Write(t Token) {
	if w.tokens == nil {
		w.tokens = make([]Token, 0)
	}
	w.tokens = append(w.tokens, t)
}

func (w *TestTokenWriter) Read() (Token, error) {
	if w.tokens == nil || len(w.tokens) == 0 {
		return EmptyToken, io.EOF
	}
	if w.position >= len(w.tokens) {
		return EmptyToken, io.EOF
	}
	t := w.tokens[w.position]
	w.position++
	return t, nil
}

func (w *TestTokenWriter) Count() int {
	if w.tokens == nil {
		return 0
	}
	return len(w.tokens)
}

func (w *TestTokenWriter) Print(show bool) string {
	out := ""
	for _, token := range w.tokens {
		if !show {
			tag := "content"
			if token.Type == Tokeniser {
				tag = "tokeniser"
			}

			ritype := token.RiType
			if token.Error != nil {
				ritype = token.Error.Error()
			}
			out += fmt.Sprintf("%04d:%03d --%30s\t(%s)\tL:%10s\tRi:%10s\n",
				token.Line, token.Pos, token.Word, tag, token.Lex, ritype)
			continue
		}

		if token.Type == Tokeniser {
			continue
		}

		out += fmt.Sprintf("%04d:%03d --%30s\n", token.Line, token.Pos, token.Word)
	}
	return out
}

func Test_TokeniserExample0(t *testing.T) {

	Convey("Tokeniser Example 0", t, func() {
		tw := new(TestTokenWriter)
		err := Tokenise(strings.NewReader(RIBExample0), tw)
		So(err, ShouldBeNil)

		fmt.Printf("\nRIB Example 0\n----------\n%s\n\n%s\n", RIBExample0, tw.Print(false))

		tw1 := new(TestTokenWriter)
		err = Lexer(tw, tw1, PorousBloomFilter{})
		So(err, ShouldBeNil)

		fmt.Printf("\nLexer\n%s\n", tw1.Print(false))

		tw2 := new(TestTokenWriter)
		err = Parser(tw1, tw2)
		So(err, ShouldBeNil)

		fmt.Printf("\nParser\n%s\n", tw2.Print(false))

	})
}

func Test_TokeniserExample1(t *testing.T) {

	//t.Skip()

	Convey("Tokeniser Example 1", t, func() {

		tw := new(TestTokenWriter)
		err := Tokenise(strings.NewReader(RIBExample1), tw)
		So(err, ShouldBeNil)
		//	So(tw.Count(),ShouldEqual,27)

		fmt.Printf("\nRIB Example 1\n----------\n%s\n\n%s\n", RIBExample1, tw.Print(false))

		tw1 := new(TestTokenWriter)
		err = Lexer(tw, tw1, PorousBloomFilter{})
		So(err, ShouldBeNil)

		fmt.Printf("\nLexer\n%s\n", tw1.Print(false))
	})
}

const RIBExample0 = `##RenderMan RIB-Structure 1.1
version 3.04
Projection "perspective" "fov" 30.0
Color [1 0 0]
Sphere 1 -1 1 360
`

const RIBExample1 = `##RenderMan RIB-Structure 1.1
version 3.04
Display "sphere.tif" "file" "rgb"
Format 320 240 1
Translate 0 0 6
WorldBegin
Projection "perspective" "float fov" 30.0
Color [1 0 0]
Sphere 1 -1 1 360
WorldEnd`


