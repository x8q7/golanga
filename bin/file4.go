package main

import (
	"io/ioutil" // 用于文件读写（Go 1.16+ 也可替换为 os 包，文末附兼容方案）
	"log"
	"strings"
)

// 定义 Page 结构体（标题 + 内容字节数组）
type Page struct {
	Title string // 作为文件名（建议后缀 .txt，避免无扩展名）
	Body  []byte // 文件内容（[]byte 适配 ioutil 读写）
}

// ========== 1. save() 方法：将 Page 写入文件 ==========
// 接收者为 *Page：避免拷贝大结构体，直接操作原数据
func (p *Page) save() error {
	// 处理文件名：为 Title 拼接 .txt 后缀（避免无扩展名，增强可读性）
	filename := strings.TrimSpace(p.Title) + ".txt"
	// 校验标题非空
	if filename == ".txt" {
		return log.Fatal("错误：Page.Title 不能为空")
	}

	// ioutil.WriteFile：将 Body 写入文件
	// 参数：文件名、内容、文件权限（0644 表示读写权限）
	err := ioutil.WriteFile(filename, p.Body, 0644)
	if err != nil {
		return err // 返回错误，由调用方处理
	}
	log.Printf("成功保存文件：%s\n", filename)
	return nil
}

// ========== 2. load() 函数：从文件读取到 *Page ==========
// 参数：title（文件名前缀，无需 .txt 后缀）
// 返回：*Page（避免拷贝）、error（错误信息）
func load(title string) (*Page, error) {
	// 拼接文件名（和 save 方法保持一致）
	filename := strings.TrimSpace(title) + ".txt"
	// 校验 title 非空
	if filename == ".txt" {
		return nil, log.Fatal("错误：title 参数不能为空")
	}

	// ioutil.ReadFile：读取文件内容到 []byte
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err // 文件不存在/读取失败时返回错误
	}

	// 创建 Page 结构体指针并返回（避免拷贝）
	page := &Page{
		Title: title, // 还原 Title（不含 .txt 后缀）
		Body:  body,  // 读取的文件内容
	}
	log.Printf("成功加载文件：%s\n", filename)
	return page, nil
}

// ========== 测试代码 ==========
func main() {
	// 1. 测试 save() 方法
	testPage := &Page{
		Title: "golang教程",                      // 文件名：golang教程.txt
		Body:  []byte("Go 语言入门教程：指针、结构体、文件操作"), // 文件内容
	}
	// 保存文件
	err := testPage.save()
	if err != nil {
		log.Fatalf("保存失败：%v\n", err)
	}

	// 2. 测试 load() 函数
	loadedPage, err := load("golang教程")
	if err != nil {
		log.Fatalf("加载失败：%v\n", err)
	}
	// 打印加载结果
	log.Printf("加载的标题：%s\n", loadedPage.Title)
	log.Printf("加载的内容：%s\n", string(loadedPage.Body))
}
