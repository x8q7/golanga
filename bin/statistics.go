package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const inputHtml = `
	    <html><body>
        <form action="#" method="post" name="bar">
            <input type="text" name="in" />
            <input type="submit" value="submit"/>
        </form>
    </body></html>
`

func calcResult(input []string) string {
	result := 0
	params := ""
	for _, v := range input {
		val, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			continue
		}
		params += v + " "
		result += val
	}
	resultStr := strconv.Itoa(result)
	eco := "<html><body>" +
		"<h4> input: " +
		params +
		"</h4>" +
		"<h4> result: " +
		resultStr +
		"</h4>" +
		"</body></html>"
	return eco
}

func formHtml(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch req.Method {
	case "GET":
		io.WriteString(w, inputHtml)
	case "POST":
		err := req.ParseForm()
		if err != nil {
			return
		}
		arr := req.Form["in"]
		resultHtml := calcResult(strings.Split(arr[0], " "))
		io.WriteString(w, resultHtml)

	default:
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "页面不存在")
	}
}

func main() {
	http.HandleFunc("/form", formHtml)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		writer.WriteHeader(http.StatusNotFound)
		io.WriteString(writer, "页面不存在")
	})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println("server start err")
		return
	}
}
