package datarw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

//ObjToJSON 序列化Json
func ObjToJSON() {
	address1 := Address{Type: "work", City: "Shanghai", Country: "China"}
	address2 := Address{Type: "home", City: "Chuzhou", Country: "Anhui"}
	vcard := VCard{FirstName: "Kai", LastName: "Wu", Addresses: []*Address{&address1, &address2}, Remark: "Personal"}

	js, _ := json.Marshal(vcard)
	fmt.Printf("JSON format: %s \n", js)

	// using an encoder: 把json转储到文件中
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vcard)
	if err != nil {
		log.Println("Error in encoding json")
	}
}

//JSONToObj 反序列化
func JSONToObj() {
	f, openErr := os.Open("vcard.json")
	if openErr != nil {
		log.Println(openErr)
		return
	}

	defer f.Close()

	js, _ := ioutil.ReadAll(f)

	vc := &VCard{}
	err := json.Unmarshal(js, vc)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(*vc)
}
