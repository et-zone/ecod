package template

var TMPDATA = `package dao

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.con/et-zone/escan"
)

var sqlDB *sql.DB
//id 自增主键， fieldtag:"select"，不需要db操作的使用db:"-"
type {{FieldName .sName}} struct{
	{{range $index,$val := .data}}
	{{FieldName $val.fname}} {{$val.ftype}} ` + "`json:\"{{Jsondata $val.fname}}\" db:\"{{$val.fname}}\" fieldtag:\"insert,select\"`" + `
	{{end}}
}

//(自选功能，可删)对外提供结构化参数，update，insert，select使用
type {{FieldName .sName}}Type struct{
	{{range $index,$val := .data}}
	{{FieldName $val.fname}} *{{$val.ftype}} ` + "`json:\"{{Jsondata $val.fname}},omitempty\" db:\"{{$val.fname}}\"`" + `
	{{end}}
}

var {{.sName}}Build = escan.NewBuilder("{{.tname}}", new({{FieldName .sName}}))

func Insert{{FieldName .sName}}s(des *[]interface{}) error {
	sql, args := {{.sName}}Build.InsertBuilderSql(des)
	_, err := sqlDB.Exec(sql, args...)
	return err
}

func Update{{FieldName .sName}}(kv map[string]interface{}, conditions map[string]*escan.Condition) error {
	sql, args :=  {{.sName}}Build.UpdateBuilderSql(kv, conditions)
	_, err := sqlDB.Exec(sql, args...)
	return err
}

func Del{{FieldName .sName}}(conditions map[string]*escan.Condition) error {
	sql, args :=  {{.sName}}Build.DeleteBuilderSql(conditions)
	_, err := sqlDB.Exec(sql, args...)
	return err
}

func Select{{FieldName .sName}}s(conditions map[string]*escan.Condition, screen *escan.Screen) (*[]Stu, error) {
	sql, args :=  {{.sName}}Build.SelectBuilderSql([]string{}, conditions, screen)
	rows, err := sqlDB.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	stus := []Stu{}
	err = escan.NewEscan().ScanAll(&stus, rows)
	if err != nil {
		return nil, err
	}
	return &stus, err
}

//fields Can be empty,If it is empty, it will get all
func Select{{FieldName .sName}}sToMap(fields []string, conditions map[string]*escan.Condition, screen *escan.Screen) (*[]map[string]string, error) {
	sql, args :=  {{.sName}}Build.SelectBuilderSql(fields, conditions, screen)
	rows, err := sqlDB.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	{{.sName}}sMap := []map[string]string{}
	err = escan.NewEscan().ScanAll(&{{.sName}}Map, rows)
	if err != nil {
		return nil, err
	}
	for i, val := range {{.sName}}Map {
		v, err := escan.ChToJsonByTagDB(val, {{FieldName .sName}}{})
		if err != nil {
			{{.sName}}Map[i] = nil
		} else {
			{{.sName}}Map[i] = v
		}
	}
	return &{{.sName}}Map, err
}
//field Can be '',If it is '', it will get count(*)
func Select{{FieldName .sName}}sCount(field string, conditions map[string]escan.Condition, screen *escan.Screen) (int, error) {
	sql, args :=  {{.sName}}Build.SelectBuilderCountSql(field, conditions, screen)
	rows, err := sqlDB.Query(sql, args...)
	if err != nil {
		return 0, err
	}
	count := map[string]int{}
	err = escan.NewEscan().ScanOne(&count, rows)
	if err != nil {
		return 0, err
	}
	return count["c"], err
}

`
