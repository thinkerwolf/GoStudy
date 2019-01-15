package struction
/** 反射 */
import (
	"reflect"
	"fmt"
)

func Reflection() {
	var h Man = Human{name:"Jack", age:56}
	
	t := reflect.TypeOf(h)
	f1, b := t.FieldByName("name")
	if b {
		fmt.Println("Relection Type and field : " , t, ", ", f1)
	}
	
}

