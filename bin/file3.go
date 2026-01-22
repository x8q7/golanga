package main

import (
	"bufio"
	"fmt"
	"os"
)

//程序中的数据结构如下，是一个包含以下字段的结构:
//
//type Page struct {
//	Title string
//	Body  []byte
//}
//请给这个结构编写一个 save() 方法，将 Title 作为文件名、Body 作为文件内容，写入到文本文件中。
//
//再编写一个 load() 函数，接收的参数是字符串 title，该函数读取出与 title 对应的文本文件。请使用 *Page 做为参数，因为这个结构可能相当巨大，我们不想在内存中拷贝它。请使用 ioutil 包里的函数

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() {
	file, err := os.OpenFile(p.Title, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(string(p.Body))
}

func (p *Page) load(title string) (body string) {
	file, err := os.Open(p.Title)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		body += line
	}
	return body
}

func main() {
	f := &Page{Title: "index.html", Body: []byte("<html></html>")}
	f.save()

	fmt.Println(f.load("index.html"))
}
