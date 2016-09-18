package util

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetDigitalPrimaryKey(t *testing.T) {
	dataSourceName := "root:@tcp(localhost:3306)/zhangboyu?charset=utf8"
	db, err := GetConnection("mysql", dataSourceName, 10, 10)
	if err != nil {
		fmt.Println(err)
	}
	c, err := GetDigitalPrimaryKey(db, "cc", "")
	fmt.Println("------------------", c)
}

func TestGetColumnTypes(t *testing.T) {
	dataSourceName := "root:@tcp(localhost:3306)/zhangboyu?charset=utf8"
	db, err := GetConnection("mysql", dataSourceName, 10, 10)
	if err != nil {
		fmt.Println(err)
	}
	c, err := GetColumnTypes(db, "cc", "")
	for ok, v := range c {
		fmt.Println(ok, v)
	}
}
func TestGetMysqlColumnNames(t *testing.T) {
	dataSourceName := "root:@tcp(localhost:3306)/zhangboyu?charset=utf8"
	db, err := GetConnection("mysql", dataSourceName, 10, 10)
	if err != nil {
		fmt.Println(err)
	}
	colums, err := GetMysqlColumnNames(db, "user", "")
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range colums {
		fmt.Println(k, "----------------", v)
	}
}
