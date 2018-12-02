package main

import (
	"net/http"
)

// ハンドラ関数　index : HTML を生成して ResponseWriter に書き出す
func index(w http.ResponseWriter, r *http.Request){
	
	// template ファイルの登録
	files := []string{"templates/layout.html", "templates/navbar.html", "templates/index.html",}
	// HTMLテンプレートのパース
	templates := template.Must(template.ParseFiles(files...))
	
	threads.err := data.Threads()
	// エラーがない場合、template 内容の書き出し
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

// メイン関数
func main() {

	// マルチプレクサの生成
	mux := http.NewServeMux()
	// ファイルの登録
	files := http.FileServer(http.Dir("/public"))
	// プレフィックス（"/static"）を削除した URL 上のファイルを探し、そのままファイルを返送
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// ルートURL をハンドラ関数にリダイレクト
	mux.HandleFunc("/", index)
	
	// main パッケージ内に定義されているハンドラ関数の登録
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signup_Account)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// サーバの生成、マルチプレクサの登録
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServer()
}
