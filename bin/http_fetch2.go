package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	url := ""
	if !scan.Scan() {
		return
	}
	url = scan.Text()

	res, err := http.Get(url)
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
