// api
package api

import (
	"EasyTransThem/pages"
	"io"
	"log"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	logRequestUrl(r)
	if r.Method == "POST" {
		//TODO
	} else {
		w.WriteHeader(400)
		_, _ = io.WriteString(w, "400 Bad Request: POST ONLY.")
	}
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	logRequestUrl(r)
	if r.Method == "GET" {
		//TODO
	} else {
		w.WriteHeader(400)
		_, _ = io.WriteString(w, "400 Bad Request: GET ONLY.")
	}
}

func Pages(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/favicon.ico" {
		logRequestUrl(r)
	}
	switch r.URL.Path {
	case "/":
		_, _ = io.WriteString(w, pages.PageIndex)
	case "/favicon.ico":
		//TODO
	default:
		_, _ = io.WriteString(w, pages.PageNotFound)
	}
}

func logRequestUrl(r *http.Request) {
	log.Printf("[Request]From:"+ r.RemoteAddr +" URL:" + r.URL.Path)
}
