package main

import "fmt"

func main() {
	a := []int{11, 22, 33, 44, 55}
	fmt.Println(a)
	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b)
	for i, v := range b {
		fmt.Println(i, v)
	}
	b = append(b, a[:]...)
	fmt.Println(b)
	b = append(b[:1], a...)
	fmt.Println(b)

	x := make([]int, 10, 100)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	jb := []string{"1", "pradip"}
	fmt.Println(jb)
	mb := []string{"2", "ram"}
	fmt.Println(mb)

	lb := [][]string{jb, mb}
	fmt.Println(lb)
	xb := [][][]string{lb, lb, lb}
	fmt.Println(xb)

	number := map[string]int{
		"pradip": 9874563210,
		"ram":    9630258741,
	}
	fmt.Println(number)
	fmt.Println(number["pradip"])

	println()

	v, ok := number["pradip"]
	fmt.Println(v, ok)

	_, ok1 := number["pradipa"]
	fmt.Println(ok1)

	if _, ok2 := number["pradip"]; ok2 {
		fmt.Println("Yehh")
	} else {
		fmt.Println("Not Yehh,")
	}

	number["shyam"] = 9512357046

	for key, val := range number {
		fmt.Println(key, val)
	}

	delete(number, "ram")
	fmt.Println(number)
}
