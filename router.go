package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

//go:embed static
var Assets embed.FS

//go:embed views
var Views embed.FS

var (
	err        error
	tmpl       *template.Template
	skey       = []byte(config.SessionKey)
	store      = sessions.NewCookieStore(skey)
	cookieName = "_nethelper_go_cookie"
	ctx        = make(map[string]interface{})
)

func index(w http.ResponseWriter, r *http.Request) {
	ctx["Current"] = getCurrent(r)
	ctx["Logined"] = isLogined(r)
	err := tmpl.ExecuteTemplate(w, "Index", ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func workDir() string {
	return config.CommandPath
}

func school(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	neibor := rooms[ip].Neibor
	schoolPath := workDir() + "/school"
	neiborPath := workDir() + "/net" + neibor

	fname := schoolPath
	if netStatus() == "allnet" || netStatus() == "net"+neibor {
		fname = neiborPath
	}

	cmd := exec.Command("sh", "-c", fname)
	err = cmd.Run()

	if err != nil {
		log.Println(err.Error())
		setSession(w, r, "flash", "操作失败")
	} else {
		setSession(w, r, "flash", "操作成功")
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func netStatus() string {
	currentPath := filepath.Join(workDir(), "current")
	data, _ := os.ReadFile(currentPath)
	return strings.TrimSpace(string(data))
}

func gotonet(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	num := rooms[ip].Num
	neibor := rooms[ip].Neibor
	netPath := workDir() + "/net" + num
	allnetPath := workDir() + "/allnet"

	fname := netPath
	if netStatus() == "allnet" || netStatus() == "net"+neibor {
		fname = allnetPath
	}

	exec.Command("sh", "-c", fname).Run()
	http.Redirect(w, r, "/", http.StatusFound)
}

func auth(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)
	session.Values["logined"] = false
	if r.FormValue("password") == config.Pass {
		session.Values["logined"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		setSession(w, r, "flash", "密码不对")
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)
	session.Values["logined"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func setSession(w http.ResponseWriter, r *http.Request, sname, str string) {
	session, _ := store.Get(r, cookieName)
	session.Values[sname] = str
	session.Save(r, w)
}

func isLogined(r *http.Request) bool {
	session, _ := store.Get(r, cookieName)
	if logined, ok := session.Values["logined"].(bool); ok {
		return logined
	}
	return false
}

func router() *mux.Router {
	store.Options.MaxAge = 0
	tmpl, _ = template.ParseFS(Views, "views/*.html")
	r := mux.NewRouter()
	fs := http.FileServer(http.FS(Assets))
	r.HandleFunc("/", index)
	r.HandleFunc("/school", checkIP(school))
	r.HandleFunc("/gotonet", checkIP(gotonet))
	r.HandleFunc("/logout", checkIP(logout))
	r.HandleFunc("/auth", auth).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notFound)
	// fs = http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/", fs))
	return r
}

func notFound(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "NotFound", ctx)
}
