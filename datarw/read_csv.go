package datarw

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadProducts() {

	file, err := os.Open("products.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	el := list.New()

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, error := reader.ReadString('\n')
		if error != nil {
			break
		}
		ss := strings.Split(line, ";")
		price, _ := strconv.ParseFloat(ss[1], 64)
		quantity, _ := strconv.ParseInt(strings.Replace(ss[2], "\n", "", -1), 0, 64)
		p := Product{
			Title:    ss[0],
			Price:    price,
			Quantity: int(quantity)}
		el.PushFront(p)
	}

	for e := el.Front(); e != nil; e = e.Next() {
		p, _ := e.Value.(Product)
		fmt.Println(p)
	}

}

type Product struct {
	Title    string
	Price    float64
	Quantity int
}
