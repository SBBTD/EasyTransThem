// api
package api

import (
	"EasyTransThem/pages"
	"EasyTransThem/settings"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	logRequestUrl(r)
	if r.Method == "POST" {
		check := func(w http.ResponseWriter, err error) {
			if err != nil {
				responseHttpErrorAndPrintLog(w, http.StatusInternalServerError, nil)
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
		responseHttpErrorAndPrintLog(w, http.StatusBadRequest, nil)
	}
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	logRequestUrl(r)
	if r.Method == "GET" {
		f := filepath.Base(r.URL.Query().Get("f"))
		if len(f) > 0 {
			fileToDel := settings.FilesPath + f
			fi, err := os.Stat(fileToDel)
			if err == nil && fi.Mode().IsRegular() {
				err = os.Remove(fileToDel)
				if err != nil {
					responseHttpErrorAndPrintLog(w, http.StatusInternalServerError, err)
				} else {
					io.WriteString(w, "Delete Success.")
					return
				}
			}
			responseHttpErrorAndPrintLog(w, http.StatusNotFound, nil)
			return
		}
	}
	responseHttpErrorAndPrintLog(w, http.StatusBadRequest, nil)
}

func FileList(w http.ResponseWriter, r *http.Request) {
	listHTML := ""
	flagEmpty := true
	fList, err := ioutil.ReadDir(settings.FilesPath)
	if err != nil {
		responseHttpErrorAndPrintLog(w, http.StatusInternalServerError, err)
		return
	}
	for _, f := range fList {
		if f.Mode().IsRegular() {
			listHTML += `<tr><td>`
			listHTML += f.Name()
			listHTML += `</td><td>`
			listHTML += `<a href="/files/` + f.Name() + `" target="_blank">Download</a>`
			listHTML += `</td><td>`
			listHTML += `<a href="#" onclick="del('` + f.Name() + `');">Delete</a>`
			listHTML += `</td></tr>`
			flagEmpty = false
		}
	}
	if flagEmpty {
		io.WriteString(w, "No Files Uploaded.")
	} else {
		io.WriteString(w, pages.PageFileList1+listHTML+pages.PageFileList2)
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
		responseHttpErrorAndPrintLog(w, http.StatusNotFound, nil)
	}
}

func responseHttpErrorAndPrintLog(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	switch statusCode {
	case http.StatusBadRequest:
		io.WriteString(w, "<html><body><h2>400 Bad Request</h2></body></html>")
	case http.StatusNotFound:
		io.WriteString(w, pages.PageNotFound)
	case http.StatusInternalServerError:
		io.WriteString(w, "<html><body><h2>500 Internal Server Error</h2></body></html>")
	}
	if err != nil {
		log.Println(err.Error())
	}
}

func logRequestUrl(r *http.Request) {
	log.Printf("[Request]From:" + r.RemoteAddr + " URL:" + r.URL.Path)
}
