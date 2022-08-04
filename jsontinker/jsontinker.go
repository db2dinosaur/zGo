package main

import (
	json "encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	fname := "config.json"
	userid := "<<unknown>>"
	password := "<<unknown>>"
	host := "<<unknown>>"
	port := "<<unknown>>"
	dbname := "<<unknown>>"
	var jsonMap map[string]interface{}
	pdata, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("ReadFile error : ", err)
	} else {
		err = json.Unmarshal([]byte(pdata), &jsonMap)
		if err != nil {
			fmt.Println("Unmarshal error : ", err)
		} else {
			userid = jsonMap["userid"].(string)
			password = jsonMap["password"].(string)
			host = jsonMap["host"].(string)
			port = jsonMap["port"].(string)
			dbname = jsonMap["dbname"].(string)
		}
	}
	fmt.Println("Values retrieved from ", fname)
	fmt.Println("userid   = ", userid)
	fmt.Println("password = ", password)
	fmt.Println("host     = ", host)
	fmt.Println("port     = ", port)
	fmt.Println("dbname   = ", dbname)
}
