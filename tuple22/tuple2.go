package tuple2

func Cmp(a, b []any) (int, bool) {
	for i := 0; i != min(len(a), len(b)); i++ {
		if a[i] == nil && b[i] != nil {
			return -1, true
		}
		if a[i] != nil && b[i] == nil {
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

		one_side_bool := a_bool_ok != b_bool_ok
		one_side_string := a_string_ok != b_string_ok
		one_side_int := a_int_ok != b_int_ok
		one_side_float64 := a_float64_ok != b_float64_ok

		both_side_diff := one_side_bool || one_side_string || one_side_int || one_side_float64
		both_side_numeric := (a_int_ok || a_float64_ok) && (b_int_ok || b_float64_ok)

		if both_side_diff && !both_side_numeric {
			return 0, false
		}

		if both_side_numeric {
			if a_int_ok && b_int_ok {
				if a_int < b_int {
					return -1, true
				}
				if a_int > b_int {
					return 1, true
				}
			}
			if a_float64_ok && b_float64_ok {
				if a_float64 < b_float64 {
					return -1, true
				}
				if a_float64 > b_float64 {
					return 1, true
				}
			}
			if a_int_ok && b_float64_ok {
				if float64(a_int) < b_float64 {
					return -1, true
				}
				if float64(a_int) > b_float64 {
					return 1, true
				}
			}
			if a_float64_ok && b_int_ok {
				if a_float64 < float64(b_int) {
					return -1, true
				}
				if a_float64 > float64(b_int) {
					return 1, true
				}
			}
		}

		if a_bool_ok && b_bool_ok {
			if !a_bool && b_bool {
				return -1, true
			}
			if a_bool && !b_bool {
				return 1, true
			}
		}
		if a_string_ok && b_string_ok {
			if a_string < b_string {
				return -1, true
			}
			if a_string > b_string {
				return 1, true
			}
		}
	}

	if len(a) < len(b) {
		return -1, true
	}
	if len(a) > len(b) {
		return 1, true
	}

	return 0, true
}

/*
func main() {
	a := []any{1, "hello", 3.14, true}
	b := []any{1, "hello", 3.14, true}
	c, ok := tuple2.Cmp(a, b)
	println(c, ok)
}
*/
