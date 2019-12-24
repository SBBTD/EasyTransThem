// boot
package boot

import (
	"EasyTransThem/settings"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var (
	h bool = false
	v bool = false
)

func init() {
	//Paths
	settings.RootPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	settings.FilesPath = settings.RootPath + "/files/"
	//Server
	flag.UintVar(&settings.ServerPort, "p", settings.ServerPort, "Server port for listening.")
	//sys
	flag.BoolVar(&h, "h", false, "Show this help.")
	flag.BoolVar(&v, "v", false, "Show version information.")

	flag.Parse()
	if h {
		println(usagePrefix())
		flag.PrintDefaults()
		os.Exit(0)
	}
	if v {
		println(version())
		os.Exit(0)
	}

	err := os.MkdirAll(settings.FilesPath, os.ModePerm)
	if err!=nil{
		log.Fatalln("[ERROR]Mkdir \"files\" Failed.")
	}

	log.Println(version())
	log.Println("[Settings] FilesPath:" + settings.FilesPath)
	log.Printf("[Settings] ServerPort:%d\n", settings.ServerPort)
}

func usagePrefix() string {
	return version() + "\nUsage:"
}

func version() string {
	return settings.AppName + " " + settings.Version + "(" + strconv.Itoa(int(settings.VersionCode)) + ")"
}
