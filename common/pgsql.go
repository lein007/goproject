package common

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Pgsql struct {
	Conn *sql.DB
}

//获取db
func (self *Pgsql) Getdb(constr string) {
	db, err := sql.Open("postgres", constr)
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	self.Conn = db
}

//插入
func (self *Pgsql) Insert(sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := self.Conn.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	return result.LastInsertId()
}
func (self *Pgsql) Execsql(sql string) error {
	_, err := self.Conn.Exec(sql)
	if err != nil {
		fmt.Printf("exec sql error: %v\n", err)
		return err
	}
	return nil
}

//修改和删除
func (self *Pgsql) Exec(sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := self.Conn.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	return result.RowsAffected()
}

//取一行数据，注意这类取出来的结果都是string
func (self *Pgsql) FetchRow(sqlstr string, args ...interface{}) (map[string]string, error) {
	stmtOut, err := self.Conn.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make(map[string]string, len(scanArgs))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			ret[columns[i]] = value
		}
		break //get the first row only
	}
	return ret, nil
}

//取多行，注意这类取出来的结果都是string
func (self *Pgsql) FetchRows(sqlstr string, args ...interface{}) ([]map[string]string, error) {
	stmtOut, err := self.Conn.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return ret, nil
}
