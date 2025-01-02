package dataset

import "serach_service/domain/document"

type DataSet struct {
	id              uint32
	documentsRecord map[uint32]*document.Document
}

func NewDataSet(id uint32) *DataSet {
	return &DataSet{
		id:              id,
		documentsRecord: make(map[uint32]*document.Document),
	}
}

func (ds *DataSet) GetId() uint32 {
	return ds.id
}

func (ds *DataSet) GetDocumentById(docId uint32) *document.Document {
	return ds.documentsRecord[docId]
}
