package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Url  string
	Name string
}

func get(api Api) {
	res, _ := http.Get(api.Url)
	body, _ := ioutil.ReadAll(res.Body)
	var x = fmt.Sprintf("Name: %s, Result: %s", api.Name, body)
	fmt.Println(x)
}

func main() {

	dat, err := ioutil.ReadFile("apilist.json")

	if err == nil {
		var apis []Api
		json.Unmarshal([]byte(dat), &apis)
		for i := 0; i < len(apis); i++ {
			get(apis[i])
		}
	}
}
