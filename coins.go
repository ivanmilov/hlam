package main

import (
	"flag"
	"fmt"
	"time"
)

// buffer for known values
var s []int64

func main() {
	numbPtr := flag.Int("n", 4, "an int64")
	flag.Parse()
	var N = int64(*numbPtr)

	// for i := 1; i <= N; i++ {
	// 	fmt.Printf("%-2d\t%4d\n", i, f_recurs(i))
	// }
	t := time.Now()
	s = make([]int64, N+1)
	for i := int64(0); i <= N; i++ {
		s[i] = 0
	}

	fmt.Println(f_recurs(N))
	fmt.Println("Время поиска: ", time.Since(t))
	// fmt.Println(f_strait(N))
}

func f_recurs(n int64) int64 {
	if s[n] != 0 {
		return s[n]
	}

	switch {
	case n == 0:
		s[n] = 1
	case 1 <= n && n < 3:
		s[n] = f_recurs(n - 1)
	case 3 <= n && n < 5:
		s[n] = f_recurs(n-1) + f_recurs(n-3)
	case 5 <= n && n < 10:
		s[n] = f_recurs(n-1) + f_recurs(n-3) + f_recurs(n-5)
	case 10 <= n:
		s[n] = f_recurs(n-1) + f_recurs(n-3) + f_recurs(n-5) + f_recurs(n-10)
	}
	return s[n]
}

// func f_strait(n int64) (sum int64) {
// 	sum = 1
// 	for i := 3; i <= n; i++ {

// 	}
// 	return
// }
