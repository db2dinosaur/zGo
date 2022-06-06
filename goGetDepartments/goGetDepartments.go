package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	/*
		values := map[string]string{"name": "John Doe", "occupation": "gardener"}
		json_data, err := json.Marshal(values)

		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post("https://httpbin.org/post", "application/json",
		    bytes.NewBuffer(json_data))
	*/
	resp, err := http.Post("http://192.168.225.225:5040/services/GILLJ/GetDepartment", "application/json", nil)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}
