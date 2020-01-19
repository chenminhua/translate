package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// {"sentences":[{"trans":"你好，世界","orig":"hello world","backend":1}],"src":"en","spell":{}}
type TranslateResult struct {
	Sentences []Sentence `json:"sentences"`
}

type Sentence struct {
	Trans   string `json:"trans"`
	Orig    string `json:"orig"`
	Backend int    `json:"backend"`
}

const GOOGLE_TRANS_URL = "https://translate.googleapis.com/translate_a/single?"

func NewQueryString(sl string, tl string, query string) string {
	return fmt.Sprintf("client=gtx&ie=UTF-8&dj=1&dt=t&sl=%s&tl=%s&q=%s", sl, tl, url.QueryEscape(query))
}

func translate(sl string, tl string, query string) string {
	//println(GOOGLE_TRANS_URL + NewQueryString(sl, tl, query) )
	res, err := http.Get(GOOGLE_TRANS_URL + NewQueryString(sl, tl, query) )
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := TranslateResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Sentences[0].Trans
}

func main() {
	translate("en", "zh", "Building Graphical Applications with Wasmer and WASI")
}
