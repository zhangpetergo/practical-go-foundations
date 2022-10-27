package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1Sum("day1/sha1/http.log.gz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sig: %v\n", sig)
	sig, err = sha1Sum("day1/sha1/sha1.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sig: %v\n", sig)

}

// 如果是.gz结尾的文件，我们对gzip解压缩后的文件进行Hash
// 不是，就对原内容进行Hash
func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file
	if strings.HasSuffix(file.Name(), ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		r = gz
	}

	// _, err = io.CopyN(os.Stdout, r, 100)
	// if err != nil {
	// 	return "", err
	// }

	w := sha1.New()
	_, err = io.Copy(w, r)
	if err != nil {
		return "", err
	}

	// 得到数字签名
	// w.Sum()形参是附加参数
	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
