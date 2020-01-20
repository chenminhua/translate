package router

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"translate/translate"
)

const (
	NoTargetLanguageError = "请指定目标语言tl，please specify target language (tl)"
)

type Handlers struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{logger: logger}
}

// curl http://localhost:8080/upload?tl=zh -F "file=@translate/en.txt" -vvv
func ReceiveFile(w http.ResponseWriter, request *http.Request) {
	sl := request.FormValue("sl")
	if sl == "" {
		sl = "auto"
	}

	tl := request.FormValue("tl")
	if tl == "" {
		http.Error(w, NoTargetLanguageError, http.StatusBadRequest)
		return
	}

	request.ParseMultipartForm(32 << 20) // limit your max input length!
	var buf bytes.Buffer
	file, header, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Printf("File name %s\n", header.Filename)
	io.Copy(&buf, file)
	contents := strings.Split(buf.String(), "\n")

	res := translate.TranslateBatch(sl, tl, contents)
	rr := ""
	for _, r := range res {
		rr += r
		rr += "\n"
	}
	w.Write([]byte(rr))

	buf.Reset()
	return
}

func Home(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("hello world"))
}

// curl "localhost:8080/trans?tl=zh&q=hello"
func Trans(w http.ResponseWriter, request *http.Request) {
	sl := request.FormValue("sl")
	if sl == "" {
		sl = "auto"
	}

	tl := request.FormValue("tl")
	if tl == "" {
		http.Error(w, NoTargetLanguageError, http.StatusBadRequest)
		return
	}

	q := request.FormValue("q")
	if q == "" {
		http.Error(w, "请输入查询内容", http.StatusBadRequest)
		return
	}

	res := translate.Translate(sl, tl, q)
	_, _ = w.Write([]byte(res))
}

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", Home)
	mux.HandleFunc("/trans", Trans)
	mux.HandleFunc("/upload", ReceiveFile)
	return mux
}
