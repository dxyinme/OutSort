package mergeSort

import (
	"fmt"
	"outSort/fileOp"
	"outSort/outSortConst"
	"reflect"
	"sort"
	"testing"
)

func TestDivide(t *testing.T) {
	fmt.Println("test Divide")
	want := []int{1, 2, 3, 3, 4, 4, 5, 5, 6, 6 ,7}
	newBlocks := Divide(fileOp.ToChannel(4,2,5,1),
		fileOp.ToChannel(6,3,6),
		fileOp.ToChannel(4,3,7,5));
	cnt := 0
	for v := range newBlocks {
		if v != want[cnt] {
			t.Errorf("Error in %d\n", cnt)
		}
		cnt++
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		L <-chan int
		R <-chan int
	}
	tests := []struct {
		name string
		args args
		want <-chan int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.L, tt.args.R); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SortChannel(t *testing.T) {
	fmt.Println("test SortChannel")
	Ch := make(chan int, outSortConst.ChannelSize)
	bef := []int{1, 7, 8, 2, 3, 4}
	want := []int{1, 7, 8, 2, 3, 4}
	sort.Ints(want)
	for _, v := range bef {
		Ch <- v
	}
	close(Ch)
	ChAfter := SortChannel(Ch)
	cnt := 0
	for v := range ChAfter {
		if v != want[cnt] {
			t.Errorf("Error in %d\n", cnt)
		}
		// fmt.Printf("%d %d\n",cnt,v);
		cnt++
	}
}

func Test_SortSlice(t *testing.T) {
	fmt.Println("test SortSlice")
	bef := []int{1, 9, 2, 1, 4, 12, 338, 28, 6}
	aft := []int{1, 9, 2, 1, 4, 12, 338, 28, 6}
	want := []int{1, 9, 2, 1, 4, 12, 338, 28, 6}
	aft = SortSlice(aft)
	sort.Ints(want)
	for _, v := range aft {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	for _, v := range want {
		fmt.Printf("%d ", v)
	}
	Lb := len(bef)
	for i := 0; i < Lb; i++ {
		if aft[i] != want[i] {
			t.Errorf("Error in %d\n", i)
		}
	}
}
