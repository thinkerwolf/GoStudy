package main

import (
	"fmt"
	"runtime"

	"github.com/GoStudy/array"
	"github.com/GoStudy/base"
	"github.com/GoStudy/concurrent"
	"github.com/GoStudy/datarw"
	"github.com/GoStudy/function"
	"github.com/GoStudy/gopackage"
	mapTest "github.com/GoStudy/map_test"
	"github.com/GoStudy/struction"
	"github.com/GoStudy/web"
)

func main() {
	// fmt.Println("============ base ============")
	// BaseLearning()

	// fmt.Println("============ web ============")
	// WebTest()

	// fmt.Println("============ map ============")
	// MapTest()

	// fmt.Println("============ database ============")
	// databasing.MysqlDatabase()

	fmt.Println("============ standard lib ============")
	StandardLibrary()

	fmt.Println("============ read write data ============")
	ReadWriteData()

}

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
}

func BaseLearning() {
	fmt.Println("Hello, GO!")

	base.Tim()
	base.Pointer()
	base.StringPointer()
	base.SwitchFunc()
	base.ForFunc()
	base.Defer()

	// 方法
	addR, minusR := function.Add(4, 6)
	fmt.Printf("Add Result %d\n", addR)
	fmt.Printf("Minus Result %d\n", minusR)
	var a1, b1 int = 100, 5
	var multiR = function.Mutiply(a1, &b1)
	fmt.Printf("Multply Result %d, a1#%d, b1#%d\n", multiR, a1, b1)
	function.FuncAsParm(5, function.AddParm)
	function.FuncAnony()
	function.Debug()

	// 数组和切片
	array.ArrayFunc()
	array.SliceFunc()

	struction.TestStruct()
	struction.AnnomyStruct()
	struction.TestInterface()
	struction.Reflection()

	concurrent.GoRoutine()
	concurrent.Channel()
	concurrent.BufferedChannel()

	concurrent.SelectChannel()
}

func WebTest() {
	web.StartSimplerServer()
}

func MapTest() {
	mapTest.MakeMap()
	mapTest.SortMap()
	mapTest.InvertMap()
}

func StandardLibrary() {
	gopackage.AList()
	gopackage.UnSafe()
}

func ReadWriteData() {
	//datarw.ReadFile()
	datarw.ReadProducts()
	datarw.ReadProducts2()
	datarw.ReadGzipFile()
	datarw.SaveLoadPage()
	datarw.StartArgs()
	datarw.ObjToJSON()
	datarw.JSONToObj()
	datarw.ReadXml()
	datarw.GobEncodeOrDecode()
	datarw.GobFile()
}
