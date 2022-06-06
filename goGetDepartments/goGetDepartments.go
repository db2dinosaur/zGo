package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	/* Use the NewRequest interface to allow us to set all of the headers */
	url := "http://192.168.225.225:5040/services/GILLJ/GetDepartments"
	var auth = "Basic " + base64.StdEncoding.EncodeToString([]byte("GILLJ:sausage"))

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))

}