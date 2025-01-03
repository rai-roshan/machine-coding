package table

import (
	"dbms_service/domain/index"
	"dbms_service/domain/schema"
	"fmt"
)

type Table struct {
	name       string
	schema     schema.Schema
	currentId  uint32
	rowsData   map[uint32]map[string]interface{}
	indexTable map[string]*index.Index
}

func NewTable(name string, schema schema.Schema) *Table {
	return &Table{
		name:       name,
		schema:     schema,
		currentId:  0,
		rowsData:   make(map[uint32]map[string]interface{}),
		indexTable: make(map[string]*index.Index),
	}
}

func (table *Table) AddRow(row map[string]interface{}) error {
	err := table.schema.ValidateSchema(row)
	if err != nil {
		return err
	}
	row["id"] = table.currentId
	table.rowsData[table.currentId] = row
	table.currentId++
	return nil
}

func (table *Table) Search(columnName string, value interface{}) (rowIds []uint32, err error) {
	// cehck if index exists
	columnIdx, ok := table.indexTable[columnName]
	if ok {
		rowIds = columnIdx.Search(value)
		return rowIds, nil
	}

	for rowId, rowData := range table.rowsData {
		columnValue, ok := rowData[columnName]
		if !ok {
			return []uint32{}, fmt.Errorf("no column found")
		}
		if columnValue == value {
			rowIds = append(rowIds, rowId)
		}
	}
	return rowIds, nil
}

func (table *Table) Update(rowId uint32, columnToValue map[string]interface{}) error {
	rowData, ok := table.rowsData[rowId]
	if !ok {
		return fmt.Errorf("row id not exists")
	}

	// discard old index for that row data
	for columnName, currentVal := range table.rowsData[rowId] {
		index, ok := table.indexTable[columnName]
		if !ok {
			continue
		}
		index.Remove(currentVal, rowId)
	}

	for columnName, newValue := range columnToValue {
		rowData[columnName] = newValue
	}
	table.rowsData[rowId] = rowData

	// update index if exists 
	for columnName, newValue := range table.rowsData[rowId] {
		index, ok := table.indexTable[columnName]
		if !ok {
			continue
		}
		index.AddRow(newValue, rowId)
	}
	return nil
}

func (table *Table) ShowAll() {
	fmt.Printf("data : %+v\n", table.rowsData)
}