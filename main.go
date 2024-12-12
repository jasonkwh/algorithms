package main

import (
	myatoi "algorithms/nov2024/my_atoi"
	regexmatch "algorithms/nov2024/regex_match"
	"fmt"
	"strconv"

	_ "net/http/pprof"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	sig := <-c
	// 	fmt.Printf("got %s signal", sig)

	// 	cancel()
	// }()

	// go func() {
	// 	log.Fatal(http.ListenAndServe("127.0.0.1:6060", nil))
	// }()

	// go func() {
	// 	for {
	// 		test := atoi1("-042")
	// 		fmt.Println(test)

	// 		test2 := atoi2("-042")
	// 		fmt.Println(test2)
	// 	}
	// }()

	// <-ctx.Done()

	//test := palindrome.IsPalindrome(9)
	//fmt.Println(test)

	match := regexmatch.IsMatch("aab", "c*a*b")
	fmt.Println(match)
}

func atoi1(s string) int {
	return myatoi.MyAtoi(s)
}

func atoi2(s string) int {
	val, _ := strconv.Atoi(s)

	return val
}
