package datarw

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Page page
type Page struct {
	Title string
	Body  []byte
}

// Save page
func Save(page Page) error {
	fileName := page.Title + ".txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	fw := bufio.NewWriter(file)
	fw.Write(page.Body)
	fw.Flush()
	return nil
}

// Load page
func Load(title string) (*Page, error) {
	fileName := title + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	bs, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		return nil, err1
	}
	page := Page{Title: title}
	page.Body = bs
	return &page, nil
}

// SaveLoadPage test
func SaveLoadPage() {
	title := "爱如潮水"
	page := Page{
		Title: title,
		Body:  []byte("321321cccc\n32132142cafe")}

	Save(page)

	pageLoad, err := Load(title)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*pageLoad)
	fmt.Println(pageLoad.Title)
}
