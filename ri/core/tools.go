/* rigo/core/tools.go */
package core

import (
	"fmt"
)

/* Format or reduce RtFloat to a nice
 * compact repesentation for RIB
 */
func Reduce(f RtFloat) string {

	str := fmt.Sprintf("%f", f)
	s := 0
	neg := false
	for i, c := range str {
		if c != '0' {
			if c == '-' {
				neg = true
				continue
			}
			s = i
			break
		}
		if c == '.' {
			break
		}
	}

	e := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != '0' {
			e = i + 1
			break
		}
		if str[i] == '.' {
			break
		}
	}

	str = str[s:e]
	if str == "." {
		str = "0"
	}
	if neg {
		str = "-" + str
	}
	if str[len(str)-1] == '.' {
		str = str[:len(str)-1]
	}

	return str
}

func Reducev(fv []RtFloat) string {
	out := ""
	for i, f := range fv {
		out += Reduce(f)
		if i < len(fv)-1 {
			out += " "
		}
	}
	return out
}

func Serialise(list []RtPointer, parameter bool) string {
	/* if parameter then incase non-arrays with "[" && "]" */
	out := ""
	for i, param := range list {
		if i > 0 {
			out += " "
		}

		/* if serialising a parameterlist (token + value, ...) then
		 * only bracket the value, not the token.
		 */
		if parameter && i%2 != 0 {

			str := param.String()

			if len(str) > 0 {
				if str[0] != '[' && str[len(str)-1] != ']' {
					str = "[" + str + "]"
				}
			}

			out += str
			continue
		}

		out += param.String()
	}
	return out
}

func Unmix(list []RtPointer) ([]RtPointer, []RtPointer) {
	params := make([]RtPointer, 0)
	values := make([]RtPointer, 0)

	for i, param := range list {
		if i%2 == 0 {
			params = append(params, param)
		} else {
			values = append(values, param)
		}
	}
	return params, values
}

func Mix(tokens []RtPointer, values []RtPointer) []RtPointer {
	params := make([]RtPointer, 0)

	for i, token := range tokens {
		params = append(params, token)
		if i >= len(values) {
			break
		}
		params = append(params, values[i])
	}
	return params
}
