package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/otiai10/gosseract"
)

func main() {
	fmt.Println("Processing ....")

	src := "./bills/"
	start := time.Now()
	client := gosseract.NewClient()
	defer client.Close()

	filepath.Walk(src, func(fpath string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() != ".DS_Store" {
			client.SetImage(fpath)
			text, _ := client.Text()
			ioutil.WriteFile("./raw/"+info.Name()+"_data.txt", []byte(text), 0666)
			getFields("./raw/", info.Name()+"_data.txt")
		}
		return nil
	})

	fmt.Printf("Elapsed Time: %v", time.Since(start))
}

func getFields(path, fname string) {
	exp := "^[0-9]*$"
	src := "./fields/"
	var data string

	file, _ := os.Open(path + fname)
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		text, err := buf.ReadString('\n')
		if strings.Contains(text, "Amount") || strings.Contains(text, "amount") || strings.Contains(text, "TIN") || strings.Contains(text, "INVOICE") || strings.Contains(text, "invoice") || strings.Contains(text, "total") || strings.Contains(text, "Total") || strings.Contains(text, "TOTAL") || strings.Contains(text, "inr") || strings.Contains(text, "INR") || strings.Contains(text, "sale") || strings.Contains(text, "Sale") || strings.Contains(text, "rupees") || strings.Contains(text, "Rupees") || strings.Contains(text, "Thousand") || strings.Contains(text, "Hundred") || match(text, exp) || strings.Contains(text, "Customer") || strings.Contains(text, "customer") || strings.Contains(text, "Name") || strings.Contains(text, "name") {
			data += text
		}

		if err != nil {
			break
		}
	}

	ioutil.WriteFile(src+fname, []byte(data), 0777)
}

func match(str, exp string) bool {
	r := regexp.MustCompile(exp)
	return r.MatchString(str)
}
