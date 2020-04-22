package main

import (
	"fmt"
	"math/rand"
	"outSort/fileOp"
	"outSort/outSortConst"
	"time"
)

var test = "test.in"

/*
随机生成一个长度为12的，只包含英文字母的字符串
*/
func RandString() []byte {
	bytes := make([]byte, 12)
	for i := 0; i < 12; i++ {
		c := rand.Intn(2)
		b := rand.Intn(26)
		if c == 0 {
			b += 'A'
		} else {
			b += 'a'
		}
		bytes[i] = byte(b)
	}
	return bytes
}

func Product(Cnt int){
	fmt.Println("Product")
	rand.Seed(time.Now().Unix())
	RndChan := make(chan outSortConst.Data)
	valArray := make([]uint32 , Cnt)
	for i := 1 ; i <= Cnt ; i++ {
		valArray[i-1] = uint32(i)
	}
	// 生成一个数组，1~Cnt ， 打乱顺序后作为ValA
	rand.Shuffle(len(valArray) , func(i, j int) {
		valArray[i],valArray[j] = valArray[j],valArray[i]
	})
	go func() {
		for i := 0; i < Cnt; i++ {
			now := outSortConst.MakeData(valArray[i] , RandString())
			//fmt.Print(now.ValA)
			//fmt.Print(" ")
			//fmt.Println(string(now.ValB))
			RndChan <- now
		}
		close(RndChan)
	}()
	fileOp.WriteToFile(test, RndChan)
}
