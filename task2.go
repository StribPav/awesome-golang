package main

import (
"fmt"
"io/ioutil"
"log"
"net/http"
"os"
)

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
