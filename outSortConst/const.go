package outSortConst

const (
	ChannelSize int = 1024
	SizeInt64 int = 8
	SizeInt32 int = 4
	SizeByte int = 12
)

type Data struct {
	ValA uint32
	ValB []byte
}

func MakeData(A uint32,B []byte) Data{
	var T Data
	T.ValA = A
	T.ValB = B
	return T
}