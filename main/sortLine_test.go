package main

import (
	"testing"
	"time"
)

func Test_sortLine(t *testing.T) {
	Product(100);
	time.Sleep(1000);
	fileName := "test.in";
	res := sortLine(fileName , 100,4);
	bef := -1;
	cnt := 0;
	for v := range(res){
		if(v < bef && cnt > 0 ) {
			t.Errorf("Error in %d %d\n",cnt,v);
		}
		bef = v;
		cnt ++;
	}
	if(cnt!=100){
		t.Errorf("failed in loss data");
	}
}