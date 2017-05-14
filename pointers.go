package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
}

func modify1(arr *[3]int) {
	(*arr)[0] = 90
}

func modify2(arr *[3]int) {
	arr[0] = 90
}

func modify3(sls []int) {
	sls[0] = 90
}

func main() {
	//Creating a pointer
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	//Zero value of a pointer is nil
	d := 25
	var p *int
	if p == nil {
		fmt.Println("\np is", p)
		p = &d
		fmt.Println("p after initialization is", p)
	}

	//dereferencing a pointer
	i := 255
	j := &i
	fmt.Println("\naddress of i is", j)
	fmt.Println("value of i is", *j)
	//changing the value pointed using dereference
	*j++
	fmt.Println("new value of i is", i)

	//passing pointer to a function
	k := 58
	fmt.Println("\nvalue of k before function call is", k)
	l := &k
	change(l)
	fmt.Println("value of k after function call is", k)

	//pointer to an array as a argument to a function
	m := [3]int{89, 90, 91}
	modify1(&m)
	fmt.Println("\nmodify1 using pointer to an array", m)

	n := [3]int{89, 90, 91}
	modify2(&n)
	fmt.Println("modify2 using pointer to an array", n)

	//slice as an argument to a function
	o := [3]int{89, 90, 91}
	modify3(o[:])
	fmt.Println("modify3 using slice", o)
}
