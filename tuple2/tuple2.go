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

		aType := reflect.TypeOf(a[i])
		bType := reflect.TypeOf(b[i])
		if aType != bType {
			return 0, false
		}

		aString, aStringOk := a[i].(string)
		aInt, aIntOk := a[i].(int)
		aFloat64, aFloat64Ok := a[i].(float64)

		bString, bStringOk := b[i].(string)
		bInt, bIntOk := b[i].(int)
		bFloat64, bFloat64Ok := b[i].(float64)

		if aStringOk && bStringOk {
			if c := cmp.Compare(aString, bString); c != 0 {
				return c, true
			}
		}

		if aIntOk && bIntOk {
			if c := cmp.Compare(aInt, bInt); c != 0 {
				return c, true
			}
		}

		if aFloat64Ok && bFloat64Ok {
			if c := cmp.Compare(aFloat64, bFloat64); c != 0 {
				return c, true
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
