package fileOp

import (
	"fmt"
	"os"
	"testing"
)

func TestToChannel(t *testing.T) {
	fmt.Println("test ToChannel")
	Ch := ToChannel(3, 1, 3, 4, 5, 2, 1)
	want := []int{3, 1, 3, 4, 5, 2, 1}
	cnt := 0
	for v := range Ch {
		if v != want[cnt] {
			t.Errorf("error in %d\n", cnt)
		}
		cnt++
	}
}

func TestWriteToFile(t *testing.T) {
	fileName := "test.out"
	Ch := ToChannel(0, 1, 2, 5, 1, 4)
	WriteToFile(fileName, Ch)
}

func TestReadFile(t *testing.T) {
	fileName := "test.out"
	Ch := ToChannel(0, 1, 2, 5, 1, 4)
	a := [] int {0, 1, 2, 5, 1, 4};
	WriteToFile(fileName, Ch)
	ChRead := ReadFile(fileName);
	cnt := 0;
	for v := range(ChRead){
		if(v != a[cnt]){
			t.Errorf("error in %d\n",cnt);
		}
		cnt ++;
	}
}

func TestReadReader(t *testing.T) {
	fileName := "test.out"
	Ch := ToChannel(0, 1, 5, 5, 1, 4)
	WriteToFile(fileName, Ch)
	File ,err := os.Open(fileName);
	if(err != nil){
		panic(err);
	}
	ChRead := ReadReader(File, 8 * 3);
	for v := range(ChRead){
		fmt.Println(v);
	}
}