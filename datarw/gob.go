package datarw

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

// gob类似于Java中的Serialization，用于序列化对象，可用于RPC调用中
type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func GobEncodeOrDecode() {
	var buf bytes.Buffer
	//var buf *bytes.Buffer 会报错 ?????

	// 序列化
	enc := gob.NewEncoder(&buf)
	enc.Encode(P{X: 1, Y: 2, Z: 3, Name: "21321"})

	//fmt.Println(buf)

	// 反序列化
	dec := gob.NewDecoder(&buf)
	var q Q
	dec.Decode(&q)

	fmt.Println(*q.X, ",", *q.Y, ",", q.Name)
}

func GobFile() {
	// 序列化
	address1 := Address{Type: "work", City: "Shanghai", Country: "China"}
	address2 := Address{Type: "home", City: "Chuzhou", Country: "Anhui"}
	vcard := VCard{FirstName: "Kai", LastName: "Wu", Addresses: []*Address{&address1, &address2}, Remark: "Personal"}
	ObjToGobFile("vcard.gob", vcard)

	// 反序列化
	var vc VCard
	err := GobFileToObj("vcard.gob", &vc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vc)
}

func ObjToGobFile(filename string, obj interface{}) error {
	fw, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return err
	}
	enc := gob.NewEncoder(fw)
	return enc.Encode(obj)
}

func GobFileToObj(filename string, obj interface{}) error {
	fr, err := os.Open(filename)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(fr)
	return dec.Decode(obj)
}
