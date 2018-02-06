package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	res, err := http.Get("http://www.mcleods.com")
	if err != nil {
		log.Fatal(error)
	}

	page, err := ioutil.ReadAll(res.body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
