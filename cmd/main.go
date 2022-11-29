package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const MysqlConfig = "betterman:RNMTQ666!@(47.101.204.54:3306)/sfrookie?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(MysqlConfig))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: false,
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
		FieldSignable: false,
		FieldWithIndexTag: false, // generate with gorm index tag
		FieldWithTypeTag: true, // generate with gorm column type tag
	})
	g.UseDB(db)
	dataMap := map[string]func(detailType string) (dataType string){
		"tinyint":   func(detailType string) (dataType string) { return "int64" },
		"smallint":  func(detailType string) (dataType string) { return "int64" },
		"mediumint": func(detailType string) (dataType string) { return "int64" },
		"bigint":    func(detailType string) (dataType string) { return "int64" },
		"int":       func(detailType string) (dataType string) { return "int64" },
	}
	g.WithDataTypeMap(dataMap)

	//// 自定义模型结体字段的标签
	//// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	//jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	//	toStringField := `balance, `
	//	if strings.Contains(toStringField, columnName) {
	//		return columnName + ",string"
	//	}
	//	return columnName
	//})
	//// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	//// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	//// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	//autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	//autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	//softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")
	//// 模型自定义选项组
	//fieldOpts := []gen.ModelOpt{jsonField}
	allModel := g.GenerateAllTable()
	g.ApplyBasic(allModel...)

	g.Execute()
}
