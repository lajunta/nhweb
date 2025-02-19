package main

import (
	"net/http"
	"os"
	"strings"
)

func getCurrent(r *http.Request) string {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	room := rooms[ip].Num
	self := "net" + room
	cpath := workDir() + "/current"
	b, _ := os.ReadFile(cpath)
	current := strings.TrimSuffix(string(b), "\n")

	if current == "allnet" || current == self {
		return "net"
	}

	return "school"
}
