package tuple2

import (
	"errors"
	"math"
)

func Cmp(a, b []any) (int, error) {
	for i := 0; i != min(len(a), len(b)); i++ {
		if a[i] == nil && b[i] != nil {
			return -1, nil
		} else if a[i] != nil && b[i] == nil {
			return 1, nil
		}

		a_bool_v, a_bool_ok := a[i].(bool)
		a_string_v, a_string_ok := a[i].(string)
		a_int_v, a_int_ok := a[i].(int)
		a_float64_v, a_float64_ok := a[i].(float64)

		b_bool_v, b_bool_ok := b[i].(bool)
		b_string_v, b_string_ok := b[i].(string)
		b_int_v, b_int_ok := b[i].(int)
		b_float64_v, b_float64_ok := b[i].(float64)

		if a_bool_ok != b_bool_ok ||
			a_string_ok != b_string_ok ||
			a_int_ok != b_int_ok ||
			a_float64_ok != b_float64_ok {
			return math.MaxInt, errors.New("type mismatch of two elements")
		}

		if a_bool_ok {
			if !a_bool_v && b_bool_v {
				return -1, nil
			} else if a_bool_v && !b_bool_v {
				return 1, nil
			}
		} else if a_string_ok {
			if a_string_v < b_string_v {
				return -1, nil
			} else if a_string_v > b_string_v {
				return 1, nil
			}
		} else if a_int_ok {
			if a_int_v < b_int_v {
				return -1, nil
			} else if a_int_v > b_int_v {
				return 1, nil
			}
		} else if a_float64_ok {
			if a_float64_v < b_float64_v {
				return -1, nil
			} else if a_float64_v > b_float64_v {
				return 1, nil
			}
		}
	}
	if len(a) < len(b) {
		return -1, nil
	} else if len(a) > len(b) {
		return 1, nil
	}
	return 0, nil
}

/*
func main() {
	s1 := []any{1, "hello", 3.14, true}
	s2 := []any{1, "hello", 3.14, true}
	b, err := TupleCmp(s1, s2)
	println(b, err)
}
*/
