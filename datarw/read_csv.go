package datarw

import (
	"bufio"
	"container/list"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Title    string
	Price    float64
	Quantity int
}

// old way
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

		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, "\r")
		line = strings.TrimSuffix(line, "\n")

		if len(line) == 0 {
			continue
		}

		ss := strings.Split(line, ",")
		price, parseErr1 := strconv.ParseFloat(ss[1], 64)
		if parseErr1 != nil {
			fmt.Println(parseErr1)
			break
		}

		//quantity, parseErr2 := strconv.Atoi(strings.Replace(ss[2], "\r\n", "", -1))
		quantity, parseErr2 := strconv.Atoi(ss[2])
		if parseErr2 != nil {
			fmt.Println(parseErr2)
			break
		}

		p := Product{
			Title:    ss[0],
			Price:    price,
			Quantity: quantity}
		el.PushFront(p)
	}

	for e := el.Front(); e != nil; e = e.Next() {
		p, _ := e.Value.(Product)
		fmt.Println(p)
	}

}

// new way
func ReadProducts2() {
	file, err := os.Open("products.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		file.Close()
	}()

	reader := csv.NewReader(file)

	for {
		strs, error := reader.Read()
		if error != nil {
			fmt.Println(error)
			break
		}

		fmt.Println(strs)

	}

}
