package main

import (
	//"github.com/GoStudy/web"
	"github.com/GoStudy/databasing"
	mapTest "github.com/GoStudy/map_test"
)

func main() {
	//BaseLearning();
	//web.StartSimplerServer()
	databasing.MysqlDatabase()
	mapTest.MakeMap()
	mapTest.SortMap()
	mapTest.InvertMap()
}
