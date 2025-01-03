package main

import (
	"dbms_service/domain/schema"
	"dbms_service/domain/schema_member"
	"dbms_service/domain/table"
	"fmt"
)

func main() {

	schemaObj := make(map[string]schema_member.SchemaMember)
	schemaObj["name"] = schema_member.NewStrSchema("name", uint32(1), uint32(10), true)
	schemaObj["age"] = schema_member.NewIntSchema("age", int32(0), int32(100), true)

	schemaInst := schema.NewSchema(schemaObj)
	tableInst := table.NewTable("user", *schemaInst)
	
	err := tableInst.AddRow(map[string]interface{}{
		"name": string("roshan"),
		"age": int32(24),
	})
	if err != nil {
		fmt.Println("err : ", err)
	}
	err = tableInst.AddRow(map[string]interface{}{
		"name": string("rohan"),
		"age": int32(14),
	})
	if err != nil {
		fmt.Println("err : ", err)
	}
	err = tableInst.AddRow(map[string]interface{}{
		"name": string("unish"),
		"age": int32(-1),
	})
	if err != nil {
		fmt.Println("err : ", err)
	}
	tableInst.ShowAll()
}