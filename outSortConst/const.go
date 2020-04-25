package outSortConst

const (
	SortSize int = 65536
	ChannelSize int = 65536
	SizeInt64 int = 8
	SizeInt32 int = 4
	SizeByte int = 12
	SizeStep int = 16
)


type Data struct {
	ValA uint32
	ValB []byte
}

type DataList []Data

func (m DataList) Len() int {
	return len(m)
}
func (m DataList) Less(i, j int) bool {
	return m[i].ValA < m[j].ValA
}
func (m DataList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func MakeData(A uint32,B []byte) Data{
	var T Data
	T.ValA = A
	T.ValB = B
	return T
}

