package gen

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// DBType database type
type DBType string

const (
	// dbMySQL Gorm Drivers mysql || postgres || sqlite || sqlserver
	dbMySQL      DBType = "mysql"
	dbPostgres   DBType = "postgres"
	dbSQLite     DBType = "sqlite"
	dbSQLServer  DBType = "sqlserver"
	dbClickHouse DBType = "clickhouse"
)

// CmdParams is command line parameters
type CmdParams struct {
	DSN               string   `yaml:"dsn"`               // consult[https://gorm.io/docs/connecting_to_the_database.html]"
	DB                string   `yaml:"db"`                // input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
	Tables            []string `yaml:"tables"`            // enter the required data table or leave it blank
	OnlyModel         bool     `yaml:"onlyModel"`         // only generate model
	OutPath           string   `yaml:"outPath"`           // specify a directory for output
	OutFile           string   `yaml:"outFile"`           // query code file name, default: gen.go
	WithUnitTest      bool     `yaml:"withUnitTest"`      // generate unit test for query code
	ModelPkgName      string   `yaml:"modelPkgName"`      // generated model code's package name
	FieldNullable     bool     `yaml:"fieldNullable"`     // generate with pointer when field is nullable
	FieldCoverable    bool     `yaml:"fieldCoverable"`    // generate with pointer when field has default value
	FieldWithIndexTag bool     `yaml:"fieldWithIndexTag"` // generate field with gorm index tag
	FieldWithTypeTag  bool     `yaml:"fieldWithTypeTag"`  // generate field with gorm column type tag
	FieldSignable     bool     `yaml:"fieldSignable"`     // detect integer field's unsigned type, adjust generated data type
}

func JudgeNilSlice(ls []string) bool {
	return len(ls) == 0
}
func operateSliceByAppend(ls []string) []string {
	tableList := make([]string, 0, len(ls))
	for _, tableName := range ls {
		_tableName := strings.TrimSpace(tableName) // trim leading and trailing space in tableName
		if _tableName == "" {                      // skip empty tableName
			continue
		}
		tableList = append(tableList, _tableName)
	}

	return tableList

}
func operateSliceByAddr(ls []string) []interface{} {
	tableList := make([]interface{}, len(ls))
	for i, tableName := range ls {
		tableList[i] = tableName
	}
	return tableList
}

func CreateStruct() {
	var cmdParse CmdParams
	fmt.Printf("%+v\n", cmdParse)
}

func SwitchCase() {
	_, _ = connectDB(DBType("mysql"), "")
}

func connectDB(t DBType, dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("dsn cannot be empty")
	}
	switch t {
	case dbMySQL:
		return nil, nil
	case dbPostgres:
		return nil, nil
	case dbSQLite:
		return nil, nil
	case dbSQLServer:
		return nil, nil
	case dbClickHouse:
		return nil, nil
	default:
		return nil, fmt.Errorf("unknow db %q (support mysql || postgres || sqlite || sqlserver for now)", t)
	}
}
