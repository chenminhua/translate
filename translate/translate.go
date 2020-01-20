package translate

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

const GOOGLE_TRANS_URL = "https://Translate.googleapis.com/translate_a/single?"

func NewQueryString(sl string, tl string, query string) string {
	return fmt.Sprintf("client=gtx&ie=UTF-8&dj=1&dt=t&sl=%s&tl=%s&q=%s", sl, tl, url.QueryEscape(query))
}

func Translate(sl string, tl string, query string) string {
	//println(GOOGLE_TRANS_URL + NewQueryString(sl, tl, query) )
	// todo 定义错误类型，方便caller 处理
	// todo we should detect errors in low level, and handle them in high level.
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

func TranslateBatch(sl string, tl string, queries []string) []string {
	res := make([]string, len(queries))
	for i, query := range queries {
		res[i] = Translate(sl, tl, query)
	}
	return res
}