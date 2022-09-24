package dialect

import "reflect"

type Dialect interface {
	// 将 Go 语言的类型转换为该数据库的数据类型
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

var dialectMap = map[string]Dialect{}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(key string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[key]
	return
}
