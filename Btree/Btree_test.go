package Btree

import (
	"flag"
	"fmt"
	"testing"
)


var btreeDegree = flag.Int("degree", 4, "B-Tree degree")

func TestBT(t *testing.T) {
	tr := New(*btreeDegree)
	for i := 0; i < 10; i++ {
		tr.ReplaceOrInsert(Token{uint32(i),int64(i)})
	}
	fmt.Println("len:       ", tr.Len())
	fmt.Println("get3:      ", tr.Search(Token{uint32(3),int64(1)}))
	fmt.Println("get100:    ", tr.Search(Token{uint32(100),int64(1)}))
	fmt.Println("del4:      ", tr.Delete(Token{uint32(4),int64(1)}))
	fmt.Println("del100:    ", tr.Delete(Token{uint32(100),int64(1)}))
	fmt.Println("replace5:  ", tr.ReplaceOrInsert(Token{uint32(5),int64(1)}))
	fmt.Println("replace100:", tr.ReplaceOrInsert(Token{uint32(100),int64(1)}))
	fmt.Println("get5:      ", tr.Search(Token{uint32(5),int64(5)}))
	fmt.Println("len:       ", tr.Len())
}
