package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling urls in golang")

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Printf("Type of the result %T\n", result)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// working with query params
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) // map

	fmt.Println(qparams["coursename"])

	for _, value := range qparams {
		fmt.Println("Param is: ", value)
	}

	//creating an url by creating object of the *url.URL
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=barun",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
