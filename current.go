package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var status = map[string]string{
	"school": "校园网",
	"allnet": "正在使用互联网",
	"net1":   "C401正在上网",
	"net2":   "C402正在上网",
	"net3":   "C403正在上网",
	"net4":   "C404正在上网",
	"net5":   "C303正在上网",
	"net6":   "C304正在上网",
	"net9":   "C503正在上网",
	"net10":  "C504正在上网",
	"net11":  "C502正在上网",
	"net13":  "C103正在上网",
}

func getCurrent() string {
	cpath := config.CommandPath + "/current"
	b, _ := ioutil.ReadFile(cpath)
	current := strings.TrimSuffix(string(b), "\n")
	return status[current]
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
	self := "net" + room
	neibor := "net" + strconv.Itoa(nnum)
	b, _ := ioutil.ReadFile(cpath)
	str := strings.TrimSuffix(string(b), "\n")
	wstr := ""
	if (str == "school") || (str == self) {
		if tag == "school" {
			wstr = "school"
		} else if tag == "net" {
			wstr = self
		}
	}
	if (str == neibor) || (str == "allnet") {
		if tag == "school" {
			wstr = neibor
		} else if tag == "net" {
			wstr = "allnet"
		}
	}

	ioutil.WriteFile(cpath, []byte(wstr), 0755)
}
