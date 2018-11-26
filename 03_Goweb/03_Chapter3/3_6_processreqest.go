package main

import (
	"fmt"
	"net/http"
)

// ハンドラ：メソッド ServerHTTP をもつインターフェース
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// サーバに届くすべてのリクエストが handler に行く
		Handler: &handler,
	}

	server.ListenAndServe()
}
