package Btree

import (
	"flag"
	"fmt"
	"testing"
)


var btreeDegree = flag.Int("degree", 4, "B-Tree degree")

func TestBT(t *testing.T) {
	tr := New(*btreeDegree)
	oplist := []int{1,1,1,1,3,3,2,2,3,3}
	val := []uint32{5,6,7,8,5,7,7,8,7,5}
	for i := 0 ; i < len(oplist) ; i++ {
		op := oplist[i]
		v := val[i]
		if op == 1 {
			tr.ReplaceOrInsert(Token{v , 0})
		}
		if op == 2 {
			tr.Delete(Token{v , 0})
		}
		if op == 3 {
			fmt.Println(tr.Search(Token{v , 0}))
		}
	}
}
