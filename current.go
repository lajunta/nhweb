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
	"net7":   "C304正在上网",
	"net8":   "C401正在上网",
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
	if tag == "school" {
		switch str {
		case "school":
			wstr = "school"
		case "allnet":
			wstr = neibor
		case neibor:
			wstr = neibor
		case self:
			wstr = "school"
		}
	}
	if tag == "net" {
		switch str {
		case "school":
			wstr = self
		case "allnet":
			wstr = "allnet"
		case neibor:
			wstr = "allnet"
		case self:
			wstr = self
		}
	}
	ioutil.WriteFile(cpath, []byte(wstr), 0755)
}
