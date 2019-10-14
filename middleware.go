package main

import (
	"net/http"
	"strings"
)

func checkIP(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addr := strings.Split(r.RemoteAddr, ":")
		ip := addr[0]
		include := false
		for _, v := range config.IPS {
			if v == ip {
				include = true
				break
			}
		}
		if include {
			f(w, r)
		} else {
			notFound(w, r)
		}
	}
}
