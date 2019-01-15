package main


import (
	"os"
	"fmt"
	"runtime"
	"github.com/GoStudy/function"
	"github.com/GoStudy/array"
	"github.com/GoStudy/base"
	"github.com/GoStudy/struction"
	"github.com/GoStudy/concurrent"
)

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
}

func readFile() {
	
os.Open("");
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