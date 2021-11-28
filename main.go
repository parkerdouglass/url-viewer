package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Simply put Hello on the root directory just to make sure the server is up and running
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello")
	})
	// Allow people to go to /proxy and handle it with the handler function that is created
	http.HandleFunc("/proxy", handler)
	// Start the web server and listen for requests
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get all of the keys in the URL, put them in the keys[] then make sure everything is ok
	keys, ok := r.URL.Query()["url"]
	if !ok {
		fmt.Println("Something went wrong with querying for URL parameter")
	}
	// Get the URL from the key
	url := string(keys[0])
	// Send a GET request to the URL
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading file")
	}
	defer res.Body.Close()
	// Read the body of the request and make it usable by us
	page, err := ioutil.ReadAll(res.Body)
	// Use the writer and write the res.body to the page
	w.Write(page)
}
