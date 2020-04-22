package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_sortLine(t *testing.T) {
	Product(101)
	time.Sleep(1000)
	fileName := "test.in"
	res := sortLine(fileName, 101 , 4)
	cnt := 0
	for v := range res {
		fmt.Println(v)
		cnt ++
	}
	fmt.Println(cnt)
	if cnt != 101 {
		t.Errorf("failed in loss data")
	}
}
