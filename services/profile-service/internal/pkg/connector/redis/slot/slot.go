package slot

// Range Диапазон слотов
type Range struct {
	Start   int
	End     int
	ShardID int
}

type Slice []*Range

func (p Slice) Len() int {
	return len(p)
}

func (p Slice) Less(i, j int) bool {
	return p[i].Start < p[j].Start
}

func (p Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
