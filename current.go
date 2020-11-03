package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func getCurrent(w http.ResponseWriter, r *http.Request) string {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	room := rooms[ip].Num
	self := "net" + room
	cpath := workDir() + "/current"
	b, _ := ioutil.ReadFile(cpath)
	current := strings.TrimSuffix(string(b), "\n")

	if current == "allnet" || current == self {
		return "net"
	}

	return "school"
}
