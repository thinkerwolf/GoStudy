package function

import (
	//"runtime"
	"log"
)

// 匿名函数作为调试

func Debug() {

//	where := func() {
//		_, file, line, _ := runtime.Caller(1)
//		log.Printf("%s:%d", file, line)
//	}
	
	where := log.Print
	
	where()
	// some code
	FuncAsParm(100, AddParm)
	where()
	// some more code
	FuncAsParm(100, AddParm)
	where()

}
