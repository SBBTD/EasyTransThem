// api
package api

import (
	"EasyTransThem/pages"
	"EasyTransThem/settings"
	"io"
	"log"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	logRequestUrl(r)
	if r.Method == "POST" {
		check := func(w http.ResponseWriter, err error) {
			if err != nil {
				w.WriteHeader(500)
				io.WriteString(w, "<h1>500 Internal Server Error</h1>")
				log.Println("[ERROR]", err.Error())
			}
		}

		file, headers, err := r.FormFile("file")
		check(w, err)

		localFile, err := os.Create(settings.FilesPath + headers.Filename)
		check(w, err)

		defer func() {
			if localFile != nil {
				localFile.Close()
			}
		}()

		_, err = io.Copy(localFile, file)
		check(w, err)

		io.WriteString(w, "Upload Success!")
	} else {
		w.WriteHeader(400)
		io.WriteString(w, "<h1>400 Bad Request: POST ONLY.</h1>")
	}
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	logRequestUrl(r)
	if r.Method == "GET" {
		//TODO
	} else {
		w.WriteHeader(400)
		io.WriteString(w, "<h1>400 Bad Request: GET ONLY.</h1>")
	}
}

func Pages(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/favicon.ico" {
		logRequestUrl(r)
	}
	switch r.URL.Path {
	case "/":
		io.WriteString(w, pages.PageIndex)
	default:
		w.WriteHeader(404)
		io.WriteString(w, pages.PageNotFound)
	}
}

func logRequestUrl(r *http.Request) {
	log.Printf("[Request]From:" + r.RemoteAddr + " URL:" + r.URL.Path)
}
