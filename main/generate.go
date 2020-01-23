package main

import (
	"fmt"
	"math/rand"
	"outSort/fileOp"
)

var test = "test.in"

func Product(Cnt int) {

	RndChan := make(chan int)
	go func() {
		for i := 0; i < Cnt; i++ {
			RndChan <- rand.Int()
		}
		close(RndChan)
	}()
	fileOp.WriteToFile(test, RndChan)

	testRead := fileOp.ReadFile(test)
	nowCnt := 0
	for {
		select {
		case v := <-testRead:
			fmt.Println(v)
			nowCnt++
			if nowCnt > 50 {
				fmt.Println("...")
				return
			}
		}
	}
}
