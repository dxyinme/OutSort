package fileOp

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"outSort/outSortConst"
)

func ToChannel(array ...outSortConst.Data) <-chan outSortConst.Data {
	ret := make(chan outSortConst.Data, outSortConst.ChannelSize)
	go func() {
		for _, v := range array {
			ret <- v
		}
		close(ret)
	}()
	return ret
}

func EncodeBuf(data outSortConst.Data) []byte{
	buf := make([]byte , outSortConst.SizeByte + outSortConst.SizeInt32)
	A := make([]byte , outSortConst.SizeInt32)
	binary.BigEndian.PutUint32(A , data.ValA)
	for i := 0 ; i < outSortConst.SizeByte + outSortConst.SizeInt32 ; i++{
		if i < 4 {
			buf[i] = A[i]
		} else {
			buf[i] = data.ValB[i - 4]
		}
	}
	return buf
}

func DecodeBuf(buf []byte) outSortConst.Data{
	o := buf[0:4]
	ValA := binary.BigEndian.Uint32(o)
	ValB := buf[4:]
	return outSortConst.MakeData(ValA ,ValB)
}

/*
write Data
*/
func WriteData(writer io.Writer, v outSortConst.Data) {
	//fmt.Println("WriteData")
	buf := EncodeBuf(v)
	now := DecodeBuf(buf)
	if v.ValA != now.ValA {
		fmt.Println("err")
	}
	writer.Write(buf)
}

func WriteToFile(fileName string, Ch <-chan outSortConst.Data) {
	File, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(File)
READLINE:
	for {
		select {
		case v, ok := <-Ch :
			if ok == false {
				break READLINE
			} else {
				WriteData(writer, v)
			}
		}
	}
	defer File.Close()
	defer writer.Flush()
}

/**
Limit = (outSortConst.SizeInt32 + outSortConst.SizeByte) * numberOfYouWant
*/
func ReadReader(reader io.Reader, Limit int) <-chan outSortConst.Data {
	ret := make(chan outSortConst.Data, outSortConst.ChannelSize)
	go func() {
		byteRead := 0
		buf := make([]byte , outSortConst.SizeInt32 + outSortConst.SizeByte)
		pre := outSortConst.Data{0 , nil}
		for {
			now , err := reader.Read(buf)
			byteRead += now
			nowData := DecodeBuf(buf)
			if pre.ValA == nowData.ValA {
				break
			} else{
				pre = nowData
			}
			ret <- nowData
			if err != nil || (Limit != -1 && byteRead >= Limit) {
				break
			}
		}
		close(ret)
	}()
	return ret
}

func ReadPos(filename string , Pos int64) outSortConst.Data {
	buf := make([]byte , outSortConst.SizeByte + outSortConst.SizeInt32)
	File , err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer File.Close()
	File.Seek(Pos,0)
	_, err1 := File.Read(buf)
	if err1 != nil {
		panic(err)
	}
	return DecodeBuf(buf)
}

