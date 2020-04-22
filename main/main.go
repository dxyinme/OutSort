package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"outSort/Btree"
	"outSort/fileOp"
	"outSort/outSortConst"
	"runtime"
	"time"
)

/*
Author : dxyinme
*/
var filename string = "test.in"
var btreeDegree = flag.Int("degree", 4, "B-Tree degree")
func main() {

	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1) // 开启对阻塞操作的跟踪

	go func() {
		// pprof
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		if err := http.ListenAndServe(":9590", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	Step := outSortConst.SizeByte + outSortConst.SizeInt32
	Cnt := 100
	tr := Btree.New(*btreeDegree)
	Product(Cnt)
	time.Sleep(5000)
	File,err := os.Open(filename)
	if err!=nil {
		panic(err)
	}
	chanData := fileOp.ReadReader(File,-1)
	var pos int64 = 0
	for v := range chanData {
		tr.ReplaceOrInsert(Btree.Token{ValA: v.ValA , Pos: pos})
		pos += int64(Step)
	}
	fmt.Println(" G ValA(查找关键字为ValA的字符串)\n " +
		"I ValA ValB(插入/替换一个记录)\n " +
		"D ValA (删除关键字为ValA的记录)\n " +
		"E (退出程序 , 进行归并排序)")
	FileSize := int64(Cnt * Step)
	for {
		var ValA uint32
		var ValB string
		var ty string
		_, _ = fmt.Scanf("%s", &ty)
		if ty[0] == 'G' {
			_, _ = fmt.Scanf("%d", &ValA)
			o := tr.Search(Btree.Token{ValA: ValA , Pos : 0})
			if o == nil {
				fmt.Println("Can't find this Item")
			}else {
				data := fileOp.ReadPos(filename,o.(Btree.Token).Pos)
				fmt.Print(data.ValA)
				fmt.Print(" ")
				fmt.Println(string(data.ValB))
			}
		}else if ty[0] == 'I' {
			_, _ = fmt.Scanf("%d %s", &ValA, &ValB)
			now := outSortConst.Data{ValA:ValA , ValB : []byte(ValB)}
			fileOp.WriteEnd(filename , now)
			tr.ReplaceOrInsert(Btree.Token{ValA: ValA , Pos: FileSize})
			FileSize += int64(Step)
		}else if ty[0] == 'D' {
			_, _ = fmt.Scanf("%d" , &ValA)
			tr.Delete(Btree.Token{ValA: ValA , Pos : 0})
		}else if ty[0] == 'E' {
			break
		}else{
			continue
		}
	}
	res := sortLine(filename , Cnt , 4)
	fileOp.WriteToFile("test.out" , res)
}
