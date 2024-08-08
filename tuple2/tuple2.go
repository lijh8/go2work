package tuple2

import (
	"math"
)

func Cmp(a, b []any) (int, bool) {
	for i := 0; i != min(len(a), len(b)); i++ {
		if a[i] == nil && b[i] != nil {
			return -1, true
		} else if a[i] != nil && b[i] == nil {
			return 1, true
		}

		a_bool, a_bool_ok := a[i].(bool)
		a_string, a_string_ok := a[i].(string)
		a_int, a_int_ok := a[i].(int)
		a_float64, a_float64_ok := a[i].(float64)

		b_bool, b_bool_ok := b[i].(bool)
		b_string, b_string_ok := b[i].(string)
		b_int, b_int_ok := b[i].(int)
		b_float64, b_float64_ok := b[i].(float64)

		if a_bool_ok != b_bool_ok ||
			a_string_ok != b_string_ok ||
			a_int_ok != b_int_ok ||
			a_float64_ok != b_float64_ok {

			return math.MaxInt, false
		}

		if a_bool_ok {
			if !a_bool && b_bool {
				return -1, true
			} else if a_bool && !b_bool {
				return 1, true
			}
		} else if a_string_ok {
			if a_string < b_string {
				return -1, true
			} else if a_string > b_string {
				return 1, true
			}
		} else if a_int_ok {
			if a_int < b_int {
				return -1, true
			} else if a_int > b_int {
				return 1, true
			}
		} else if a_float64_ok {
			if a_float64 < b_float64 {
				return -1, true
			} else if a_float64 > b_float64 {
				return 1, true
			}
		}
	}

	if len(a) < len(b) {
		return -1, true
	} else if len(a) > len(b) {
		return 1, true
	}

	return 0, true
}

/*
func main() {
	s1 := []any{1, "hello", 3.14, true}
	s2 := []any{1, "hello", 3.14, true}
	b, ok := tuple2.Cmp(s1, s2)
	println(b, ok)
}
*/
