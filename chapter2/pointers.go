package main

import "fmt"
/**
Mental note:

1. Using refs points to an address in memory rather than the value
2. Syntax is similar to php when passing a variable as a ref &$myvar, which
when comparing them will return false as the address in memory are different (value x ref)
 */
func main() {
	fmt.Println(f() == f())

	total := 1
	incr(&total)
	fmt.Println(total)
}

func f() *int  {
	v :=  1
	return &v
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}