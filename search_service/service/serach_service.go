package search_service

import (
	"fmt"
	"serach_service/domain/dataset"
	"serach_service/domain/dataset_index"
	"serach_service/domain/document"
	order_strategy "serach_service/stratergy"
)

type SearchService interface {
	CreateDataSet(datasetId uint32) *dataset.DataSet
	CreateDocumentInDataset(dataset *dataset.DataSet, docId uint32, content string) *document.Document
	Search(keyword string, orderAlgo order_strategy.OrderStratergy)
}

type searchService struct {
	dataSetTable      map[uint32]*dataset.DataSet
	dataSetIndexTable map[uint32]*dataset_index.DataSetIndex
}

func NewSearchService() SearchService {
	return &searchService{
		dataSetTable:      make(map[uint32]*dataset.DataSet),
		dataSetIndexTable: make(map[uint32]*dataset_index.DataSetIndex),
	}
}

func (ss *searchService) CreateDataSet(datasetId uint32) *dataset.DataSet {
	datasetObj := dataset.NewDataSet(datasetId)
	datasetIndexObj := dataset_index.NewDatasetIndex(datasetId)

	ss.dataSetIndexTable[datasetId] = datasetIndexObj
	ss.dataSetTable[datasetId] = datasetObj

	fmt.Printf("dataset idx : %+v\n", ss.dataSetIndexTable)
	fmt.Printf("dataset table : %+v\n", ss.dataSetTable)
	return datasetObj
}

func (ss *searchService) CreateDocumentInDataset(dataset *dataset.DataSet, docId uint32, content string) *document.Document {
	doc := document.NewDocument(docId, content)
	datasetId := dataset.GetId()

	ss.dataSetIndexTable[datasetId].IndexDocument(doc)
	return doc
}

func (ss *searchService) Search(keyword string, orderAlgo order_strategy.OrderStratergy) {
	// result := make(map[uint32][]uint32)
	for datasetId, datasetIdx := range ss.dataSetIndexTable {
		resultMetaData := datasetIdx.Search(keyword)

		resultMetaData = orderAlgo.Order(resultMetaData)

		if len(resultMetaData) > 0 {
			fmt.Println("dataset id : ", datasetId)
			for _, metaData := range resultMetaData {
				fmt.Println("doc id : ", metaData.DocId)
			}
		}
	}
}

// func OrderByFreq(result []*dataset_index.IndexDocMetaData) []*dataset_index.IndexDocMetaData {
// 	sort.Slice(result, func(i, j int) bool {
// 		return result[i].TokenFrequence > result[j].TokenFrequence
// 	})
// 	return result
// }
