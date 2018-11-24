package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"strings"
)

// プロフィール用 struct
type profile struct {
	Name string `json:"name"`
	Age	 int `json:"age"`
	Gender string `json:"gender"`
	Favorite_foods []string `json:"favorite_foods"`
}

func main() {
	
	var (
		name = flag.String("name", "Unknown", "-name")
		age = flag.Int("age", -1, "-age")
		gender = flag.String("gender", "NULL", "-gender")
		favorite_foods = flag.String("favorite_foods", "NULL", "favorite_foods")
	)
	flag.Parse()

	if *age==-1 && *gender=="NULL" && *favorite_foods=="NULL" {
		res, err1 := http.Get("http://localhost:8080/Profile/"+*name)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer res.Body.Close()
		
		byteArray, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println(string(byteArray))
	
	} else if *age!=-1 && *gender!="NULL" && *favorite_foods!="NULL" {

		_profile := profile{
			Name:           *name,
			Age:            *age,
			Gender:         *gender,
			Favorite_foods: strings.Split(*favorite_foods, ","),
		}
		_bytes, err1 := json.Marshal(&_profile)
		if err1 != nil {
			log.Fatal(err1)
		}

		res, err2 := http.Post("http://localhost:8080/Profile/", "application/json", bytes.NewReader(_bytes))
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println(res)

	} else {
		fmt.Println("invalid args")
	}


}