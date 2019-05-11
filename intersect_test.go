package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	set "github.com/deckarep/golang-set"
)

var (
	n      = 1000
	strLen = 3
)

func TestIntersect(t *testing.T) {
	s1 := generateRandStr(2*n, strLen)
	s2 := generateRandStr(10*n, strLen)
	begin := time.Now()
	s3 := Intersect(s1, s2)
	fmt.Println(len(s3))
	fmt.Printf("slice intersect: %s\n", time.Now().Sub(begin))
}

func TestSetIntersect(t *testing.T) {
	s1 := generateRandStr(2*n, strLen)
	s2 := generateRandStr(10*n, strLen)
	begin := time.Now()
	s1set := set.NewSetFromSlice(s1)
	s2set := set.NewSetFromSlice(s2)
	s3set := s1set.Intersect(s2set).ToSlice()
	fmt.Println(len(s3set))
	fmt.Printf("set intersect: %s\n", time.Now().Sub(begin))
}

func generateRandStr(n int, strLen int) []interface{} {
	ss := make([]interface{}, n)
	for i := 0; i < n; i++ {
		ss[i] = RandString(strLen)
	}
	rand.Shuffle(len(ss), func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})
	return ss
}
