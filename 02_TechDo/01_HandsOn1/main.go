// ソース元:https://qiita.com/yukpiz/items/64e2874fd37a9dc7365d

// 事前準備：外部パッケージの導入
// $ go get github.com/julienschmidt/httprouter

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

// /Hello/:lang にハンドルされている Hello 関数
func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lang := p.ByName("lang")  //lang パラメータを取得する
	fmt.Fprintf(w, lang)      //レスポンスに値を書き込む
}

// /Example にハンドルされている Example 関数
func Example(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close() // Example 関数が終了するときに実行される defer ステートメント

	// リクエストボディを読み取る
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// リクエストの読み取りに失敗した場合、400 エラー
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// JSONパラメータを構造体にするための定義
	type ExampleParameter struct {
		ID		int		`json:"id"`
		Name	string	`json:"name"`
	}
	var param ExampleParameter

	// bodyBytes を ExampleParameter 構造体に変換
	err = json.Unmarshal(bodyBytes, &param)
	if err != nil {
		// JSONパラメータ類を構造体の変換に失敗した場合、400 エラー
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 構造体に変換した ExampleParameter を文字列にしてレスポンスに書き込む
	fmt.Fprintf(w, fmt.Sprintf("%+v\n", param))
}

// /FizzBuzz/:num にハンドルされている FizzBuzz 関数
func FizzBuzz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	num, _ := strconv.Atoi(p.ByName("num"))

	if num < 1 {
		http.Error(w, "400", http.StatusBadRequest)
	}

	for i:=1; i<=num; i++ {
		if i%3==0 {
			if i%5==0 {
				fmt.Fprintf(w, fmt.Sprintf("%d FizzBuzz!\n", i))
			} else {
				fmt.Fprintf(w, fmt.Sprintf("%d Fizz\n", i))
			}
		} else if i%5==0 {
			fmt.Fprintf(w, fmt.Sprintf("%d Buzz\n", i))
		} else {
			fmt.Fprintf(w, fmt.Sprintf("%d\n", i))
		}
	}
}

// プロフィール用 struct
type profile struct {
	Name string `json:"name"`
	Age	 int `json:"age"`
	Gender string `json:"gender"`
	Favorite_foods []string `json:"favorite_foods"`
}
var profiles map[string]profile

func InitProfiles() {

	profiles = make(map[string]profile)
	// Bob, Alice のプロフィールをセット
	bob := profile {
		Name:           "Bob",
		Age:            25,
		Gender:         "Man",
		Favorite_foods: []string{"Hamburger","Cookie","Chocolate"},
	}
	alice := profile {
		Name:           "Alice",
		Age:            24,
		Gender:         "Woman",
		Favorite_foods: []string{"Apple","Orange","Melon"},
	}
	profiles["Bob"] = bob
	profiles["Alice"] = alice
}

// /Profile/:name にハンドルされている Profile 関数
// :name が Bob または Alice ならプロフィールを JSON 形式で返す
func Profile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// GET メソッドの時、プロフィールの表示
	if r.Method == http.MethodGet {
		name := p.ByName("name")

		if name != "Bob" && name != "Alice" {
			http.Error(w, "invalid name", http.StatusBadRequest)
		}
		if name == "Bob" {
			bytes, _ := json.Marshal(&profiles[0])
			fmt.Fprintf(w, fmt.Sprintf("%s\n",string(bytes)))
		}
		if name == "Alice" {
			bytes, _ := json.Marshal(&profiles[1])
			fmt.Fprintf(w, fmt.Sprintf("%s\n",string(bytes)))
		}
	}

	// POST メソッドの時、プロフィールの登録
	if r.Method == http.MethodPost {

	}
}

func main() {
	router := httprouter.New() //HTTPルーターを初期化

	// /Hello に GET リクエストがあったら Hello 関数にハンドルする
	// :lang はパラメータとして扱われる
	router.GET("/Hello/:lang", Hello)

	// /Example に POST リクエストがあったら Example 関数にハンドルする
	router.POST("/Example", Example)

	router.GET("/FizzBuzz/:num", FizzBuzz)

	InitProfiles()
	router.GET("/Profile/:name", Profile)
	router.POST("/Profile", Profile)

	// Web サーバを 8080 ポートで立ち上げる
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

// Web サーバの起動
// $ go run sample.go

// Web サーバにアクセス
// ・ブラウザで http://localhost:8080/Hello/golang にアクセス
// ・$ curl -XPOST -d "{\"id\": 1, \"name\": \"yukpiz\"}" http://localhost:8080/Example