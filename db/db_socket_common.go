package db

import (
	"fmt"
)

const (
	SQL_WHERE_CLAUSE_EQ_STRING string = "and %s='%s'"
)

func (db_socket *DBSocket) gen_where_clause_string(op, col, value string) string {
	return fmt.Sprintf(op, col, value)
}

func (db_socket *DBSocket) fmt_where_clause_sql(dst_sql, where_clause string) string {
	return fmt.Sprintf(dst_sql, where_clause)
}
