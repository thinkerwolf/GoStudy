package struction

/** 反射 */
import (
	"fmt"
	"reflect"
)

func Reflection() {
	var h Man = Human{Name: "Jack", Age: 56}

	t := reflect.TypeOf(h)
	f1, b := t.FieldByName("name")
	if b {
		fmt.Println("Relection Type and field : ", t, ", ", f1)
	}

}
