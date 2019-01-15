package map_test

import (
	"fmt"
	"sort"
)

func InvertMap() {
	drinkMap := map[string]string {
		"juice":"果汁",
		"beer":"啤酒",
		"orange":"橘子",
	}
	
	for k, v := range drinkMap {
		fmt.Println("Key:", k, " Value:", v)
	}
	
	ks := make([]string, len(drinkMap))
	i := 0
	for k,_ := range drinkMap {
		ks[i] = k
		i++
	}
	
	sort.Strings(ks)
	fmt.Println("Sorted map")
	for _, k := range ks {
		fmt.Println("Key:", k, " Value:", drinkMap[k])
	}
	
}

