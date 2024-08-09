package tuple2

import "cmp"

func Cmp(a, b []any) (int, bool) {
	for i := 0; i != min(len(a), len(b)); i++ {
		if a[i] == nil || b[i] == nil {
			return 0, false
		}

		if _, boolean := a[i].(bool); boolean {
			return 0, false
		}
		if _, boolean := b[i].(bool); boolean {
			return 0, false
		}

		aString, aStringOk := a[i].(string)
		aInt, aIntOk := a[i].(int)
		aFloat64, aFloat64Ok := a[i].(float64)

		bString, bStringOk := b[i].(string)
		bInt, bIntOk := b[i].(int)
		bFloat64, bFloat64Ok := b[i].(float64)

		stringDiff := aStringOk != bStringOk
		intDiff := aIntOk != bIntOk
		float64Diff := aFloat64Ok != bFloat64Ok

		typeDiff := stringDiff || intDiff || float64Diff
		if typeDiff {
			return 0, false
		}

		if c := cmp.Compare(aString, bString); c != 0 {
			return c, true
		}
		if c := cmp.Compare(aInt, bInt); c != 0 {
			return c, true
		}
		if c := cmp.Compare(aFloat64, bFloat64); c != 0 {
			return c, true
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
	a := []any{"abc", 123, 3.14}
	b := []any{"abc", 123, 3.14}
	c, ok := tuple2.Cmp(a, b)
	fmt.Println(c, ok)
}
*/
