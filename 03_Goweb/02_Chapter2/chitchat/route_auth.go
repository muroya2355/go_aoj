package main

import (
	"net/http"
)

// クッキーを使ったアクセス制御 ユーザを認証し、クライアントにクッキーを返す
func authenticate(w http.ResponseWriter, r *http.Request) {

	// フォームの値を取得
	r.ParseForm()
	// フォームに入力された email の値から、user 構造体を返す
	user, _ := data.UserByEmail(r.PostFormValue("email"))

	// 暗号化されたパスワードが一致するか確認
	if user.Password == data.Encrypt(r.PostFormValue("password")) {

		// 一致した場合、セッションを生成しクッキーを発行
		session := user.CreateSession()
		cookie := http.Cookie{
			Name: "_cookie",
			// ログインユーザIDをクッキーに登録
			Value:    session.Uuid,
			HttpOnly: true,
		}

		// レスポンスヘッダにクッキーを追加
		http.SetCookie(w, &cookie)
		// ルートURL にリダイレクト
		http.Redirect(w, r, "/", 302)
	} else {

		// 認証に失敗した場合、ログインURL にリダイレクト
		http.Redirect(w, r, "/login", 302)
	}

}
