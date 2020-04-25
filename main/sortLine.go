package main

import (
	"os"
	"outSort/fileOp"
	"outSort/mergeSort"
	"outSort/outSortConst"
)

/*
cntNum 用的是 数字个数
*/
func sortPre(filename string , tempFilename string, cntNum, blockSize int) {
	sum := 0
	for sum < cntNum {
		File , err := os.Open(filename)
		if err!=nil {
			panic(err)
		}
		File.Seek(int64(sum*outSortConst.SizeStep),0)
		now := cntNum - sum
		if now > blockSize {
			now = blockSize
		}
		chanNow := fileOp.ReadReader(File , outSortConst.SizeStep * now)
		res := mergeSort.SortChannel(chanNow)
		fileOp.WriteChannel(tempFilename,res)
		sum += now
	}
}


func sortMerge(tempFilename string , cntNum, blockSize int) <-chan outSortConst.Data {
	sum := 0
	var res []<-chan outSortConst.Data
	for sum < cntNum {
		File , err := os.Open(tempFilename)
		if err!=nil {
			panic(err)
		}
		File.Seek(int64(sum*outSortConst.SizeStep),0)
		now := cntNum - sum
		if now > blockSize {
			now = blockSize
		}
		chanNow := fileOp.ReadReader(File , outSortConst.SizeStep * now)
		res = append(res , chanNow)
		sum += now
	}
	return mergeSort.Divide(res...)
}