package schema

import (
	"dbms_service/domain/schema_member"
	"fmt"
)

type Schema struct {
	columnToSchem map[string]schema_member.SchemaMember
}

func (s *Schema) ValidateSchema(rowData map[string]interface{}) (err error) {
	
	// check if all required field present
	for columnName, schemaMem := range s.columnToSchem {
		if schemaMem.IsRequired() {
			_, ok := rowData[columnName]
			if !ok {
				return fmt.Errorf("required column")
			}
		}
	}

	for columnName, value := range rowData {
		schemaMem, ok := s.columnToSchem[columnName]
		if !ok {
			return fmt.Errorf("column not in schema")
		}
		err = schemaMem.Validate(value)
		if err != nil {
			return
		}
	}

	return nil
}

func NewSchema(columnToSchem map[string]schema_member.SchemaMember) *Schema {
	return &Schema{
		columnToSchem: columnToSchem,
	}
}