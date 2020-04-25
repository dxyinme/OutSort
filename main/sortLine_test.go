package main

import (
	"fmt"
	"os"
	"outSort/fileOp"
	"testing"
	"time"
)

func Test_sortPre(t *testing.T){
	filename := "test.in"
	tempFilename := "test.temp"
	f1,err1 := os.Create(filename)
	if err1!=nil {
		panic(err1)
	}
	f1.Close()
	f2,err2 := os.Create(tempFilename)
	if err2!=nil {
		panic(err2)
	}
	Product(101)
	time.Sleep(1000)

	f2.Close()
	sortPre(filename , tempFilename, 101 , 20)
	File,err := os.Open(tempFilename)
	if err != nil {
		panic(err)
	}
	chanNow := fileOp.ReadReader(File, 101*16)
	for v := range chanNow {
		fmt.Println(v)
	}
}

func Test_sortMerge(t *testing.T){
	Cnt := 1000000
	bls := 65536
	filename := "test.in"
	tempFilename := "test.temp"
	outFilename := "test.out"
	f1,err1 := os.Create(filename)
	if err1!=nil {
		panic(err1)
	}
	f1.Close()
	f2,err2 := os.Create(tempFilename)
	if err2!=nil {
		panic(err2)
	}
	f2.Close()
	Product(Cnt)
	sortPre(filename , tempFilename, Cnt , bls)
	fmt.Println("pre")
	ch := sortMerge(tempFilename , Cnt , bls)
	fileOp.WriteToFile(outFilename , ch)
	File,err := os.Open(outFilename)
	if err != nil {
		panic(err)
	}
	chanNow := fileOp.ReadReader(File, Cnt*16)
writeloop:
	for {
		select {
		case v, ok := <-chanNow:
			if ok == false {
				break writeloop
			}
			fmt.Println(v)
		}
	}
}