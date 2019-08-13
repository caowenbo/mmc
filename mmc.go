package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const SELECT = "select"
const SHOW = "show"

type SelectResultRow map[string]interface{}

func GetSelect(rows *sql.Rows) []SelectResultRow {
	var selectResult []SelectResultRow
	cols, err := rows.Columns()
	if err != nil {
		return nil
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil
	}
	for rows.Next() {
		vals := make([]interface{}, len(cols))
		for i, _ := range cols {
			switch colTypes[i].DatabaseTypeName() {
			case "DOUBLE":
				vals[i] = new(float64)
			case "JSON", "TEXT":
				vals[i] = new(string)
			case "FLOAT":
				vals[i] = new(float32)
			case "INT", "SMALLINT", "TINYINT":
				vals[i] = new(int32)
			case "BIGINT", "DECIMAL":
				vals[i] = new(int64)
			default:
				vals[i] = new(string)
			}
		}
		selectResultRow := make(SelectResultRow, len(cols))
		err = rows.Scan(vals...)
		for i, _ := range cols {
			selectResultRow[cols[i]] = vals[i]
		}
		selectResult = append(selectResult, selectResultRow)
	}
	return selectResult
}
func ExecSql(strSql string, db *sql.DB) error {
	if strings.HasPrefix(strings.ToLower(strSql), SELECT) || strings.HasPrefix(strings.ToLower(strSql), SHOW) {
		rows, err := db.Query(strSql)
		if err != nil {
			fmt.Printf(`{"code":5,"message":"exec the sql:%s,error:%s","result":""}`, strSql, err.Error())
			return err
		} else {
			result := GetSelect(rows)
			jsonStr, err := json.Marshal(result)

			if err != nil {
				fmt.Printf(`{"code":6,"message":"query with sql:%s,error:%s","result":""}`, strSql, err.Error())
				return err
			}
			fmt.Printf(`{"code":0,"message":"OK","result":%s}`, jsonStr)
		}
	} else {
		rt, err := db.Exec(strSql)
		if err != nil {
			fmt.Printf(`{"code":5,"message":"exec the sql:%s,error:%s","result":""}`, strSql, err.Error())
			return err
		} else {
			LastInsertId, _ := rt.LastInsertId()
			RowsAffected, _ := rt.RowsAffected()
			fmt.Printf(`{"code":0,"message":"OK","result":{"insert":%d,"alter":%d}}`, LastInsertId, RowsAffected)
		}
	}
	return nil
}

func DBTool(args []string) {
	if len(args) != 4 {
		fmt.Println(`{"code":1,"message":"usage:dbcu host user password sql","result":""}`)
		os.Exit(1)
	}
	dsName := args[1] + ":" + args[2] + "@tcp(" + args[0] + ")/"

	db, err := sql.Open("mysql", dsName)
	if err != nil {
		fmt.Printf(`{"code":2,"message":"can not open db use:%s,%s","result":""}`, dsName, err.Error())
		os.Exit(2)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf(`{"code":3,"message":"can not connect to db use:%s,%s","result":""}`, dsName, err.Error())
		os.Exit(3)
	}
	err = ExecSql(args[3], db)
	if err != nil {
		os.Exit(4)
	}
}
func main() {
	args := os.Args[1:]
	DBTool(args)
}
