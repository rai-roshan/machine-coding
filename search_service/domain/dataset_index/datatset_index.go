package dataset_index

import (
	"serach_service/domain/document"
	"strings"
)

type IndexDocMetaData struct {
	DocId           uint32
	TokenFrequence  uint32
	SearchFrequency uint32
}

type DataSetIndex struct {
	dataSetId  uint32
	indexTable map[string]map[uint32]*IndexDocMetaData
}

func NewDatasetIndex(datasetId uint32) *DataSetIndex {
	return &DataSetIndex{
		dataSetId:  datasetId,
		indexTable: make(map[string]map[uint32]*IndexDocMetaData),
	}
}

func (dsi *DataSetIndex) Search(keyword string) []*IndexDocMetaData {
	// break the keywords to token
	// search for each token
	// union the result

	tokenToFreq := getTokenFreq(keyword)

	tokenLen := len(tokenToFreq)

	docIdToFreq := make(map[uint32]uint32)
	docIdToDocMeta := make(map[uint32]*IndexDocMetaData)

	for token, _ := range tokenToFreq {
		docs, ok := dsi.indexTable[token]
		if !ok {
			continue
		}
		for _, docMeta := range docs {
			docIdToFreq[docMeta.DocId]++
			docIdToDocMeta[docMeta.DocId] = docMeta
		}
	}

	var resultDocIds []uint32
	for docId, freq := range docIdToFreq {
		if freq == uint32(tokenLen) {
			resultDocIds = append(resultDocIds, docId)
		}
	}

	var result []*IndexDocMetaData
	for _, docId := range resultDocIds {
		result = append(result, docIdToDocMeta[docId])
	}

	return result
}

func (dsi *DataSetIndex) IndexDocument(doc *document.Document) {
	// get content and break it into tokens
	// frequency of each keyword
	tokenToFreq := getTokenFreq(doc.GetContent())
	docId := doc.GetId()

	// update it in indexTable
	for token, freq := range tokenToFreq {
		_, ok := dsi.indexTable[token]
		if !ok {
			dsi.indexTable[token] = make(map[uint32]*IndexDocMetaData)
		}
		dsi.indexTable[token][docId] = &IndexDocMetaData{
			DocId:          docId,
			TokenFrequence: freq,
		}
	}
}

func getTokenFreq(content string) map[string]uint32 {
	tokens := strings.Fields(content)
	tokenToFreq := make(map[string]uint32)
	for _, token := range tokens {
		tokenToFreq[token]++
	}
	return tokenToFreq
}
