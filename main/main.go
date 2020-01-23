package main

import (
	"fmt"
	"outSort/fileOp"
)

/*
Author : dxyinme
*/
var Name string
var NameSorted string
var TyD string
var Cnt int
func main() {
	fmt.Printf("generate / sort / test ? (G / S / T)")
	fmt.Scanf("%s",&TyD);
	if(TyD[0] == 'S'){
		fmt.Printf("OutSort\nInput Your file name : ")
		fmt.Scanf("%s", &Name)
		fmt.Printf("Input Your sorted-file name : ")
		fmt.Scanf("%s", &NameSorted)
		fmt.Printf("Number of number : ");
		fmt.Scanf("%d" , &Cnt);
		fmt.Printf("sort {%s} into {%s}\n" , Name , NameSorted);
		SortedChan := sortLine(Name , Cnt , 8);
		fileOp.WriteToFile(NameSorted , SortedChan);
	} else if (TyD[0] == 'G'){
		fmt.Println("Generate!");
		fmt.Printf("input generate number:");
		fmt.Scanf("%d" , &Cnt);
		Product(Cnt);
	} else if(TyD[0]  == 'T'){
		fmt.Printf("Input Your sorted-file name : ")
		fmt.Scanf("%s", &NameSorted)
		SortedChan := fileOp.ReadFile(NameSorted);
		cnt := 0;
		bef := -1;
		for v := range(SortedChan){
			if(bef > v && cnt > 0){
				fmt.Errorf("error in %d : %d\n",cnt , v);
				return ;
			}
			cnt ++;
			bef = v;
		}
		fmt.Printf("test finished\n");
	}
}
