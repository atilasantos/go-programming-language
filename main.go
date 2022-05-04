package main

import "fmt"

func main() {
	a := make(map[int][]string)
	a[1] = []string{`John`, `Maria`, `Carlo`}
	a[0] = []string{`Johnny`, `Mariana`, `Carlos`}

	fmt.Print(a[0])
}
