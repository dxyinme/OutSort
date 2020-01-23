package mergeSort

import (
	"fmt"
	"outSort/outSortConst"
)

func Merge(L, R <-chan int) <-chan int {
	ret := make(chan int, outSortConst.ChannelSize)
	go func() {
		vLeft, hasLeft := <-L
		vRight, hasRight := <-R
		for hasLeft || hasRight {
			if hasLeft && hasRight {
				if vLeft < vRight {
					ret <- vLeft
					vLeft, hasLeft = <-L
				} else {
					ret <- vRight
					vRight, hasRight = <-R
				}
			} else {
				if hasLeft {
					ret <- vLeft
					vLeft, hasLeft = <-L
				} else {
					ret <- vRight
					vRight, hasRight = <-R
				}
			}
		}
		close(ret)
	}()
	return ret
}

/*
	Blocks[i] 是一个无序的chan int
*/
func Divide(Blocks ...<-chan int) <-chan int {
	LenB := len(Blocks)
	if LenB == 1 {
		Blocks[0] = SortChannel(Blocks[0]);
		return Blocks[0]
	}
	mid := LenB / 2
	return Merge(Divide(Blocks[0:mid]...), Divide(Blocks[mid:LenB]...))
}

func mergeWork(L, R []int) []int {
	var ret []int
	cl, cr := len(L), len(R)
	nl, nr := 0, 0
	for nl < cl && nr < cr {
		if L[nl] <= R[nr] {
			ret = append(ret, L[nl])
			nl++
		} else {
			ret = append(ret, R[nr])
			nr++
		}
	}
	for nl < cl {
		ret = append(ret, L[nl])
		nl++
	}
	for nr < cr {
		ret = append(ret, R[nr])
		nr++
	}
	return ret
}

func divWork(a []int) []int {
	Len := len(a)
	if Len == 1 {
		return a
	}
	mid := Len / 2
	return mergeWork(divWork(a[0:mid]), divWork(a[mid:Len]))
}

func SortSlice(now []int) []int {
	// sort.Ints(now)
	fmt.Println("sortSlice")
	return divWork(now)
}

func SortChannel(Vec <-chan int) <-chan int {
	ret := make(chan int, outSortConst.ChannelSize)
	go func() {
		var now []int
		for v := range Vec {
			now = append(now, v)
		}
		now = SortSlice(now)
		for _, v := range now {
			ret <- v
		}
		close(ret)
	}()
	return ret
}
