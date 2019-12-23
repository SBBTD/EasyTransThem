// TempTransThatFile project main.go
package main

import (
	_ "EasyTransThem/boot"
	_ "EasyTransThem/router"
	"EasyTransThem/settings"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(int(settings.ServerPort)), nil))
}
