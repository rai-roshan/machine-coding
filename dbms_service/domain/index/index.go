package index

type Index struct {
	name  string
	table map[interface{}][]uint32
}

func NewIndex(name string) *Index {
	return &Index{
		name:  name,
		table: make(map[interface{}][]uint32),
	}
}

func (i *Index) AddRow(value interface{}, rowId uint32) {
	i.table[value] = append(i.table[value], rowId)
}

func (i *Index) Search(value interface{}) []uint32 {
	return i.table[value]
}

func (i *Index) Remove(value interface{}, rowId uint32) {
	var newRowIdsList []uint32
	for _, existRowId := range i.table[value] {
		if existRowId == rowId {
			continue
		} 
		newRowIdsList = append(newRowIdsList, existRowId)
	}
	i.table[value] = newRowIdsList
}