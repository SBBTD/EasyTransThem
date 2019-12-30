// router
package router

import (
	"EasyTransThem/api"
	"EasyTransThem/settings"
	"net/http"
)

func init() {
	//pages
	http.HandleFunc("/", api.Pages)
	http.HandleFunc("/list/", api.FileList)

	//api
	http.HandleFunc("/api/upload", api.UploadFile)
	http.HandleFunc("/api/delete", api.DeleteFile)

	//files
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(settings.FilesPath))))
}
