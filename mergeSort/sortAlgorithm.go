package mergeSort

import (
	"fmt"
	"outSort/outSortConst"
	"sort"
)

func Merge(L, R <-chan outSortConst.Data) <-chan outSortConst.Data {
	ret := make(chan outSortConst.Data, outSortConst.ChannelSize)
	go func() {
		vLeft, hasLeft := <-L
		vRight, hasRight := <-R
		for hasLeft || hasRight {
			if hasLeft && hasRight {
				if vLeft.ValA < vRight.ValA {
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
	Blocks[i] 是一个有序的chan Data
*/
func Divide(Blocks ...<-chan outSortConst.Data) <-chan outSortConst.Data {
	LenB := len(Blocks)
	if LenB == 1 {
		Blocks[0] = Blocks[0]
		return Blocks[0]
	}
	mid := LenB / 2
	return Merge(Divide(Blocks[0:mid]...), Divide(Blocks[mid:LenB]...))
}

func mergeWork(L, R []outSortConst.Data) []outSortConst.Data {
	var ret []outSortConst.Data
	cl, cr := len(L), len(R)
	nl, nr := 0, 0
	for nl < cl && nr < cr {
		if L[nl].ValA <= R[nr].ValA {
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

func divWork(a []outSortConst.Data) []outSortConst.Data {
	Len := len(a)
	if Len == 1 {
		return a
	}
	mid := Len / 2
	return mergeWork(divWork(a[0:mid]), divWork(a[mid:Len]))
}

func SortSlice(now outSortConst.DataList) outSortConst.DataList {
	fmt.Println("sortSlice")
	sort.Sort(now)
	return now
}

func SortChannel(Vec <-chan outSortConst.Data) <-chan outSortConst.Data {
	ret := make(chan outSortConst.Data, outSortConst.ChannelSize)
	go func() {
		var now outSortConst.DataList
		READLOOP:
		for {
			select {
			case v, ok := <-Vec:
				if ok == false {
					break READLOOP
				}
				now = append(now, v)
			}
		}
		now = SortSlice(now)
		for _, v := range now {
			ret <- v
		}
		close(ret)
	}()
	return ret
}
