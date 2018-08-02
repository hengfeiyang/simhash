package main

import (
	"fmt"

	"github.com/safeie/simhash"
)

func main() {
	s1 := simhash.Simhash("this is a project for golang implementation of simhash algorithm")
	s2 := simhash.Simhash("this is a project for java   implementation of simhash algorithm")

	fmt.Println("distance:", simhash.Distance(s1, s2))
	fmt.Println("similars:", simhash.Similar(s1, s2))
}
