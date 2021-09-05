package main

import (
	"fmt"
	"net"

	"net/http"
	"encoding/json"
	"github.com/braintree/manners"
	"github.com/julienschmidt/httprouter"
)

// func handleSignal() {
// 	signalChan := make(chan os.Signal, 10)
// 	signal.Notify(signalChan, syscall.SIGTERM)
// 	go func() {
// 		for {
// 			s := <-signalChan
// 			if s == syscall.SIGTERM {
// 				manners.Close()
// 			}
// 		}
// 	}()
// }

type Parameters struct {
	Sid     string `json:"sid"`
}

func SetPatrameters(r *http.Request) []byte {
	var parameters = Parameters{}
	parameters.Sid = r.URL.Query().Get("sid")

	paramsJson, err := json.Marshal(&parameters)
	if err != nil {
		panic(err)
	}
	return paramsJson
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

  // TODO:Cookieをセットしたり
	// 外部パラメータをJSON形式で取得
	paramsJson := SetPatrameters(r)

  // 確認のためにbyteをstringにキャストする
	fmt.Print(string(paramsJson))

	// UserAgent判定
	// url := banner.Url

	// if regexp.MustCompile(`iPhone`).Match([]byte(r.UserAgent())) {
	// 	url = banner.UrlIos
	// }
	// if regexp.MustCompile(`iPad`).Match([]byte(r.UserAgent())) {
	// 	url = banner.UrlIos
	// }
	// if regexp.MustCompile(`Android`).Match([]byte(r.UserAgent())) {
	// 	url = banner.UrlAndroid
	// }

  url := "https://google.com"
	// リダイレクト処理
	http.Redirect(w, r, url, 302)
}

func main() {

	var listener net.Listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8089))
	if err != nil {
		// logger.Error("Failed to start server", zap.String("status", err.Error()))
	}

	router := httprouter.New()
	router.GET("/", Index)
	// TODO: handle の役割を知りたい
	// router.GET("/click/:media_code/:banner_sym/", handle.Click)

	// handleSignal()
	manners.Serve(listener, router)
}