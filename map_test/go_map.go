package map_test

import (
	"fmt"
)

func MakeMap() {
	
	map1 := new(map[string]int)
	
	fmt.Println(*map1)
	
	map2 := make(map[string]int, 20)
	map2["a"] = 23
	
	if value, ok := map2["a"]; ok {
		fmt.Println(value)
	}
	
	for key, value := range map2 {
		fmt.Println("key->", key, ", value->", value)
	}
	

}