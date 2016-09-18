package util

import (
	"bytes"
	"database/sql"
	"strings"
)

type MysqlTable struct {
}

/**
 * 查询表数值类型的主键
 *
 * @param conn
 * @param catalog
 * @param schema
 * @param table
 * @return
 */
func GetDigitalPrimaryKey(db *sql.DB, tableName string, keywordEscaper string) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(" desc ")
	buffer.WriteString(tableName)
	rows, err := db.Query(buffer.String())
	defer rows.Close()
	if err != nil {
		return "", err
	}
	var field, types, null, key, defaults, extra string
	for rows.Next() {
		rows.Scan(&field, &types, &null, &key, &defaults, &extra)
		if key != "" && strings.Contains(types, "int") {
			return field, nil
		}
	}
	return "", err
}

/**
 * 获取表的字段类型
 *
 * @param connection
 * @param table
 * @return
 */
func GetColumnTypes(db *sql.DB, tableName string, keywordEscaper string) (map[string]string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(" desc ")
	buffer.WriteString(tableName)
	rows, err := db.Query(buffer.String())
	defer rows.Close()
	if err != nil {
		//	fmt.Println(err)
		return nil, err
	}
	var field, types, null, key, defaults, extra string
	columnType := make(map[string]string, 0)
	for rows.Next() {
		rows.Scan(&field, &types, &null, &key, &defaults, &extra)
		//fmt.Println(field, types, null, key, defaults, extra, "-----------------------------")
		columnType[field] = types
	}
	return columnType, err
}

/**
 * 获取表的字段名称
 *
 * @param db
 * @param tableName
 * @return
 */
func GetMysqlColumnNames(db *sql.DB, tableName string, keywordEscaper string) ([]string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(" SELECT * FROM  ")
	buffer.WriteString(keywordEscaper)
	buffer.WriteString(tableName)
	buffer.WriteString(keywordEscaper)
	buffer.WriteString("  WHERE 1=1 ")
	buffer.WriteString(" Limit 1 ")
	rows, err := db.Query(buffer.String())
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	ColumnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	colums := make([]string, 0)
	for _, v := range ColumnNames {
		colums = append(colums, v)
	}
	return colums, err
}

/*
获取数据库连接
*/

func GetConnection(driverName, dataSourceName string, maxOpenConns, maxIdleConns int) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(-1)
	db.Ping()
	return db, err

}
