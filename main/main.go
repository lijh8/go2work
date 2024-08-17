package main

func main() {
	{
		a := make([]int, 3)
		a = append(a, 1, 2, 3)
		println(a)
	}
	{
		a := []int{1, 2, 3}
		println(a)
	}
	{
		a := make(map[string]int)
		a["aaa"] = 100
	}
	{
		a := map[string]int{"aaa": 100}
		println(a)
	}

	type T struct {
		name string
		num  int
	}

	{
		a := new(T)
		a.name = "aaa"
		a.num = 100
		println(a)
	}
	{
		a := &T{name: "aaa", num: 100}
		a.name = "bbb"
		println(a)
	}
}
