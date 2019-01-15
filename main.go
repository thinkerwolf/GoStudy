package main

import (
	//"web"
	"databasing"
	mapTest "map_test"
)

func main() {
	//BaseLearning();
	//web.StartSimplerServer()
	databasing.MysqlDatabase()
	mapTest.MakeMap()
	mapTest.SortMap()
	mapTest.InvertMap()
}
