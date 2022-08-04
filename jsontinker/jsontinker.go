package main

import (
	json "encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	fname := "config.json"
	var jsonMap map[string]interface{}
	pdata, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("ReadFile error : ", err)
	} else {
		err = json.Unmarshall([]byte(pdata), &jsonMap)
		if err != nil {
			fmt.Println("Unmarshall error : ", err)
		} else {
			userid = jsonMap["userid"]
			password = jsonMap["password"]
			host = jsonMap["host"]
			port = jsonMap["port"]
			dbname = jsonMap["dbname"]
		}
	}
	fmt.Println("Values retrieved from ", fname)
	fmt.Println("userid   = ", userid)
	fmt.Println("password = ", password)
	fmt.Println("host     = ", host)
	fmt.Println("port     = ", port)
	fmt.Println("dbname   = ", dbname)
}
