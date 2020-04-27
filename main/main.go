package main

import (
	"flag"
	"fmt"
	"os"
	"outSort/Btree"
	"outSort/fileOp"
	"outSort/outSortConst"
	"time"
)

/*
Author : dxyinme
*/
var filename string = "test.in"
var tempFilename string = "test.temp"
var outFilename string = "test.out"
var btreeDegree = flag.Int("degree", 4, "B-Tree degree")
func main() {
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
	f3,err3 := os.Create(outFilename)
	if err3!=nil {
		panic(err3)
	}
	f3.Close()
	Cnt := 1000000
	tr := Btree.New(*btreeDegree)
	Product(Cnt)
	File,err := os.Open(filename)
	if err!=nil {
		panic(err)
	}
	BtreeTime := time.Now()
	chanData := fileOp.ReadReader(File,-1)
	var pos int64 = 0
	for v := range chanData {
		tr.ReplaceOrInsert(Btree.Token{ValA: v.ValA , Pos: pos})
		pos += int64(outSortConst.SizeStep)
	}
	fmt.Printf("B tree build time : %v \n" , time.Since(BtreeTime))
	fmt.Println(" G ValA(查找关键字为ValA的字符串)\n " +
		"I ValA ValB(插入/替换一个记录)\n " +
		"D ValA (删除关键字为ValA的记录)\n " +
		"E (退出程序 , 进行归并排序)")
	FileSize := int64(Cnt * outSortConst.SizeStep)
	for {
		var ValA uint32
		var ValB string
		var ty string
		_, _ = fmt.Scanf("%s", &ty)
		if ty[0] == 'G' {
			_, _ = fmt.Scanf("%d", &ValA)
			Start := time.Now()
			o := tr.Search(Btree.Token{ValA: ValA , Pos : 0})
			if o == nil {
				fmt.Println("Can't find this Item")
			}else {
				data := fileOp.ReadPos(filename,o.(Btree.Token).Pos)
				fmt.Print(data.ValA)
				fmt.Print(" ")
				fmt.Println(string(data.ValB))
			}
			fmt.Printf("Search use time : %v \n " , time.Since(Start))
		}else if ty[0] == 'I' {
			_, _ = fmt.Scanf("%d %s", &ValA, &ValB)
			Start := time.Now()
			now := outSortConst.Data{ValA:ValA , ValB : []byte(ValB)}
			fileOp.WriteEnd(filename , now)
			tr.ReplaceOrInsert(Btree.Token{ValA: ValA , Pos: FileSize})
			FileSize += int64(outSortConst.SizeStep)
			Cnt ++
			fmt.Printf("Insert use time : %v \n " , time.Since(Start))
		}else if ty[0] == 'D' {
			_, _ = fmt.Scanf("%d" , &ValA)
			Start := time.Now()
			tr.Delete(Btree.Token{ValA: ValA , Pos : 0})
			fmt.Printf("Delete use time : %v \n " , time.Since(Start))
		}else if ty[0] == 'E' {
			break
		}else{
			continue
		}
	}
	Start := time.Now()
	sortPre(filename, tempFilename , Cnt , outSortConst.SortSize)
	res := sortMerge(tempFilename,Cnt,outSortConst.SortSize)
	fileOp.WriteToFile(outFilename , res)
	fmt.Printf("sort time : %v \n " , time.Since(Start))
}
