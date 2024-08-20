package main

/*
func main2() {
	a := 10
	func() {
		a++
	}()
	println(a) // 11

}
*/

func main() {

	values := []int{1, 2, 3}

	// problem
	for _, val := range values {
		func() {
			println(val)
		}()
	}

	// temporary fix
	// for _, val := range values {
	// 	func(val int) { // val, not anymore. Fixed
	// 		println(val)
	// 	}(val) // val, not anymore. Fixed
	// }
}
