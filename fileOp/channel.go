package fileOp

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
	"outSort/outSortConst"
)

func ToChannel(array ...int) <-chan int {
	ret := make(chan int, outSortConst.ChannelSize)
	go func() {
		for _, v := range array {
			ret <- v
		}
		close(ret)
	}()
	return ret
}

/*
write int64
*/
func WriteInt(writer io.Writer, v int) {
	buf := make([]byte, outSortConst.SizeInt64)
	binary.BigEndian.PutUint64(buf, uint64(v))
	writer.Write(buf)
}

func WriteToFile(fileName string, Ch <-chan int) {
	File, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(File)
	for v := range Ch {
		// fmt.Println(v)
		WriteInt(writer, v)
	}
	defer File.Close()
	defer writer.Flush()
}
/**
 Limit = outSortConst.SizeInt64 * numberOfYouWant
 */
func ReadReader(reader io.Reader,Limit int)<-chan int{
	ret := make(chan int, outSortConst.ChannelSize);
	go func() {
		buf := make([]byte, outSortConst.SizeInt64);
		byteRead := 0
		for {
			now, err := reader.Read(buf)
			byteRead += now
			if now > 0 {
				ret <- int(binary.BigEndian.Uint64(buf))
			}
			if err != nil || (Limit != -1 && byteRead >= Limit) {
				break
			}
		}
		close(ret)
	}()
	return ret;
}

func ReadFile(fileName string) <-chan int {
	File, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(File);
	return ReadReader(buf,-1);
}
