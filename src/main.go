package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

var userNameSearch string

func main() {
	fmt.Println("Welcome to Moriarty!")

	// Arguments
	cliArgs := os.Args[1:]

	//TODO: create a argument parser
	userNameSearch = cliArgs[0]

	//read json file (its equals to sherlock-project/sherlock)
	jsonFile, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf(string(jsonFile))

	var sites map[string]interface{}

	json.Unmarshal([]byte(jsonFile), &sites)

	for key, value := range sites {
		fmt.Println(key)
		fmt.Println(reflect.ValueOf(value)) //Parei aqui seila como cata os valores
		fmt.Println("_______________________________________________________")

	}

}
