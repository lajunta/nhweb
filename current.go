package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
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
	cpath := config.CommandPath + "/current"
	b, _ := ioutil.ReadFile(cpath)
	current := strings.TrimSuffix(string(b), "\n")

	if current == "allnet" || current == self {
		return "net"
	}
	return "school"
}

func setCurrent(room string, tag string) {
	cpath := config.CommandPath + "/current"
	num, _ := strconv.Atoi(room)
	if num == 7 {
		num = 6
	}
	nnum := 0
	if num%2 == 0 {
		nnum = num - 1
	} else {
		nnum = num + 1
	}
	self := "net" + strconv.Itoa(num)
	neibor := "net" + strconv.Itoa(nnum)
	b, _ := ioutil.ReadFile(cpath)
	current := strings.TrimSuffix(string(b), "\n")
	wstr := ""
	if (current == "school") || (current == self) {
		if tag == "school" {
			wstr = "school"
		} else if tag == "net" {
			wstr = self
		}
		ioutil.WriteFile(cpath, []byte(wstr), 0755)
	}
	if (current == neibor) || (current == "allnet") {
		if tag == "school" {
			wstr = neibor
		} else if tag == "net" {
			wstr = "allnet"
		}
		ioutil.WriteFile(cpath, []byte(wstr), 0755)
	}
}
