package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/miniyk2012/awesomego/good_example/1-session/session"

	_ "github.com/miniyk2012/awesomego/good_example/1-session/memory"
)

var globalSessions *session.Manager

func init() {
	var err error
	globalSessions, err = session.NewManager("memory", "goSessionid", 3600)
	if err != nil {
		fmt.Println(err)
		return
	}
	go globalSessions.GC()
	fmt.Println("fd")
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("goSessionid")
	if err == nil {
		fmt.Fprintln(w, cookie.Value)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	val := sess.Get("username")
	if val != nil {
		fmt.Fprintln(w, val)
	} else {
		sess.Set("username", "jerry")
		fmt.Fprintln(w, "set session")
	}
}

func loginOut(w http.ResponseWriter, r *http.Request) {
	//销毁
	globalSessions.SessionDestroy(w, r)
	fmt.Fprintln(w, "session destroy")
}

func main() {
	http.HandleFunc("/", sayHelloHandler) //	设置访问路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", loginOut) //销毁
	log.Fatal(http.ListenAndServe(":8080", nil))
}
