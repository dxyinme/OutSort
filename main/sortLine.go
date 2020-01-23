package main

import (
	"os"
	"outSort/fileOp"
	"outSort/mergeSort"
	"outSort/outSortConst"
)

func sortLine(fileName string, cntNum, blockNum int) <-chan int {
	blockSize := cntNum / blockNum
	var res []<-chan int
	for i := 0; i < blockNum; i++ {
		File, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		File.Seek(int64(i*blockSize), 0)
		chanNow := fileOp.ReadReader(File, outSortConst.SizeInt64*blockSize)
		res = append(res, chanNow)
	}
	return mergeSort.Divide(res...)
}
