package main

import (
	b64 "encoding/base64"
	json "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {

	/* Use the NewRequest interface to allow us to set all of the headers */
	url := "http://192.168.225.225:5040/services/GILLJ/GetDepartments"
	var auth = "Basic " + b64.StdEncoding.EncodeToString([]byte("GILLJ:sausage"))

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

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
	log.Println(string(body))

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(body), &jsonMap)
	rj, err := json.MarshalIndent(jsonMap, "", "  ")
	if err != nil {
		log.Println("Error marshalling JSON : ", err)
	}
	log.Println(string(rj))
	stc, err := strconv.Atoi(jsonMap["StatusCode"])
	//	statusCode := stc.(map[int]interface{})
	if stc == 200 {
		log.Println("--- OKAY ---")
	} else {
		log.Println(">>> FAIL <<<\nStatusCode = ", stc)
	}

}
