package order_strategy

import (
	"serach_service/domain/dataset_index"
	"sort"
)

type OrderStratergy interface {
	Order(data []*dataset_index.IndexDocMetaData) []*dataset_index.IndexDocMetaData
}

type OrderByTokenFreq struct{}
func (obs *OrderByTokenFreq) Order(dataList []*dataset_index.IndexDocMetaData) []*dataset_index.IndexDocMetaData {
	sort.Slice(dataList, func(i, j int) bool {
		return dataList[i].TokenFrequence > dataList[j].TokenFrequence
	})
	return dataList
}

type OrderBySearchFreq struct{}
func (orderBySearchFreq *OrderBySearchFreq) Order(dataList []*dataset_index.IndexDocMetaData) []*dataset_index.IndexDocMetaData {
	sort.Slice(dataList, func(i, j int) bool {
		return dataList[i].SearchFrequency > dataList[j].SearchFrequency
	})
	return dataList
}
