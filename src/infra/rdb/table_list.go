package rdb

const (
	_ = iota
	UsersTable
)

var tableList = map[int]string{
	UsersTable: "users",
}

func TableName(tableNum int) string {
	return tableList[tableNum]
}

func AllTableNames() []string {
	allTableNames := make([]string, len(tableList))
	i := 0
	for _, tableName := range tableList {
		allTableNames[i] = tableName
		i++
	}
	return allTableNames
}
