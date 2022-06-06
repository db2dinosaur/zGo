package main

import (
	"bytes"
	b64 "encoding/base64"
	json "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	ep := len(os.Args)
	if ep > 3 {
		ep = 3
	}
	args := os.Args[1:ep]
	mgr := ""
	dept := ""
	if ep > 2 {
		mgr = args[1]
	}
	if ep > 1 {
		dept = args[0]
	}
	jsonReq := []byte(`{
		"mgr": "` + mgr + `",
		"dept":"` + dept + `"
	}`)

	/* Use the NewRequest interface to allow us to set all of the headers */
	url := "http://192.168.225.225:5040/services/GILLJ/GetEmployeesByDepartment"
	var auth = "Basic " + b64.StdEncoding.EncodeToString([]byte("GILLJ:sausage"))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
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
	//	log.Println(string(body))

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(body), &jsonMap)
	var statusCode = jsonMap["StatusCode"].(float64)
	stc := int(statusCode)
	if stc == 200 {
		log.Println("--- OKAY ---")
		rj, err := json.MarshalIndent(jsonMap, "", "  ")
		if err != nil {
			log.Fatal("Error marshalling JSON : ", err)
			os.Exit(11)
		} else {
			log.Println(string(rj))
		}
	} else {
		log.Fatal(">>> FAIL <<<\nStatusCode = ", stc)
		os.Exit(12)
	}
}
