package datarw

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func ReadGzipFile() {
	fi, err := os.Open("gzip.txt.gz")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fi.Close()

	reader, err1 := gzip.NewReader(fi)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	bufReader := bufio.NewReader(reader)
	for {
		str, err2 := bufReader.ReadString('\n')
		if err2 != nil {
			break
		}
		fmt.Println(str)
	}

	/*buf := make([]byte, 500, 1024)
	str := ""

	for {
		len, readerError := reader.Read(buf)
		if len == 0 || readerError != nil {
			fmt.Println(readerError)
			break
		}

		str += string(buf)
	}*/

}
