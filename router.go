package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	tmpl       *template.Template
	skey       = []byte(config.SessionKey)
	store      = sessions.NewCookieStore(skey)
	cookieName = "_nethelper_go_cookie"
	ctx        = make(map[string]interface{})
)

func index(w http.ResponseWriter, r *http.Request) {
	if !isLogined(w, r) {
		setFlash(w, r, "请先登录")
		http.Redirect(w, r, "/login", 302)
	}
	ctx["Current"] = getCurrent()
	ctx["Info"] = ""
	tmpl.ExecuteTemplate(w, "Index", ctx)
}

func school(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	fields := strings.Split(ip, ".")
	room := fields[2]

	cmd := exec.Command("sh", config.CommandPath+"/school")
	err := cmd.Run()
	ctx["Current"] = getCurrent()
	if err != nil {
		log.Println(err.Error())
		ctx["Info"] = "运行失败"
	} else {
		ctx["Info"] = "运行成功"
		setCurrent(room, "school")
	}
	tmpl.ExecuteTemplate(w, "Index", ctx)
}

func gotonet(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	fields := strings.Split(ip, ".")
	room := fields[2]
	fname := "/net" + string(room)
	cmd := exec.Command("sh", config.CommandPath+fname)
	err := cmd.Run()
	ctx["Current"] = getCurrent()
	if err != nil {
		ctx["Info"] = "运行失败"
	} else {
		ctx["Info"] = "运行成功"
		setCurrent(room, "net")
	}

	tmpl.ExecuteTemplate(w, "Index", ctx)
}

func login(w http.ResponseWriter, r *http.Request) {
	ctx["Info"] = getFlash(w, r)
	tmpl.ExecuteTemplate(w, "Login", ctx)
}

func auth(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)
	session.Values["logined"] = false
	if r.FormValue("password") == config.Pass {
		session.Values["logined"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	} else {
		setFlash(w, r, "密码不对")
		http.Redirect(w, r, "/login", 302)
	}
}

func getFlash(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, cookieName)
	if str, ok := session.Values["flash"].(string); ok {
		session.Values["flash"] = nil
		session.Save(r, w)
		return str
	}
	return ""
}

func setFlash(w http.ResponseWriter, r *http.Request, str string) {
	session, _ := store.Get(r, cookieName)
	session.Values["flash"] = str
	session.Save(r, w)
}

func isLogined(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, cookieName)
	if auth, ok := session.Values["logined"].(bool); !ok || !auth {
		return false
	}
	return true
}

func router() *mux.Router {

	store.Options.MaxAge = 0
	tmpl, _ = tmpl.ParseGlob("views/*")

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/school", school)
	r.HandleFunc("/gotonet", gotonet)
	r.HandleFunc("/login", login)
	r.HandleFunc("/auth", auth).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notFound)
	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return r
}

func notFound(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "NotFound", ctx)
}
