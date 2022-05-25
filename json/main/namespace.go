package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {

	file, err := os.ReadFile("resources/namespace.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonText := string(file)

	var jsonObject map[string]interface{}

	err1 := json.Unmarshal([]byte(jsonText), &jsonObject)
	if err1 != nil {
		return
	}

	_, err2 := fmt.Println(jsonObject["$meta"])
	if err2 != nil {
		return
	}

	for k, v := range jsonObject["$meta"].(map[string]interface{}) {
		_, err := fmt.Println("value type:", reflect.ValueOf(v).Kind().String())
		if err != nil {
			return
		}
		if reflect.ValueOf(v).Kind().String() == "map" {
			_, err := fmt.Println(k)
			if err != nil {
				return
			}
			for l, w := range v.(map[string]interface{}) {
				_, err := fmt.Println("  ", l, "=>", w)
				if err != nil {
					return
				}
			}
		} else {
			_, err := fmt.Println(k, "=>", v)
			if err != nil {
				return
			}
		}
	}

	//fmt.Println(jsonObject)
}
