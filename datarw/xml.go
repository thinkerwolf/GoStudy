package datarw

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// go默认的xml解析器是SAX解析
func ReadXml() {
	reader := strings.NewReader("<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>")

	decoder := xml.NewDecoder(reader)

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
		}
	}

}
