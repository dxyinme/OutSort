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
func sortLine(fileName string, cntNum, blockNum int) <-chan outSortConst.Data {
	blockSize := cntNum / blockNum
	if cntNum % blockNum != 0 {
		blockSize ++
	}
	now := int64(0)
	var res []<-chan outSortConst.Data
	for i := 0; i < blockNum; i++ {
		File, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		File.Seek(now, 0)
		chanNow := fileOp.ReadReader(File, (outSortConst.SizeInt32 + outSortConst.SizeByte) * blockSize)
		res = append(res, chanNow)
		now += int64(blockSize * (outSortConst.SizeByte + outSortConst.SizeInt32))
	}
	//for _,v := range res {
	//	for {
	//		now , ok := <-v
	//		if ok == false {
	//			break
	//		}
	//		fmt.Printf(" %d",now.ValA)
	//	}
	//	fmt.Println()
	//}
	return mergeSort.Divide(res...)
}
