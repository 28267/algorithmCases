package util

func F1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func F2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func F3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func F4() (r int) {
	defer func(r int) {
		r++
	}(r)
	return 1
}

func IncreaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func IncreaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}
