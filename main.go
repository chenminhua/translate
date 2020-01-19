package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// {"sentences":[{"trans":"你好，世界","orig":"hello world","backend":1}],"src":"en","spell":{}}
type JsonType struct {
	Sentences []Sentence `json:"sentences"`
}

type Sentence struct {
	Trans   string `json:"trans"`
	Orig    string `json:"orig"`
	Backend int    `json:"backend"`
}

const GOOGLE_TRANS_URL = "https://translate.googleapis.com/translate_a/single?" +
	"client=gtx&ie=UTF-8&dj=1&sl=en&tl=zh&dt=t&q="

func translate(query string) {
	res, err := http.Get(GOOGLE_TRANS_URL + query)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	println(string(body))
	arr := JsonType{}
	json.Unmarshal(body, &arr)
	println(arr.Sentences[0].Trans)
}

func main() {
	translate("hello")
}
