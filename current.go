package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func getCurrent(w http.ResponseWriter, r *http.Request) string {
	addr := strings.Split(r.RemoteAddr, ":")
	fields := strings.Split(addr[0], ".")
	room := fields[2]
	// c304 is special num
	if room == "7" {
		room = "6"
	}
	self := "net" + room
	cpath := workDir() + "/current"
	b, _ := ioutil.ReadFile(cpath)
	current := strings.TrimSuffix(string(b), "\n")

	if current == "allnet" || current == self {
		return "net"
	}
	return "school"
}
