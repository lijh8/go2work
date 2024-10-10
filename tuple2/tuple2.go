package tuple2

import (
	"cmp"
	"reflect"
)

// two tuples should be the same length and
// element types are same for same position.
func Compare(a, b []any) (int, bool) {
	if len(a) != len(b) {
		return 0, false
	}

	for i := range a {
		if a[i] == nil || b[i] == nil {
			return 0, false
		}

		if _, boolean := a[i].(bool); boolean {
			return 0, false
		}
		if _, boolean := b[i].(bool); boolean {
			return 0, false
		}

		if a, b := reflect.TypeOf(a[i]), reflect.TypeOf(b[i]); a != b {
			return 0, false
		}

		if a, aOk := a[i].(string); aOk {
			if b, bOk := b[i].(string); bOk {
				if c := cmp.Compare(a, b); c != 0 {
					return c, true
				}
			}
		}

		if a, aOk := a[i].(int); aOk {
			if b, bOk := b[i].(int); bOk {
				if c := cmp.Compare(a, b); c != 0 {
					return c, true
				}
			}
		}

		if a, aOk := a[i].(float64); aOk {
			if b, bOk := b[i].(float64); bOk {
				if c := cmp.Compare(a, b); c != 0 {
					return c, true
				}
			}
		}

		if a, aOk := a[i].([]any); aOk {
			if b, bOk := b[i].([]any); bOk {
				if c, ok := Compare(a, b); ok && c != 0 {
					return c, true
				} else if !ok {
					return 0, false
				}
			}
		}
	}
	return 0, true
}

/*
import (
	"tuple2"
)

func main() {
	a := []any{"abc", 123, 3.14}
	b := []any{"abc", 123, 3.14}
	c, ok := tuple2.Compare(a, b)
	fmt.Println(c, ok)
}
*/
