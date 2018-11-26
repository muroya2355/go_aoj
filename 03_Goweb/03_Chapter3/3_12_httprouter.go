package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Params 型には名前付きパラメータが入っており、メソッド ByName で取得できる
func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s\n", p.ByName("name"))
}

func main() {

	// マルチプレクサの生成
	mux := httprouter.New()
	// GETメソッド用のURL に対してハンドラ関数を登録
	mux.GET("/hello/:name", hello)

	// マルチプレクサをサーバ構造体に入れて使う
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
