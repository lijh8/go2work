package tuple2

import (
	"cmp"
	"reflect"
)

func Cmp(a, b []any) (int, bool) {
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
	}
	return 0, true
}

/*
func main() {
	a := []any{"abc", 123, 3.14}
	b := []any{"abc", 123, 3.14}
	c, ok := tuple2.Cmp(a, b)
	fmt.Println(c, ok)
}
*/
