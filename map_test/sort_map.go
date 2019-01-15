package map_test

import (
	"fmt"
	"sort"
)

/** 
  对map进行排序
 */
func SortMap() {
	m := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
                            "delta": 87, "echo": 56, "foxtrot": 12,
                            "golf": 34, "hotel": 16, "indio": 87,
                            "juliet": 65, "kili": 43, "lima": 98}
	
	var ks []string = make([]string, len(m))
	i := 0;
	for k, _ := range m {
		ks[i] = k
		i++
	}
	
	sort.Strings(ks)
	
	for _, k := range ks {
		fmt.Printf("Key: %v, Value: %v / ", k, m[k]);
	}
	
}

